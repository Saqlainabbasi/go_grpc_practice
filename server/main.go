package main

import (
	"log"
	"net"

	pb "github.com/Saqlainabbasi/go_grcp_practice/proto"

	"google.golang.org/grpc"
)

// define a Port
const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	//declare a listner variable
	//use the net package to create listner
	//take two params network type and port/address
	listener, err := net.Listen("tcp", port)
	//check for the error
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	//create a grpc server
	grpcServer := grpc.NewServer()
	//use the pb file form the proto folder
	// pass the grpc server and service
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	//log the server starting message
	log.Printf("Server started at %v", listener.Addr())
	//start the grpc server
	//check for the err if any
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
