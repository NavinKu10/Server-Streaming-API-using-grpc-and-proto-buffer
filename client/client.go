package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"

	pb "grpc/cust"
)

func callMultiply(c pb.MultiplyServiceClient) {
	stream, err := c.Multiply(context.Background(), &pb.Request{A: 2})

	if err != nil {
		log.Fatalf("callMultiply err %v", err)
	}

	log.Println("Squaring of numbers started from 2...")

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("server finished streaming")
			return
		}

		if err != nil {
			log.Fatalf("callMultiply recvErr %v", err)
		}

		log.Printf("result %v", resp)
	}

}
func main() {

	conn, err := grpc.Dial("localhost:8088", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMultiplyServiceClient(conn)

	log.Printf("Client is running.... %v", client)

	callMultiply(client)
}
