package storage

import (
	"fmt"
	"server/internal/domain"
)

type Storage struct {
	StorageServer domain.StorageServer
}

func (s *Storage) GetTreatemtsByPatientID(patient_id string) (domain.TreatmentLight, error) {
	fmt.Printf("GetTreatemtsByPatientID, provider/storage")
	return domain.TreatmentLight{}, nil
}

func (s *Storage) GetTreatment(treatment_id string) (domain.Treatment, error) {
	fmt.Printf("GetTreatment, provider/storage")
	return domain.Treatment{}, nil
}
