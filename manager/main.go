package manager

import (
	auctionService "auction-system/proto"
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

type State int

const (
	Primary State = iota
	Backup
)

type ReplicaManager struct {
	auctionService.UnimplementedAuctionServiceServer
	State            State
	Port             int
	LamportTimestamp int32
	Clients          map[int]auctionService.AuctionServiceClient
	ReplicaManager   map[int]auctionService.AuctionServiceServer
	ctx              context.Context
}

func main() {

	// Retrieve network ports from terminal args
	port, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ports := make([]int32, 0)
	ports = append(ports, port)
	ports = append(ports, int32(port)+1)
	ports = append(ports, int32(port)+2)

	// Don't touch
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Instantiate Peer
	manager := &ReplicaManager{
		Port:             int32(port),
		Clients:          make(map[int32]p2p.p),
		State:            FREE,
		CSQueue:          make(map[int32]int32),
		LamportTimestamp: 0,
		ctx:              ctx,
	}

	// TCP-listener to port
	list, err := net.Listen("tcp", fmt.Sprintf(":%v", port))

	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}

	grpcServer := grpc.NewServer()
	p2p.RegisterPeerServiceServer(grpcServer, peer)

	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("Failed to serve %v", err)
		}
	}()

	fmt.Printf("Listening on port %v \n", port)
	fmt.Printf("Dialing services: \n")

	// Discover other clients/services
	for i := 0; i < len(ports); i++ {
		nodePort := ports[i]

		if nodePort == int32(port) {
			continue
		}

		// Don't touch
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(fmt.Sprintf(":%v", nodePort), grpc.WithInsecure(), grpc.WithBlock())
		defer conn.Close()

		if err != nil {
			log.Fatalf("Failed connecting to %v: %s", nodePort, err)
		}

		client := p2p.NewPeerServiceClient(conn)
		peer.Clients[nodePort] = client
		peer.LamportTimestamp += 1

		fmt.Printf("|– Successfully connected to %v \n", nodePort)
	}

	// Aggree with network of clients
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		peer.AskNetworkForPermission()
	}
}

// grpc heartbeat(), hvis primary.

// state primary || backup
// lamportTimestamp
// backups[]

// grpc bid(), track highest bid

// grpc result(), output highest bid || result if time is over

// startTime() start to end.

// grpc startElection(), new leader, update client

// compareLamport()
// func max(ownTimestamp int32, theirTimestamp int32) int32 {
//	return int32(math.Max(float64(ownTimestamp), float64(theirTimestamp)))
//}
