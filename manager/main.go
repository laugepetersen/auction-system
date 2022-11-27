package main

import (
	auctionService "auction-system/proto"
	"bufio"
	"context"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type ManagerState int
type ElectionState int

const (
	Primary ManagerState = iota
	Backup
)

const (
	Void ElectionState = iota
	Waiting
	Done
)

type ReplicaManager struct {
	auctionService.UnimplementedAuctionServiceServer
	
	State            ManagerState
	ElectionState    ElectionState
	Port             int
	LamportTimestamp int
	Clients          map[int]auctionService.AuctionServiceClient
	ReplicaManagers  map[int]auctionService.AuctionServiceClient
	PrimaryManager   int

	maxBid 			 int
	latestBidPort 	 int

	ctx              context.Context
}

func main() {

	// Retrieve network ports from terminal args
	var serverPorts [3]int
	var clientPorts [3]int

	args1, _ := strconv.ParseInt(os.Args[1], 10, 32) // The primary port
	args2, _ := strconv.ParseInt(os.Args[2], 10, 32) // port1 (own port)
	args3, _ := strconv.ParseInt(os.Args[3], 10, 32) // port2 (other port)
	args4, _ := strconv.ParseInt(os.Args[4], 10, 32) // port3 (other port)
	serverPorts[0] = int(args2)
	serverPorts[1] = int(args3)
	serverPorts[2] = int(args4)
	ownPort := serverPorts[0]

	args5, _ := strconv.ParseInt(os.Args[5], 10, 32) // client1
	args6, _ := strconv.ParseInt(os.Args[6], 10, 32) // client2
	args7, _ := strconv.ParseInt(os.Args[7], 10, 32) // client3
	clientPorts[0] = int(args5)
	clientPorts[1] = int(args6)
	clientPorts[2] = int(args7)

	// Don't touch
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Instantiate Manager
	manager := &ReplicaManager{
		State:            Backup,
		ElectionState:    Void,
		Port:             ownPort,
		LamportTimestamp: 0,
		Clients:          make(map[int]auctionService.AuctionServiceClient),
		ReplicaManagers:  make(map[int]auctionService.AuctionServiceClient),
		PrimaryManager:   int(args1),
		ctx:              ctx,
	}

	// Make first port Primary Manager
	if args1 == int64(ownPort) {
		manager.State = Primary
	}

	// TCP-listener to port
	list, err := net.Listen("tcp", fmt.Sprintf(":%v", ownPort))

	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}

	grpcServer := grpc.NewServer()
	auctionService.RegisterAuctionServiceServer(grpcServer, manager)

	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("Failed to serve %v", err)
		}
	}()

	fmt.Printf("Listening on port %v \n", ownPort)
	fmt.Println()
	fmt.Printf("Dialing ReplicaManagers: \n")

	// Dial replica managers
	for i := 0; i < len(serverPorts); i++ {
		serverPort := serverPorts[i]

		if serverPort == manager.Port {
			continue
		}

		// Dial RM with the ports
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(fmt.Sprintf(":%v", serverPort), grpc.WithInsecure(), grpc.WithBlock())
		defer conn.Close()

		if err != nil {
			log.Fatalf("Failed connecting to %v: %s", serverPort, err)
		}

		backupManager := auctionService.NewAuctionServiceClient(conn)
		manager.ReplicaManagers[serverPort] = backupManager
		manager.LamportTimestamp += 1

		fmt.Printf("|- Successfully connected to Manager:%v \n", serverPort)
	}

	// Dial clients
	go func() {
		for i := 0; i < len(clientPorts); i++ {
			clientPort := clientPorts[i]

			// Dial clients with the ports
			var conn *grpc.ClientConn
			conn, err := grpc.Dial(fmt.Sprintf(":%v", clientPort), grpc.WithInsecure(), grpc.WithBlock())
			defer conn.Close()

			if err != nil {
				log.Fatalf("Failed dialing client clientport: %v, err: %s", clientPort, err)
			}

			client := auctionService.NewAuctionServiceClient(conn)
			manager.Clients[clientPort] = client
			manager.LamportTimestamp += 1

			fmt.Printf("|- Successfully connected to Client:%v \n", clientPort)
		}
	}()

	go manager.ListenForHeartBeat()

	// Enter to crash
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		os.Exit(0)
	}

}

