package storage

import (
	"server/internal/domain"
	"testing"
)

type MockStorageProvider struct{}

func (msp *MockStorageProvider) GetTreatments(patientID string) ([]domain.LightTreatment, error) {
	// Simulate returning a list of light treatments for testing purposes
	return []domain.LightTreatment{
		{TreatmentID: "1", TreatmentName: "Treatment 1", TreatmentStatus: "Completed", TreatmentProgress: 100},
		{TreatmentID: "2", TreatmentName: "Treatment 2", TreatmentStatus: "InProgress", TreatmentProgress: 50},
	}, nil
}

func (msp *MockStorageProvider) TreatmentByID(treatmentID string) (domain.Treatment, error) {
	// Simulate returning a treatment by ID for testing purposes
	return domain.Treatment{
		TreatmentID: "1",
		DoctorID:    "doctorID",
		PatientID:   "patientID",
		Status:      "InProgress",
		StartedAt:   "2024-02-08",
		FinishedAt:  "",
		DeletedAt:   "",
		// Add other necessary fields
	}, nil
}

func (msp *MockStorageProvider) GetPatientsByDoctor(doctorID string) ([]string, error) {
	// Simulate returning a list of patient ids for testing purposes
	return []string{
		"asdasd",
		"xyzxyz",
		"asd123",
	}, nil
}

func TestGetTreatments(t *testing.T) {
	mockStorage := &MockStorageProvider{}

	treatments, err := mockStorage.GetTreatments("testPatientID")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(treatments) != 2 {
		t.Errorf("Expected 2 treatments, got %d", len(treatments))
	}
}

func TestTreatmentByID(t *testing.T) {
	mockStorage := &MockStorageProvider{}

	treatment, err := mockStorage.TreatmentByID("testTreatmentID")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if treatment.TreatmentID != "1" {
		t.Errorf("Expected treatment with ID '1', got '%s'", treatment.TreatmentID)
	}
}

func TestGetPatientsByDoctor(t *testing.T) {
	mockStorage := &MockStorageProvider{}

	patients, err := mockStorage.GetPatientsByDoctor("testDoctorID")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(patients) != 3 {
		t.Errorf("Expected 3 patients, got %d", len(patients))
	}
}
