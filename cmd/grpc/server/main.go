package main

import (
	"fmt"
	"log"
	"net"

	"rest-vs-grpc/internal/pb"

	"google.golang.org/grpc"
)

const portServer = 10000

func main() {

	log.Println("GRPC running on", fmt.Sprintf("localhost:%d", portServer))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", portServer))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := pb.Server{}

	grpcServer := grpc.NewServer()
	pb.RegisterServerGrpcServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("GRPC failed to serve: %s", err)
	}
}
