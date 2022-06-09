package main

import (
	"context"
	"go-grpc-sum-number/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := calculator.NewCalculatorServiceClient(conn)

	ctx := context.Background()

	result1, err := client.Sum(ctx, &calculator.SumRequest{
		Number1: 6,
		Number2: 4,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result1)
	result2, err := client.Sum(ctx, &calculator.SumRequest{
		Number1: 7,
		Number2: 13,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result2)
	result3, err := client.DevidedBy(ctx, &calculator.DevidedByRequest{
		Number1: 20,
		Number2: 4,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result3)
}
