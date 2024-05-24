package main

import (
	"context"
	"grpc-compute/compute"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error in connecting to server: %v", err)
	}

	defer conn.Close()

	c := compute.NewComputeClient(conn)

	response1, err := c.Square(context.Background(), &compute.RequestBody{
		X: 3,
	})

	if err != nil {
		log.Fatalf("Error in getting response from square: %v", err)
	}

	log.Println("Square of number is ", response1.GetX())

	response2, err := c.Cube(context.Background(), &compute.RequestBody{
		X: 2,
	})

	if err != nil {
		log.Fatalf("Error in getting the response from cube: %v", err)
	}

	log.Println("Cube of number is ", response2.GetX())
}
