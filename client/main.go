package main

import (
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

	//doSum(client)
	//doDevidedBy(client)
	//doFibonacci(client)
	//doAVG(client)
	doFilterDuplicateNumber(client)
}
