// internal/handlers/treatment/treatment.go

package treatment

import (
	"fmt"
	"server/internal/domain"
)

type StorageInterface interface {
	GetTreatments(patientID string) ([]domain.LightTreatment, error)
	TreatmentByID(treatmentID string) (domain.Treatment, error)
	GetPatientsByDoctor(doctorID string) ([]string, error)
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
	return treatments, err
}

func (t *Treatment) GetTreatment(treatmentID string) (domain.Treatment, error) {
	treatment, err := t.StorageProvider.TreatmentByID(treatmentID)
	if err != nil {
		fmt.Println("Error getting treatment entities: ", err)
	}
	fmt.Printf("GetTreatemts, provider/handler")
	return treatment, err
}

func (t *Treatment) DoctorPatients(doctorID string) ([]string, error) {
	fmt.Println("START DoctorPatients handler")
	patients, err := t.StorageProvider.GetPatientsByDoctor(doctorID)
	if err != nil {
		fmt.Println("Error getting patient IDs by doctor: ", err)
		return nil, err
	}
	fmt.Println("END DoctorPatients handler")
	return patients, nil
}
