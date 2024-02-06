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
	treatment := GenerateRandomTreatment(treatmentID)
	fmt.Printf("GetTreatment, provider/storage")
	return treatment, nil
}
