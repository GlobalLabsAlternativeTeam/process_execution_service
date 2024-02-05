package storage

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

func LightTreatmentMock() []byte {
	lightTreatment := struct {
		Treatment_ID       string `json:"treatment_id"`
		Treatment_Name     string `json:"treatment_name"`
		Treatment_Status   string `json:"treatment_status"`
		Treatment_Progress int    `json:"treatment_progress"`
	}{
		Treatment_ID:       "",
		Treatment_Name:     "",
		Treatment_Status:   "",
		Treatment_Progress: rand.Intn(100),
	}

	result, err := json.Marshal(lightTreatment)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil
	}

	return result
}

func TreatmentMock() []byte {
	lightTreament := []byte(`{

	}`)
	return lightTreament
}
