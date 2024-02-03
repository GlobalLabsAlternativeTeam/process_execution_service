package main

import (
	"context"
	"fmt"
	"log"
	"net"
	process_execution_service "server/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	process_execution_service.UnimplementedProcessExecutionServiceServer
}

func (s *server) GetTreatemtsByPatientID(
	ctx context.Context, req *process_execution_service.GetTreatemtsByPatientIDRequest,
) (*process_execution_service.GetTreatemtsByPatientIDResponse, error) {
	fmt.Println("dd")
	return nil, nil
}

func (s *server) GetTreatmentByID(
	ctx context.Context, req *process_execution_service.GetTreatmentByIDRequest,
) (*process_execution_service.GetTreatmentByIDResponse, error) {
	fmt.Println("dd")
	return nil, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	process_execution_service.RegisterProcessExecutionServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// func (s *server) GetBookList(ctx context.Context, in *process_execution_service.GetBookListRequest) (*process_execution_service.GetBookListResponse, error) {
// 	return &process_execution_service.GetBookListResponse{
// 		Books: getSampleBooks(),
// 	}, nil
// }

//	func getSampleBooks() []*process_execution_service.Book {
//		sampleBooks := []*process_execution_service.Book{
//			{
//				Title:     "The Hitchhiker's Guide to the Galaxy",
//				Author:    "Douglas Adams",
//				PageCount: 42,
//			},
//			{
//				Title:     "The Lord of the Rings",
//				Author:    "J.R.R. Tolkien",
//				PageCount: 1234,
//			},
//		}
//		return sampleBooks
//	}