func (manager *ReplicaManager) ListenForHeartBeat() {
	for {
		time.Sleep(time.Millisecond * 2000)

		if manager.isPrimary() {
			continue
		}

		// TODO: Lamport
		pingData := &auctionService.Ping{
			Host:    int32(manager.Port),
			Lamport: int32(manager.LamportTimestamp),
		}
		response, err := manager.ReplicaManagers[manager.PrimaryManager].PingClient(manager.ctx, pingData)

		if err != nil {
			fmt.Printf("|- Error in pinging PrimaryManager: %v \n", manager.PrimaryManager)
			// TODO: Start Election due to error, likely a crash
			break
		} else {
			fmt.Printf("|- RM: %v successfully pinged and recived response from: %v \n", manager.Port, response.Host)
		}
	}
}

func (manager *ReplicaManager) PingClient(ctx context.Context, in *auctionService.Ping) (*auctionService.Ping, error) {
	// TODO: Lamport
	fmt.Printf("|- Primary manager succesfully pinged by: %v \n", in.Host)
	return &auctionService.Ping{
		Host:    int32(manager.Port),
		Lamport: int32(manager.LamportTimestamp),
	}, nil
}

func (manager *ReplicaManager) Bid(ctx context.Context, in *auctionService.BidMessage) (*auctionService.Ack, error) {

	fmt.Printf("|- Successfully received bid request from: %v, with amount: %v \n", in.Host, in.Amount)

	if in.Amount <= int32(manager.maxBid) {
		// TODO: Lamport
		return &auctionService.Ack{
			Ack:     true,
			Message: "Bid is lower or equal to current max bid",
			Lamport: int32(manager.LamportTimestamp),
		}, nil
	}

	// Update auction fields
	manager.maxBid = int(in.Amount)
	manager.latestBidPort = int(in.Host)

	if manager.isPrimary() {
		
		// Notify the other replica managers a bid was made
		for port, replicaManager := range manager.ReplicaManagers {
			reponse, err := replicaManager.Bid(manager.ctx, in)

			if err != nil {
				fmt.Printf("Error notifying replica manager: %v with new bid, err: %v", port, err)
			}

			fmt.Printf("|- Successfully notifed RM of bid, returned message: %v \n", reponse.Message)
		}

		// TOOD: Lamport
		return &auctionService.Ack{
			Ack:     true,
			Message: "Thank you user, a bid has successfully been placed",
			Lamport: int32(manager.LamportTimestamp),
		}, nil
	} else {
		// TODO: Lamport
		return &auctionService.Ack{
			Ack:     true,
			Message: fmt.Sprintf("Thank you primary manager, fields are updated, for port: %v, with amount: %v", manager.Port, manager.maxBid),
			Lamport: int32(manager.LamportTimestamp),
		}, nil
	}

}

func (manager *ReplicaManager) GetResult(ctx context.Context, in *auctionService.Ping) (*auctionService.Ack, error) {

	return &auctionService.Ack{}, nil
}

func (manager *ReplicaManager) Vote(ctx context.Context, in *auctionService.VoteReq) (*auctionService.VoteRes, error) {
	voteRes := &auctionService.VoteRes{
		Answer: false,
	}

	if manager.LamportTimestamp < int(in.Lamport) {
		voteRes.Answer = true
		return voteRes, nil
	}

	if manager.LamportTimestamp > int(in.Lamport) {
		return voteRes, nil
	}

	if manager.LamportTimestamp == int(in.Lamport) {
		if manager.Port < int(in.Host) {
			voteRes.Answer = true
			return voteRes, nil
		}
	}
	return voteRes, nil
}

func (manager *ReplicaManager) StartElection() {
	// for port, _manager := range manager.ReplicaManagers {
	// 	res, _ := _manager.Vote(manager.ctx, &auctionService.VoteReq{
	// 		Host:    int32(manager.Port),
	// 		Lamport: int32(manager.LamportTimestamp),
	// 	})
	// }
}

func (manager *ReplicaManager) isPrimary() bool {
	return (manager.State == Primary)
}

func (manager *ReplicaManager) setState(State ManagerState) {
	manager.State = State
}

func (manager *ReplicaManager) setElectionState(State ElectionState) {
	manager.ElectionState = State
}

func max(ownTimestamp int32, theirTimestamp int32) int32 {
	return int32(math.Max(float64(ownTimestamp), float64(theirTimestamp)))
}
