package storage_test

import (
	"server/internal/providers/storage"
	"testing"
)

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
