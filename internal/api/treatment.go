package api

import (
	"context"
	"fmt"
	"server/internal/domain"
	process_execution_service "server/proto"
)

type TreatmentHandler interface {
	PatientTreatments(patientID string) ([]domain.LightTreatment, error)
	GetTreatment(treatmentID string) (domain.Treatment, error)
}

type TreatmentServer struct {
	process_execution_service.UnimplementedProcessExecutionServiceServer
	TreatmentHandler TreatmentHandler
}

func (s *TreatmentServer) GetTreatmentsByPatientID(
	ctx context.Context, req *process_execution_service.GetTreatmentsByPatientIDRequest,
) (*process_execution_service.GetTreatmentsByPatientIDResponse, error) {
	fmt.Println("START GetTreatmentsByPatientID API")
	patientID := req.PatientId
	treatments, err := s.TreatmentHandler.PatientTreatments(patientID)

	if err != nil {
		fmt.Println("Error calling treatment API, GetTreatmentsByPatientID ", err)
		return nil, err
	}
	var grpcTreatments []*process_execution_service.TreatmentLight
	for _, lt := range treatments {
		grpcTreatments = append(grpcTreatments, &process_execution_service.TreatmentLight{
			TreatmentId:       lt.TreatmentID,
			TreatmentName:     lt.TreatmentName,
			TreatmentStatus:   lt.TreatmentStatus,
			TreatmentProgress: float32(lt.TreatmentProgress),
		})
	}

	response := &process_execution_service.GetTreatmentsByPatientIDResponse{
		PatientId:      req.PatientId,
		TreatmentLight: grpcTreatments,
	}
	fmt.Println("END GetTreatmentsByPatientID API")
	return response, nil

}

func (s *TreatmentServer) GetTreatmentByID(
	ctx context.Context, req *process_execution_service.GetTreatmentByIDRequest,
) (*process_execution_service.GetTreatmentByIDResponse, error) {
	fmt.Println("START GetTreatmentByID API")
	treatmentID := req.TreatmentId
	treatment, err := s.TreatmentHandler.GetTreatment(treatmentID)
	if err != nil {
		fmt.Println("Error calling treatment API, GetTreatmentByID ", err)
		return nil, err
	}

	response := &process_execution_service.GetTreatmentByIDResponse{
		Treatment: domain.TreatmentToGRPC(&treatment),
	}
	fmt.Println("END GetTreatmentByID API")
	return response, nil
}
