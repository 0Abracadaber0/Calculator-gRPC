package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"calculator/calculatorpb" // Adjust path as needed

	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServer // Embed the unimplemented server
}

func (s *server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Println("Sum function was invoked")
	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()
	sum := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResult: sum,
	}
	return res, nil
}

func (s *server) Sub(ctx context.Context, req *calculatorpb.SubRequest) (*calculatorpb.SubResponse, error) {
	fmt.Println("Sub function was invoked")
	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()
	sub := firstNumber - secondNumber
	res := &calculatorpb.SubResponse{
		SubResult: sub,
	}
	return res, nil
}

func (s *server) Mul(ctx context.Context, req *calculatorpb.MulRequest) (*calculatorpb.MulResponse, error) {
	fmt.Println("Mul function was invoked")
	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()
	mul := firstNumber * secondNumber
	res := &calculatorpb.MulResponse{
		MulResult: mul,
	}
	return res, nil
}

func (s *server) Div(ctx context.Context, req *calculatorpb.DivRequest) (*calculatorpb.DivResponse, error) {
	fmt.Println("Div function was invoked")
	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()
	if secondNumber == 0 {
		return nil, fmt.Errorf("cannot divide by zero")
	}
	div := firstNumber / secondNumber
	res := &calculatorpb.DivResponse{
		DivResult: div,
	}
	return res, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServer(s, &server{})
	log.Printf("Server started on port 8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func main() {
	StartServer()
}
