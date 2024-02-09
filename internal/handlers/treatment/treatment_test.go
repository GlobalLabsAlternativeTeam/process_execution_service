package treatment_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"server/internal/domain"

	"server/internal/handlers/treatment"
)

type MockStorageProvider struct{}

func (msp *MockStorageProvider) GetTreatments(patientID string) ([]domain.LightTreatment, error) {
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

func TestGetTreatments(t *testing.T) {
	mockStorageProvider := &MockStorageProvider{}
	treatmentService := &treatment.Treatment{StorageProvider: mockStorageProvider}

	t.Run("ValidPatientID", func(t *testing.T) {
		expectedTreatments := []domain.LightTreatment{
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
		resultTreatments, err := treatmentService.PatientTreatments("validPatientID")

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if !reflect.DeepEqual(resultTreatments, expectedTreatments) {
			t.Errorf("Expected treatments %+v, got %+v", expectedTreatments, resultTreatments)
		}
	})

	t.Run("InvalidPatientID", func(t *testing.T) {
		_, err := treatmentService.PatientTreatments("invalidPatientID")

		if err == nil {
			t.Error("Expected an error, but got nil")
		}
	})
}

func (msp *MockStorageProvider) TreatmentByID(treatmentID string) (domain.Treatment, error) {
	if treatmentID == "validID" {
		task := domain.Task{
			ID:          1,
			Level:       1,
			Name:        "Test Task",
			Status:      "Pending",
			BlockedBy:   []interface{}{},
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
			TreatmentID:     "validID",
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

func TestGetTreatment(t *testing.T) {
	mockStorageProvider := &MockStorageProvider{}
	treatmentService := &treatment.Treatment{StorageProvider: mockStorageProvider}

	t.Run("ValidTreatmentID", func(t *testing.T) {
		expectedTreatment := domain.Treatment{
			TreatmentID: "validID",
			DoctorID:    "doctorID",
			PatientID:   "patientID",
			Status:      "InProgress",
			StartedAt:   "2024-02-08",
			FinishedAt:  "",
			DeletedAt:   "",
			PatternInstance: domain.PatternInstance{
				SchemaInstanceID:      "schemaID",
				SchemaID:              "schemaID",
				AuthorID:              "authorID",
				SchemaName:            "Test Schema",
				PatternInstanceStatus: "Active",
				CreatedAt:             time.Now(),
				UpdatedAt:             time.Now(),
				DeletedAt:             time.Time{},
				Tasks: []domain.Task{
					{
						ID:          1,
						Level:       1,
						Name:        "Test Task",
						Status:      "Pending",
						BlockedBy:   []interface{}{},
						Responsible: "John Doe",
						TimeLimit:   time.Now().Unix(),
						Children:    []domain.Task{},
						Comment: struct {
							Value string `json:"value"`
						}{
							Value: "Test comment",
						},
					},
				},
			},
		}
		resultTreatment, err := treatmentService.GetTreatment("validID")

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if !reflect.DeepEqual(resultTreatment, expectedTreatment) {
			t.Errorf("Expected treatment %+v, got %+v", expectedTreatment, resultTreatment)
		}
	})

	t.Run("InvalidTreatmentID", func(t *testing.T) {
		_, err := treatmentService.GetTreatment("invalidID")

		if err == nil {
			t.Error("Expected an error, but got nil")
		}
	})
}
