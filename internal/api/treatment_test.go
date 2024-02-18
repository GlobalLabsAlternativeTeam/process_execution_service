package api_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"server/internal/api"
	"server/internal/domain"

	process_execution_service "server/proto"
)

type MockTreatmentHandler struct{}

func (mth *MockTreatmentHandler) PatientTreatments(patientID string) ([]domain.LightTreatment, error) {
	if patientID == "validPatientID" {
		treatments := []domain.LightTreatment{
			{
				TreatmentID:       "treatment1",
				TreatmentName:     "Treatment 1",
				TreatmentStatus:   "Completed",
				TreatmentProgress: 100,
			},
			{
				TreatmentID:       "treatment2",
				TreatmentName:     "Treatment 2",
				TreatmentStatus:   "InProgress",
				TreatmentProgress: 50,
			},
		}
		return treatments, nil
	}
	return nil, errors.New("Patient not found")
}

func (mth *MockTreatmentHandler) GetTreatment(treatmentID string) (domain.Treatment, error) {
	if treatmentID == "validTreatmentID" {
		task := domain.Task{
			ID:          1,
			Level:       1,
			Name:        "Test Task",
			Status:      "Pending",
			BlockedBy:   []int64{},
			Responsible: "John Doe",
			TimeLimit:   time.Now().Unix(),
			Children:    []domain.Task{},
			Comment: struct {
				Value string `json:"value"`
			}{
				Value: "Test comment",
			},
		}
		patternInstance := domain.PatternInstance{
			SchemaInstanceID:      "schemaID",
			SchemaID:              "schemaID",
			AuthorID:              "authorID",
			SchemaName:            "Test Schema",
			PatternInstanceStatus: "Active",
			CreatedAt:             time.Now(),
			UpdatedAt:             time.Now(),
			DeletedAt:             time.Time{},
			Tasks:                 []domain.Task{task},
		}
		return domain.Treatment{
			TreatmentID:     "validTreatmentID",
			DoctorID:        "doctorID",
			PatientID:       "patientID",
			Status:          "InProgress",
			StartedAt:       "2024-02-08",
			FinishedAt:      "",
			DeletedAt:       "",
			PatternInstance: patternInstance,
		}, nil
	}
	return domain.Treatment{}, errors.New("Treatment not found")
}

func (mth *MockTreatmentHandler) DoctorPatients(doctorID string) ([]string, error) {
	if doctorID != "validDoctorID" {
		return nil, errors.New("Doctor not found")
	}

	patients := []string{
		"asdasd",
		"xyzxyz",
		"asd123",
	}
	return patients, nil
}

func (mth *MockTreatmentHandler) CreateTreatment(doctorID string, patientID string, schema domain.Schema) (domain.Treatment, error) {
	return domain.Treatment{}, nil
}

func (mth *MockTreatmentHandler) CompleteTasks(treatmentID string, taskIDs []int64) []domain.TaskLight {
	return []domain.TaskLight{}
}

func TestGetTreatmentByID(t *testing.T) {
	mockHandler := &MockTreatmentHandler{}
	apiHandler := api.TreatmentServer{TreatmentHandler: mockHandler}

	tests := []struct {
		name          string
		request       *process_execution_service.GetTreatmentByIDRequest
		expectedError bool
	}{
		{
			name: "ValidTreatmentID",
			request: &process_execution_service.GetTreatmentByIDRequest{
				TreatmentId: "validTreatmentID",
			},
			expectedError: false,
		},
		{
			name: "InvalidTreatmentID",
			request: &process_execution_service.GetTreatmentByIDRequest{
				TreatmentId: "invalidTreatmentID",
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := apiHandler.GetTreatmentByID(context.Background(), tt.request)

			if (err != nil) != tt.expectedError {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if tt.expectedError {
				return
			}

			if response.Treatment == nil {
				t.Error("Expected non-nil treatment")
				return
			}

			if response.Treatment.TreatmentId != "validTreatmentID" {
				t.Errorf("Expected treatment ID 'validTreatmentID', got '%s'", response.Treatment.TreatmentId)
			}
		})
	}
}

func TestGetDoctorPatients(t *testing.T) {
	mockHandler := &MockTreatmentHandler{}
	apiHandler := api.TreatmentServer{TreatmentHandler: mockHandler}

	tests := []struct {
		name           string
		request        *process_execution_service.GetPatientsByDoctorIDRequest
		expectedError  bool
		expectedLength int
	}{
		{
			name: "ValidDoctorID",
			request: &process_execution_service.GetPatientsByDoctorIDRequest{
				DoctorId: "validDoctorID",
			},
			expectedError:  false,
			expectedLength: 3,
		},
		{
			name: "InvalidDoctorID",
			request: &process_execution_service.GetPatientsByDoctorIDRequest{
				DoctorId: "invalidDoctorID",
			},
			expectedError:  true,
			expectedLength: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := apiHandler.GetPatientsByDoctorID(context.Background(), tt.request)

			if (err != nil) != tt.expectedError {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if tt.expectedError {
				return
			}

			if len(response.PatientIds) != tt.expectedLength {
				t.Errorf("Expected %d patient IDs, got %d", tt.expectedLength, len(response.PatientIds))
			}
		})
	}
}
