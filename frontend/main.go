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
	PrimaryManager   auctionService.AuctionServiceClient
	ctx              context.Context
}

func main() {
	args1, _ := strconv.ParseInt(os.Args[1], 10, 32) // primary manager port
	args2, _ := strconv.ParseInt(os.Args[2], 10, 32) // client port

	primaryManagerPort := int(args1)
	ownPort := int(args2)

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

	// Error on serve
	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("Failed to serve %v", err)
		}
	}()

	// Dial Primary Manager
	fmt.Printf("Dialing Primary Manager: \n")

	var conn *grpc.ClientConn
	conn, dialErr := grpc.Dial(fmt.Sprintf(":%v", primaryManagerPort), grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()

	if dialErr != nil {
		log.Fatalf("Failed dialing primary manager %v: %s", primaryManagerPort, err)
	}

	PrimaryManager := auctionService.NewAuctionServiceClient(conn)
	manager.PrimaryManager = PrimaryManager
	manager.LamportTimestamp += 1

	fmt.Printf("|- Successfully connected to Primary Manager:%v \n", primaryManagerPort)

	for {
	}
}
