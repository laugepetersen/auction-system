package main

import (
	auctionService "auction-system/proto"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
)


type Frontend struct {
	auctionService.UnimplementedAuctionServiceServer
	Port             int
	LamportTimestamp int
	PrimaryManager	 auctionService.AuctionServiceClient
	ctx              context.Context
}

func main() {
	clientPorts := []int{6100, 6101, 6102}

	arg, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := clientPorts[arg]

	// Don't touch
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Instantiate Frontend
	manager := &Frontend{
		Port:             ownPort,
		LamportTimestamp: 0,
		PrimaryManager:   nil,
		ctx:              ctx,
	}

	// TCP-listener to port
	list, err := net.Listen("tcp", fmt.Sprintf(":%v", ownPort))

	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	} else {
		fmt.Printf("Listening on port %v \n", ownPort)
	}

	grpcServer := grpc.NewServer()
	auctionService.RegisterAuctionServiceServer(grpcServer, manager)

	//error serve 
	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("Failed to serve %v", err)
		}
	}()

	// Discover Primary Manager
	PrimaryManagerPort := 5100

	fmt.Printf("Dialing Primary Manager: \n")

	var conn *grpc.ClientConn
	conn, dialErr := grpc.Dial(fmt.Sprintf(":%v", PrimaryManagerPort), grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()

	if dialErr != nil {
		log.Fatalf("Failed connecting to %v: %s", PrimaryManagerPort, err)
	}

	PrimaryManager := auctionService.NewAuctionServiceClient(conn)
	manager.PrimaryManager = PrimaryManager
	manager.LamportTimestamp += 1

	fmt.Printf("|- Successfully connected to Primary Manager:%v \n", PrimaryManagerPort)

	for {
	}
}