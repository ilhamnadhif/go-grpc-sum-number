package main

import (
	"context"
	"go-grpc-sum-number/calculator"
	"google.golang.org/grpc/status"
	"io"
	"log"
)

func doSum(client calculator.CalculatorServiceClient) {
	result1, err := client.Sum(context.Background(), &calculator.SumRequest{
		Number1: 6,
		Number2: 4,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result1)

}
func doDevidedBy(client calculator.CalculatorServiceClient) {
	result, err := client.DevidedBy(context.Background(), &calculator.DevidedByRequest{
		Number1: 0,
		Number2: 4,
	})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Println(e.Message())
			log.Fatalln(e.Code())
		} else {
			log.Fatalln(err)
		}
	}
	log.Println(result)

}

func doFibonacci(client calculator.CalculatorServiceClient) {
	stream, err := client.Fibonacci(context.Background(), &calculator.FibonacciRequest{
		Value: 5,
	})
	if err != nil {
		log.Fatalln(err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(msg.Value)
	}

}

func doAVG(client calculator.CalculatorServiceClient) {
	stream, err := client.AVG(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	reqs := []*calculator.AVGRequest{
		{Number: 10},
		{Number: 5},
	}
	for _, req := range reqs {
		stream.Send(req)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Result)

}

func doFilterDuplicateNumber(client calculator.CalculatorServiceClient) {
	reqs := []*calculator.FilterSumNumberRequest{
		{Number: 1},
		{Number: 2},
		{Number: 1},
		{Number: 3},
	}
	stream, err := client.FilterSameNumber(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	for _, req := range reqs {
		err := stream.Send(req)
		if err != nil {
			log.Fatalln(err)
		}
	}
	err = stream.CloseSend()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(res.Number)
	}

}
