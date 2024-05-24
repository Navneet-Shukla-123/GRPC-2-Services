package main

import (
	"context"
	"grpc-compute/compute"
	"log"
	"net"

	"google.golang.org/grpc"
)

type myComputeServer struct {
	compute.UnimplementedComputeServer
}

// Square will compute the square of the request
func (s myComputeServer) Square(ctx context.Context, req *compute.RequestBody) (*compute.ResponseBody, error) {

	requestData := req.GetX()

	newdata := requestData * requestData

	return &compute.ResponseBody{
		X: int64(newdata),
	}, nil

}

// Cube will compute the cube of the request
func (s myComputeServer) Cube(ctx context.Context, req *compute.RequestBody) (*compute.ResponseBody, error) {
	requestData := req.GetX()

	newData := requestData * requestData * requestData
	return &compute.ResponseBody{
		X: int64(newData),
	}, nil

}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error in starting the server : %v", err)
	}
	service := &myComputeServer{}

	serverRegister := grpc.NewServer()

	compute.RegisterComputeServer(serverRegister, service)

	err = serverRegister.Serve(lis)
	if err != nil {
		log.Fatalf("Error in listening to server : %v", err)
	}

}
