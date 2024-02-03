package api

import (
	"fmt"
	"server/internal/domain"
	process_execution_service "server/proto"
)

type TreatmentHandeler interface {
	GetTreatemtsByPatientID(patient_id string) (domain.TreatmentLight, error)
	GetTreatmentByID(treatment_id string) (domain.Treatment, error)
}

type TreatmentAPI struct {
	treatmentHandeler TreatmentHandeler
}

func (t *TreatmentAPI) GetTreatemtsByPatientID(
	req *process_execution_service.GetTreatemtsByPatientIDRequest,
) (*process_execution_service.GetTreatemtsByPatientIDResponse, error) {
	var patient_id = "aa"
	_, err := t.treatmentHandeler.GetTreatemtsByPatientID(patient_id)
	if err != nil {
		fmt.Println("Error getting treatment entities: ", err)
	}

	return nil, nil
}

func (t *TreatmentAPI) GetTreatmentByID(
	req *process_execution_service.GetTreatmentByIDRequest,
) (*process_execution_service.GetTreatmentByIDResponse, error) {
	var treatment_id = "aa"
	_, err := t.treatmentHandeler.GetTreatmentByID(treatment_id)
	if err != nil {
		fmt.Println("Error getting treatment entity: ", err)
	}

	return nil, nil
}
