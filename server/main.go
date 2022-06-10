package main

import (
	"context"
	"go-grpc-sum-number/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
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
	if request.Number1 == 0 || request.Number2 == 0 {
		return nil, status.Error(codes.InvalidArgument, "tidak bisa dibagi dengan angka 0")
	}
	return &calculator.DevidedByResponse{
		Result: request.Number1 / request.Number2,
	}, nil
}
func (server *CalculatorServer) Fibonacci(request *calculator.FibonacciRequest, stream calculator.CalculatorService_FibonacciServer) error {
	a := 0
	b := 1
	sum := 0
	for i := 0; i < int(request.Value); i++ {
		sum = a + b
		stream.Send(&calculator.FibonacciResponse{
			Value: int64(sum),
		})
		a = b
		b = sum
	}
	return nil
}

func (server *CalculatorServer) AVG(stream calculator.CalculatorService_AVGServer) error {

	var index, total float32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calculator.AVGResponse{
				Result: total / index,
			})
		}
		if err != nil {
			return err
		}

		index++
		total += req.Number

	}
}

func (server *CalculatorServer) FilterSameNumber(stream calculator.CalculatorService_FilterSameNumberServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = stream.Send(&calculator.FilterSumNumberResponse{
			Number: req.Number,
		})
		if err != nil {
			return err
		}

	}

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
