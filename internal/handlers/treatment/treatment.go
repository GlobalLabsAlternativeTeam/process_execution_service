// internal/handlers/treatment/treatment.go

package treatment

import (
	"fmt"
	"server/internal/domain"
)

type StorageInterface interface {
	GetTreatments(patientID string) ([]domain.LightTreatment, error)
	TreatmentByID(treatmentID string) (domain.Treatment, error)
}

type Treatment struct {
	StorageProvider StorageInterface
}

func (t *Treatment) PatientTreatments(patientID string) ([]domain.LightTreatment, error) {
	fmt.Println("START PatientTreatments handler")
	treatments, err := t.StorageProvider.GetTreatments(patientID)
	if err != nil {
		fmt.Println("Error getting treatment entities: ", err)
	}
	fmt.Println("END PatientTreatments handler")
	return treatments, nil
}

func (t *Treatment) GetTreatment(treatmentID string) (domain.Treatment, error) {
	treatment, err := t.StorageProvider.TreatmentByID(treatmentID)
	if err != nil {
		fmt.Println("Error getting treatment entities: ", err)
	}
	fmt.Printf("GetTreatemts, provider/handler")
	return treatment, nil
}
