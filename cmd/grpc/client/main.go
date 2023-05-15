package main

import (
	"context"
	"log"
	"rest-vs-grpc/internal/cli"
	"rest-vs-grpc/internal/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.Println("---- gRPC CLIENT TEST ----")
	log.Printf("There are %v users in response\n", len(pb.ResponsePb.Users))

	qtyRequests, err := cli.HandleParams()
	if err != nil {
		log.Fatal(err)
	}

	result, err := manyCalls(qtyRequests)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("average of %v requests is: %v/request", qtyRequests, result)
}

func manyCalls(numCalls int) (time.Duration, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn := getGRPCClient(ctx)
	defer conn.Close()

	var callsSum time.Duration
	for i := 0; i < numCalls; i++ {
		startCall := time.Now()
		_, err := callAPI(conn)
		if err != nil {
			return 0, err
		}

		requestTime := time.Since(startCall)
		callsSum = callsSum + requestTime
	}

	return callsSum / time.Duration(numCalls), nil
}

func callAPI(conn *grpc.ClientConn) (result *pb.Response, err error) {

	client := pb.NewServerGrpcClient(conn)

	resp, err := client.GetUsers(context.Background(), &pb.Null{})
	if err != nil {
		return nil, err
	}

	return resp, err
}

func getGRPCClient(ctx context.Context) *grpc.ClientConn {
	var opts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	}

	conn, err := grpc.DialContext(ctx, "localhost:10000", opts...)

	if err != nil {
		log.Fatal(err)
	}

	return conn
}
