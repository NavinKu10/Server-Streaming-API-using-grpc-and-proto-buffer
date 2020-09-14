package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "grpc/cust"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Multiply(req *pb.Request,
	stream pb.MultiplyService_MultiplyServer) error {

	log.Println("Squaring of numbers started.....")
	a := req.GetA()
	for i := a; i < a+5; i++ {
		stream.Send(&pb.Response{
			Result: i * i,
		})
		time.Sleep(2 * time.Second)
	}

	return nil

}
func main() {

	lis, err := net.Listen("tcp", ":8088")

	if err != nil {
		log.Fatalf("Error is %v", err)
	}

	gs := grpc.NewServer()
	pb.RegisterMultiplyServiceServer(gs, &server{})

	fmt.Println("Server is running.....")
	err = gs.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}

}
