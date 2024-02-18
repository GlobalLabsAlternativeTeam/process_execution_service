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
	DoctorPatients(doctorID string) ([]string, error)
	CreateTreatment(doctorID string, patientID string, schema domain.Schema) (domain.Treatment, error)
	CompleteTasks(treatmentID string, taskIDs []int64) []domain.TaskLight
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

	fmt.Println(domain.TreatmentToGRPC(&treatment))

	response := &process_execution_service.GetTreatmentByIDResponse{
		Treatment: domain.TreatmentToGRPC(&treatment),
	}
	fmt.Println("END GetTreatmentByID API")
	return response, nil
}

func (s *TreatmentServer) GetPatientsByDoctorID(
	ctx context.Context, req *process_execution_service.GetPatientsByDoctorIDRequest,
) (*process_execution_service.GetPatientsByDoctorIDResponse, error) {
	fmt.Println("START GetPatientsByDoctorID API")
	doctorID := req.DoctorId
	patients, err := s.TreatmentHandler.DoctorPatients(doctorID)

	if err != nil {
		fmt.Println("Error calling treatment API, GetPatientsByDoctorID ", err)
		return nil, err
	}

	response := &process_execution_service.GetPatientsByDoctorIDResponse{
		DoctorId:   req.DoctorId,
		PatientIds: patients,
	}
	fmt.Println("END GetPatientsByDoctorID API")

	return response, nil
}

func (s *TreatmentServer) CreateTreatment(ctx context.Context, req *process_execution_service.CreateTreatmentRequest,
) (*process_execution_service.CreateTreatmentResponse, error) {
	fmt.Println("START CreateTreatment API")
	doctorID := req.DoctorId
	patientID := req.PatientId
	proto_schema := req.Schema

	schema := domain.ProtoToSchema(proto_schema)

	treatment, err := s.TreatmentHandler.CreateTreatment(doctorID, patientID, schema)

	if err != nil {
		fmt.Println("Error calling treatment API, CreateInstance ", err)
		return nil, err
	}

	response := &process_execution_service.CreateTreatmentResponse{
		Treatment: domain.TreatmentToGRPC(&treatment),
	}

	fmt.Println("END CreateTreatment API")
	return response, nil
}

func (s *TreatmentServer) CompleteTasks(ctx context.Context, req *process_execution_service.CompleteTasksRequest,
) (*process_execution_service.CompleteTasksResponse, error) {
	fmt.Println("START CompleteTasks API")
	treatmentID := req.InstanceId
	tasks := req.TaskIds

	result := s.TreatmentHandler.CompleteTasks(treatmentID, tasks)

	var grpcTasksLight []*process_execution_service.TaskLight
	for _, tl := range result {
		grpcTasksLight = append(grpcTasksLight, &process_execution_service.TaskLight{
			TaskId: int64(tl.TaskID),
			Status: domain.ConvertTaskStatus(tl.Status),
		})
	}

	response := &process_execution_service.CompleteTasksResponse{
		TasksLight: grpcTasksLight,
	}
	fmt.Println("END CompleteTasks API")
	return response, nil

}
