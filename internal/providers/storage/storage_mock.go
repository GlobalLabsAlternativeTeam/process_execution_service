package storage

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"server/internal/domain"
	"time"
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
func GenerateRandomTreatment() domain.Treatment {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate random strings for IDs and statuses
	randomString := func() string {
		const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		b := make([]byte, 10)
		for i := range b {
			b[i] = charset[rand.Intn(len(charset))]
		}
		return string(b)
	}

	// Generate random treatment
	treatment := domain.Treatment{
		TreatmentdID: randomString(),
		DoctorID:     randomString(),
		PatientID:    randomString(),
		Status:       randomString(),
		StartedAt:    time.Now().String(), // Replace with actual logic to generate time strings
		FinishedAt:   time.Now().String(),
		DeletedAt:    time.Now().String(),
		PatternInstance: domain.PatternInstance{
			SchemaInstanceID:      randomString(),
			SchemaID:              randomString(),
			AuthorID:              randomString(),
			SchemaName:            randomString(),
			PatternInstanceStatus: randomString(),
			CreatedAt:             randTime(),
			UpdatedAt:             randTime(),
			DeletedAt:             randTime(),
			Tasks:                 nil, // Populate tasks if needed
		},
	}

	return treatment
}

func randTime() time.Time {
	min := time.Date(2000, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
