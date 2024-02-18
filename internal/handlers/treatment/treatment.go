// internal/handlers/treatment/treatment.go

package treatment

import (
	"fmt"
	"server/internal/domain"
	"time"

	"github.com/google/uuid"
)

type StorageInterface interface {
	GetTreatments(patientID string) ([]domain.LightTreatment, error)
	TreatmentByID(treatmentID string) (domain.Treatment, error)
	GetPatientsByDoctor(doctorID string) ([]string, error)
	CreateTreatment(treatmentID string, doctorID string,
		patientID string, patternInstance domain.PatternInstance) (domain.Treatment, error)
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

func (t *Treatment) CreateTreatment(doctorID string, patientID string, schema domain.Schema) (domain.Treatment, error) {
	fmt.Println("START CreateTreatment handler")
	patternInstance := t.CreateInstance(schema)
	treatmentStatus := "RUNNING"
	treatment, err := t.StorageProvider.CreateTreatment(doctorID, patientID, treatmentStatus, patternInstance)
	if err != nil {
		fmt.Println("Error createing treatment: ", err)
		return domain.Treatment{}, err
	}
	fmt.Println("END CreateTreatment handler")
	return treatment, nil

}

func (t *Treatment) CreateInstance(schema domain.Schema) domain.PatternInstance {

	//TODO: Check that there is no sucj id for another instance in the same treatment (alsmost impossible)
	id := uuid.New().String()

	patternInstance := domain.PatternInstance{
		SchemaInstanceID:      id,
		SchemaID:              schema.SchemaID,
		AuthorID:              schema.AuthorID,
		SchemaName:            schema.SchemaName,
		PatternInstanceStatus: "NOT_STARTED",
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
		DeletedAt:             time.Time{},
		Tasks:                 schema.Tasks,
	}
	return patternInstance

}
