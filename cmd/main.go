// cmd/main.go

package main

import (
	"log"
	"net"
	"server/internal/api"
	"server/internal/handlers/treatment"
	"server/internal/providers/storage"
	process_execution_service "server/proto"

	// "github.com/GlobalLabsAlternativeTeam/process_execution_service/internal/api"
	// "github.com/GlobalLabsAlternativeTeam/process_execution_service/internal/handlers/treatment"
	// "github.com/GlobalLabsAlternativeTeam/process_execution_service/internal/providers/storage"

	"google.golang.org/grpc"
)

func main() {
	// Create a listener on TCP port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create instances of your dependencies (handlers, storage, etc.)
	storageService := &storage.Storage{}
	treatmentHandler := &treatment.Treatment{StorageProvider: storageService}
	apiService := &api.TreatmentServer{TreatmentHandler: treatmentHandler}

	// Create a new gRPC server
	server := grpc.NewServer()

	// Register the ProcessExecutionService server
	process_execution_service.RegisterProcessExecutionServiceServer(server, apiService)

	// Serve and listen for incoming requests
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
