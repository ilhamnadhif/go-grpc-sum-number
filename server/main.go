package main

import (
	"context"
	"go-grpc-sum-number/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type CalculatorServer struct {
	proto.CalculatorServer
}

func (server *CalculatorServer) Sum(ctx context.Context, request *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	return &proto.CalculatorResponse{
		Result: request.Number1 + request.Number2,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalln(err.Error())
	}

	grpcServer := grpc.NewServer()
	proto.RegisterCalculatorServer(grpcServer, &CalculatorServer{})
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln(err.Error())
	}
}
