package server

import (
	"log"
	"net"
	"google.golang.org/grpc"

)


type server struct{}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to list %v", err)
	}

	grpcServer := grpc.NewServer()
	
}