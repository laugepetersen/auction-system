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
	ctx              context.Context
}

func main() {

	// Retrieve network ports from terminal args
	serverPorts := []int{7101, 7106, 7103}
	// clientPorts := []int{6100, 6101, 6102}

	arg, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := serverPorts[arg]

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
		PrimaryManager:   serverPorts[0],
		ctx:              ctx,
	}

	// Make first port Primary Manager
	if arg == 0 {
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

	// TODO: Dial clients

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
