package main

import (
	"context"
	"go-grpc-sum-number/calculator"
	"google.golang.org/grpc"
	"log"
	"net"
)

type CalculatorServer struct {
	//calculator.CalculatorServer
	calculator.UnimplementedCalculatorServiceServer
}

func (server *CalculatorServer) Sum(ctx context.Context, request *calculator.SumRequest) (*calculator.SumResponse, error) {
	return &calculator.SumResponse{
		Result: request.Number1 + request.Number2,
	}, nil
}
func (server *CalculatorServer) DevidedBy(ctx context.Context, request *calculator.DevidedByRequest) (*calculator.DevidedByResponse, error) {
	return &calculator.DevidedByResponse{
		Result: request.Number1 / request.Number2,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalln(err.Error())
	}

	grpcServer := grpc.NewServer()
	calculator.RegisterCalculatorServiceServer(grpcServer, &CalculatorServer{})
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln(err.Error())
	}
}
