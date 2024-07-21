package main

import (
	"calculator/calculatorpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func CallSum(firstNumber, secondNumber int32, c calculatorpb.CalculatorClient, ctx context.Context) {
	request := &calculatorpb.SumRequest{
		FirstNumber:  firstNumber,
		SecondNumber: secondNumber,
	}

	response, err := c.Sum(ctx, request)
	if err != nil {
		log.Fatalf("could not sum: %v", err)
	}
	fmt.Printf("Sum Result: %v\n", response.SumResult)
}

func CallSub(firstNumber, secondNumber int32, c calculatorpb.CalculatorClient, ctx context.Context) {
	request := &calculatorpb.SubRequest{
		FirstNumber:  firstNumber,
		SecondNumber: secondNumber,
	}
	response, err := c.Sub(ctx, request)
	if err != nil {
		log.Fatalf("could not sub: %v", err)
	}
	fmt.Printf("Sub Result: %v\n", response.SubResult)
}

func CallMul(firstNumber, secondNumber int32, c calculatorpb.CalculatorClient, ctx context.Context) {
	request := &calculatorpb.MulRequest{
		FirstNumber:  firstNumber,
		SecondNumber: secondNumber,
	}
	response, err := c.Mul(ctx, request)
	if err != nil {
		log.Fatalf("could not mul: %v", err)
	}
	fmt.Printf("Mul Result: %v\n", response.MulResult)
}

func CallDiv(firstNumber, secondNumber int32, c calculatorpb.CalculatorClient, ctx context.Context) {
	request := &calculatorpb.DivRequest{
		FirstNumber:  firstNumber,
		SecondNumber: secondNumber,
	}

	response, err := c.Div(ctx, request)
	if err != nil {
		log.Fatalf("could not div: %v", err)
	}
	fmt.Printf("Div Result: %v\n", response.DivResult)
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient("localhost:8080", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	c := calculatorpb.NewCalculatorClient(conn)
	ctx := context.Background()

	CallSum(1, 2, c, ctx)
	CallSub(10, 2, c, ctx)
	CallMul(10, 20, c, ctx)
	CallDiv(15, 3, c, ctx)
}
