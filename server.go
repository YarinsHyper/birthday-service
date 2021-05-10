package main

import (
	"fmt"
	"log"
	"net"

	"github.com/yarinBenisty/birthday-service/controller"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen on port 8000: %v", err)
	}

	grpcServer := grpc.NewServer()

	fmt.Printf("birthday service running on port %d", 8000)

	//connecting to mongodb
	controller.NewService("mongodb://root:example@0.0.0.0:27017")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port 8000: %v", err)
	}

}
