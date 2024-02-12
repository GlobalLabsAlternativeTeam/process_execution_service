// internal/providers/storage/storage.go

package storage

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"server/internal/domain"
)

type Storage struct {
	// Add necessary dependencies
}

func (s *Storage) GetTreatments(patientID string) ([]domain.LightTreatment, error) {
	fmt.Println(" START GetTreatments, provider/storage")
	var treatments []domain.LightTreatment
	for i := 0; i < rand.Intn(10)+1; i++ {
		var lightTreatment domain.LightTreatment
		curLightTreatment, err := GenerateRandomLightTreatmentJSON()
		if err != nil {
			fmt.Println("GetTreatments, provider/storage Error generating random LightTreatment JSON")
			fmt.Println(err)
			continue
		}

		err = json.Unmarshal([]byte(curLightTreatment), &lightTreatment)
		if err != nil {
			fmt.Println("GetTreatments, provider/storage Error unmarshaling LightTreatment")
			fmt.Println(err)
			continue
		}

		treatments = append(treatments, lightTreatment)
	}

	fmt.Println("END GetTreatments, provider/storage ")

	return treatments, nil
}

func (s *Storage) TreatmentByID(treatmentID string) (domain.Treatment, error) {
	fmt.Println(" START TreatmentByID, provider/storage")
	var treatment domain.Treatment
	json_treatment, err := GenerateRandomTreatment(treatmentID)
	if err != nil {
		fmt.Println("TreatmentByID, provider/storage Error generating random Treatment JSON")
		fmt.Println(err)
	}
	err = json.Unmarshal([]byte(json_treatment), &treatment)
	if err != nil {
		fmt.Println("TreatmentByID, provider/storage Error unmarshaling Treatment")
		fmt.Println(err)
	}
	fmt.Println(" END TreatmentByID, provider/storage")
	return treatment, nil
}

func (s *Storage) GetPatientsByDoctor(doctorID string) ([]string, error) {
	fmt.Println(" START GetPatientsByDoctor, provider/storage")
	var patients []string

	for i := 0; i < rand.Intn(10)+1; i++ {
		var patient = GenerateRandomString(6)
		patients = append(patients, patient)
	}

	fmt.Println("END GetPatientsByDoctor, provider/storage ")

	return patients, nil
}
