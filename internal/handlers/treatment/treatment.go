package treatment

import (
	"fmt"
	"server/internal/domain"
)

// import {

// }

type StorageInterface interface {
	GetTreatemtsByPatientID(patient_id string) (domain.TreatmentLight, error)
	GetTreatment(treatment_id string) (domain.Treatment, error)
}

type TreatmentHandler struct {
	storageProvider StorageInterface
}

func (t *TreatmentHandler) GetTreatemtsByPatientID(patient_id string) (domain.TreatmentLight, error) {
	treatmnents, err := t.storageProvider.GetTreatemtsByPatientID(patient_id)
	if err != nil {
		fmt.Println("Error getting treatment entities: ", err)
	}
	fmt.Printf("GetTreatemtsByPatientID, provider/handler")
	return treatmnents, nil
}

func (t *TreatmentHandler) GetTreatments(treatment_id string) (domain.Treatment, error) {
	treatmemt, err := t.storageProvider.GetTreatment(treatment_id)
	if err != nil {
		fmt.Println("Error getting treatment entities: ", err)
	}
	fmt.Printf("GetTreatemts, provider/handler")
	return treatmemt, nil
}
