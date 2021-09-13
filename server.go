package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/yarinBenisty/birthday-service/controller"
	pb "github.com/yarinBenisty/birthday-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

//BirthdayServer is is a structure that holds the birthday grpc server
// and its services and configuration.
type BirthdayServer struct {
	*grpc.Server
	port                string
	healthCheckInterval int
	birthdayService     *controller.Service
}

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

	// Create a health server and register it on the grpc server.
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)

	birthdayServer := &BirthdayServer{
		Server:              grpcServer,
		port:                "8000",
		healthCheckInterval: 3,
		birthdayService:     bdService,
	}

	go birthdayServer.healthCheckWorker(healthServer)

}

func (s BirthdayServer) healthCheckWorker(healthServer *health.Server) {
	for {
		if s.birthdayService.HealthCheck(10 * time.Second) {
			healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
		} else {
			healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_NOT_SERVING)
		}

		time.Sleep(time.Second * time.Duration(s.healthCheckInterval))
	}
}
