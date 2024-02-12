package storage

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"server/internal/domain"
	"time"
)

func GenerateRandomLightTreatmentJSON() (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomString := func(length int) string {
		b := make([]byte, length)
		for i := range b {
			b[i] = charset[rand.Intn(len(charset))]
		}
		return string(b)
	}
	statuses := []string{"RUNNING", "BLOCKED", "COMPLETED"}
	status := statuses[rand.Intn(len(statuses))]

	lightTreatment := domain.LightTreatment{
		TreatmentID:       randomString(10),
		TreatmentName:     randomString(10),
		TreatmentStatus:   status,
		TreatmentProgress: rand.Intn(100) + 1,
	}
	fmt.Print(lightTreatment)

	// Marshal the LightTreatment struct to JSON
	jsonBytes, err := json.Marshal(lightTreatment)
	if err != nil {
		return "", err
	}

	// Convert JSON bytes to string
	jsonString := string(jsonBytes)

	return jsonString, nil
}

func GenerateRandomPatternInstance() domain.PatternInstance {
	// Generate random pattern instance ID
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	instanceIDLength := rand.Intn(20) + 5 // Random length between 5 and 25 characters
	instanceIDBytes := make([]byte, instanceIDLength)
	for i := range instanceIDBytes {
		instanceIDBytes[i] = charset[rand.Intn(len(charset))]
	}
	instanceID := string(instanceIDBytes)

	// Generate random schema instance ID
	schemaInstanceID := instanceID

	// Generate random schema ID
	schemaID := instanceID

	// Generate random author ID
	authorID := instanceID

	// Generate random schema name
	schemaName := instanceID

	// Generate random pattern instance status
	patternInstanceStatuses := []string{"NOT_STARTED", "RUNNING", "BLOCKED", "COMPLETED", "CANCELLED"}
	patternInstanceStatus := patternInstanceStatuses[rand.Intn(len(patternInstanceStatuses))]

	// Generate random created at, updated at, and deleted at times
	createdAt := randTime()
	updatedAt := randTime()
	deletedAt := randTime()

	// Generate random tasks
	numTasks := rand.Intn(5) + 1 // Random number of tasks (up to 5)
	tasks := make([]domain.Task, numTasks)
	for i := range tasks {
		tasks[i] = GenerateRandomTask()
	}

	// Construct and return the pattern instance
	return domain.PatternInstance{
		SchemaInstanceID:      schemaInstanceID,
		SchemaID:              schemaID,
		AuthorID:              authorID,
		SchemaName:            schemaName,
		PatternInstanceStatus: patternInstanceStatus,
		CreatedAt:             createdAt,
		UpdatedAt:             updatedAt,
		DeletedAt:             deletedAt,
		Tasks:                 tasks,
	}
}

func GenerateRandomTreatment(treatmentID string) (string, error) {
	// Generate random treatment ID

	// Generate random doctor ID
	doctorID := treatmentID

	// Generate random patient ID
	patientID := treatmentID

	// Generate random treatment status
	statuses := []string{"RUNNING", "BLOCKED", "COMPLETED"}
	status := statuses[rand.Intn(len(statuses))]

	// Generate random started at, finished at, and deleted at times
	startedAt := randTime()
	finishedAt := randTime()
	deletedAt := randTime()

	// Generate random pattern instance
	patternInstance := GenerateRandomPatternInstance()

	// Construct and return the treatment
	treatment := domain.Treatment{
		TreatmentID:     treatmentID,
		DoctorID:        doctorID,
		PatientID:       patientID,
		Status:          status,
		StartedAt:       startedAt.String(),
		FinishedAt:      finishedAt.String(),
		DeletedAt:       deletedAt.String(),
		PatternInstance: patternInstance,
	}

	// Marshal the Treatment struct to JSON
	jsonBytes, err := json.Marshal(treatment)
	if err != nil {
		return "", err
	}

	// Convert JSON bytes to string
	jsonString := string(jsonBytes)

	return jsonString, nil
}

func randTime() time.Time {
	min := time.Date(2000, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func GenerateRandomTask() domain.Task {
	// Generate random task ID
	id := rand.Intn(1000)

	// Generate random task name
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	nameLength := rand.Intn(20) + 5 // Random length between 5 and 25 characters
	nameBytes := make([]byte, nameLength)
	for i := range nameBytes {
		nameBytes[i] = charset[rand.Intn(len(charset))]
	}
	name := string(nameBytes)

	// Generate random task level
	level := rand.Intn(10)

	// Generate random task status
	taskStatuses := []string{"NOT_STARTED", "IN_PROGRESS", "BLOCKED", "DONE"}
	status := taskStatuses[rand.Intn(len(taskStatuses))]

	// Generate random task time limit
	timeLimit := rand.Int63n(24*60) + 1 // Random value between 1 and 1440 minutes (1 day)

	// Generate random task responsible
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(15) + 1
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	responsible := string(result)

	// Generate random task blocked by
	blockedBy := make([]interface{}, rand.Intn(5)) // Random number of blocked tasks
	for i := range blockedBy {
		blockedBy[i] = rand.Intn(1000) // Assuming task IDs are integers
	}

	// Generate random task comment
	const commentCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	commentLength := rand.Intn(50) + 10 // Random length between 10 and 60 characters
	commentBytes := make([]byte, commentLength)
	for i := range commentBytes {
		commentBytes[i] = commentCharset[rand.Intn(len(commentCharset))]
	}
	comment := string(commentBytes)

	// Generate random task children (recursive call to GenerateRandomTask)
	var children []domain.Task
	if rand.Float64() < 0.5 {
		numChildren := rand.Intn(5) // Random number of children tasks (up to 5)
		children = make([]domain.Task, numChildren)
		for i := range children {
			children[i] = GenerateRandomTask()
		}
	}

	// Construct and return the task
	return domain.Task{
		ID:          id,
		Level:       level,
		Name:        name,
		Status:      status,
		BlockedBy:   blockedBy,
		Responsible: responsible,
		TimeLimit:   timeLimit,
		Children:    children,
		Comment: struct {
			Value string `json:"value"`
		}{Value: comment},
	}
}
