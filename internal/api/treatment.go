// internal/api/api.go

package api

import (
	"context"
	"fmt"
	"server/internal/domain"
	process_execution_service "server/proto"
)

type TreatmentHandler interface {
	PatientTreatments(patientID string) ([]domain.TreatmentLight, error)
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
	patientID := "aa"
	treatments, err := s.TreatmentHandler.PatientTreatments(patientID)

	if err != nil {
		fmt.Println("Error calling treatment API, GetTreatmentsByPatientID ", err)
		return nil, err
	}
	var grpcTreatments []*process_execution_service.TreatmentLight
	for _, lt := range treatments {
		grpcTreatments = append(grpcTreatments, &process_execution_service.TreatmentLight{
			TreatmentId:       lt.Treatment_ID,
			TreatmentName:     lt.Treatment_Name,
			TreatmentStatus:   lt.Treatment_Status,
			TreatmentProgress: float32(lt.Treatment_Progress),
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
	treatmentID := "aa"
	_, err := s.TreatmentHandler.GetTreatment(treatmentID)
	if err != nil {
		fmt.Println("Error calling treatment API, GetTreatmentByID ", err)
		return nil, err
	}
	return nil, nil
}
