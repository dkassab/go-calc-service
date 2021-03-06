package main

import (
	"flag"
	pb "github.com/dkassab/go-calc-service/calc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":33333"
)

// CalculatorServer is used to implement Calculator.CalcServer.
type calculatorServer struct{}

// Add implements Calculator.Add
func (s *calculatorServer) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddReply, error) {
	return &pb.AddReply{Result: in.Value1 + in.Value2}, nil
}

// Subtract implements Calculator.Subtract
func (s *calculatorServer) Subtract(ctx context.Context, in *pb.SubtractRequest) (*pb.SubtractReply, error) {
	return &pb.SubtractReply{Result: in.Value1 - in.Value2}, nil
}

// Multiply implements Calculator.Multiply
func (s *calculatorServer) Multiply(ctx context.Context, in *pb.MultiplyRequest) (*pb.MultiplyReply, error) {
	return &pb.MultiplyReply{Result: in.Value1 * in.Value2}, nil
}

// Divide implements Calculator.Divide
func (s *calculatorServer) Divide(ctx context.Context, in *pb.DivideRequest) (*pb.DivideReply, error) {
	return &pb.DivideReply{Result: float32(in.Value1) / float32(in.Value2)}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", ":33333")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &calculatorServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
