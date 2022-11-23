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
	Waiting ElectionState = iota
	Done
)

type ReplicaManager struct {
	auctionService.UnimplementedAuctionServiceServer
	State            ManagerState
	Port             int
	LamportTimestamp int
	Clients          map[int]auctionService.AuctionServiceClient
	ReplicaManagers  map[int]auctionService.AuctionServiceClient
	ctx              context.Context
}

func main() {
	// Retrieve network ports from terminal args
	serverPorts := []int{5100, 5101, 5102}
	// clientPorts := []int{6100, 6101, 6102}

	arg, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := serverPorts[arg]

	// Don't touch
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Instantiate Manager
	manager := &ReplicaManager{
		State:            Backup,
		Port:             ownPort,
		LamportTimestamp: 0,
		Clients:          make(map[int]auctionService.AuctionServiceClient),
		ReplicaManagers:  make(map[int]auctionService.AuctionServiceClient),
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
	fmt.Printf("Dialing ReplicaManagers: \n")

	discoverReplicaManagers(serverPorts, manager)
	//go discoverClients(clientPorts, manager)

	// Heartbeats
	// go func() {
	// 	for {
	// 		time.Sleep(time.Millisecond * 2000)
	// 		manager.SendHeartBeatToAllManagers()
	// 	}
	// }()

	go func() {
		for {
			time.Sleep(time.Millisecond * 2000)
			manager.SendHeartBeatToAllManagers()
		}
	}()

	// Enter to crash
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		os.Exit(0)
	}
}

func discoverClients(clientPorts []int, manager *ReplicaManager) {
	for i := 0; i < len(clientPorts); i++ {
		clientPort := clientPorts[i]

		if clientPort == manager.Port { 
			continue
		}

		// Dial clients with the ports
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(fmt.Sprintf(":%v", clientPort), grpc.WithInsecure(), grpc.WithBlock())
		defer conn.Close()

		if err != nil {
			log.Fatalf("Failed connecting to %v: %s", clientPort, err)
		}

		client := auctionService.NewAuctionServiceClient(conn)
		manager.Clients[clientPort] = client
		manager.LamportTimestamp += 1

		fmt.Printf("|- Successfully connected to Client:%v \n", clientPort)
	}
}

func discoverReplicaManagers(serverPorts []int, manager *ReplicaManager) {
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

		client := auctionService.NewAuctionServiceClient(conn)
		manager.ReplicaManagers[serverPort] = client
		manager.LamportTimestamp += 1

		fmt.Printf("|- Successfully connected to Manager:%v \n", serverPort)
	}
}

func (manager *ReplicaManager) SendHeartBeatToAllManagers() {
	// Only primary sends heartbeat
	if !manager.isPrimary() {
		fmt.Println("!manager.isPrimary() > SendHeartBeatToAllManagers()")
		return
	}

	fmt.Println("Sending heartbeat...")

	// TODO: Some lamport stuff here
	req := &auctionService.Ping{
		Host:    int32(manager.Port),
		Lamport: int32(manager.LamportTimestamp),
	}

	for port, backupManager := range manager.ReplicaManagers {
		_, err := backupManager.HeartBeat(manager.ctx, req)

		if err != nil {
			log.Fatalf("|- Failed to send heartbeat to Manager:%v: %s \n", port, err)
		}
	}

}

func (manager *ReplicaManager) HeartBeat(ctx context.Context, in *auctionService.Ping) (*auctionService.Ping, error) {
	// Slet senere
	if manager.isPrimary() {
		fmt.Println("manager.isPrimary in HeartBeat !!!!! ERROR !!!!!")
		return &auctionService.Ping{}, nil
	}

	fmt.Printf("|- Received heartbeat from PrimaryManager:%v \n", in.Host)

	// TODO: Some lamport here
	returnPing := &auctionService.Ping{
		Host:    int32(manager.Port),
		Lamport: int32(manager.LamportTimestamp + 1),
	}
	return returnPing, nil
}

func (manager *ReplicaManager) isPrimary() bool {
	return (manager.State == Primary)
}

func max(ownTimestamp int32, theirTimestamp int32) int32 {
	return int32(math.Max(float64(ownTimestamp), float64(theirTimestamp)))
}
