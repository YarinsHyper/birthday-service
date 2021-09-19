package main

import (
	"fmt"
	"log"
	"net"

	"github.com/yarinBenisty/birthday-service/controller"
	pb "github.com/yarinBenisty/birthday-service/proto"
	"google.golang.org/grpc"
)

//BirthdayServer is is a structure that holds the birthday grpc server
// and its services and configuration.

func main() {
	addr := "mongodb://root:example@0.0.0.0:27017"

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen on port 8000: %v", err)
	}

	grpcServer := grpc.NewServer()
	fmt.Printf("birthday service running on port %d\n", 8000)

	// connecting to mongodb
	bdService := controller.NewService(addr)
	pb.RegisterBirthdayFunctionsServer(grpcServer, bdService)
	fmt.Printf("Successfully connected to mongo!\n")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc server over port 8000: %v", err)
	}
}
