package storage_test

import (
	"server/internal/domain"
	"server/internal/providers/storage"
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
	storageService, err := storage.NewStorage("./test_storage.json")
	if err != nil {
		t.Errorf("Failed to create storage: %v", err)
	}

	t.Run("Patient with treatments", func(t *testing.T) {
		treatments, err := storageService.GetTreatments("pat002")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(treatments) != 2 {
			t.Errorf("Expected 2 treatments, got %d", len(treatments))
		}
	})

	t.Run("Patient with no treatments", func(t *testing.T) {
		treatments, err := storageService.GetTreatments("invalidPatient")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(treatments) != 0 {
			t.Errorf("Expected 0 treatments, got %d", len(treatments))
		}
	})
}

func TestTreatmentByID(t *testing.T) {
	storageService, err := storage.NewStorage("./test_storage.json")
	if err != nil {
		t.Errorf("Failed to create storage: %v", err)
	}

	t.Run("Valid treatment", func(t *testing.T) {
		var treatmentId string = "treatment001"
		treatment, err := storageService.TreatmentByID(treatmentId)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if treatment.TreatmentID != treatmentId {
			t.Errorf("Expected treatment with ID '%s', got '%s'", treatmentId, treatment.TreatmentID)
		}
	})

	t.Run("Invalid treatment", func(t *testing.T) {
		_, err := storageService.TreatmentByID("invalidTreatment")
		if err == nil {
			t.Errorf("Expected error because of invalid treatment, got nil")
		}
	})
}

func TestGetPatientsByDoctor(t *testing.T) {
	storageService, err := storage.NewStorage("./test_storage.json")
	if err != nil {
		t.Errorf("Failed to create storage: %v", err)
	}

	t.Run("Doctor with patients", func(t *testing.T) {
		patients, err := storageService.GetPatientsByDoctor("doc001")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(patients) != 2 {
			t.Errorf("Expected 2 patients, got %d", len(patients))
		}
	})

	t.Run("Doctor without patients", func(t *testing.T) {
		patients, err := storageService.GetPatientsByDoctor("doc999")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(patients) != 0 {
			t.Errorf("Expected 0 patients, got %d", len(patients))
		}
	})
}
