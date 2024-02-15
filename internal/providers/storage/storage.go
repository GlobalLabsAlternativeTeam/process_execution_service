// internal/providers/storage/storage.go

package storage

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"server/internal/domain"
)

type Storage struct {
	filePath   string
	treatments map[string]domain.Treatment
}

func NewStorage(filePath string) (*Storage, error) {
	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// If the file doesn't exist, create an empty JSON file
		if err := createEmptyJSONFile(filePath); err != nil {
			return nil, fmt.Errorf("error creating storage file: %v", err)
		}
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading storage file: %v", err)
	}

	var treatments []domain.Treatment
	err = json.Unmarshal(data, &treatments)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	treatmentsMap := make(map[string]domain.Treatment)
	for _, treatment := range treatments {
		treatmentsMap[treatment.TreatmentID] = treatment
	}

	return &Storage{
		filePath:   filePath,
		treatments: treatmentsMap,
	}, nil
}

func (s *Storage) SaveToFile() error {
	treatmentsSlice := make([]domain.Treatment, 0, len(s.treatments))
	for _, treatment := range s.treatments {
		treatmentsSlice = append(treatmentsSlice, treatment)
	}

	data, err := json.MarshalIndent(treatmentsSlice, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshalling treatments to JSON: %v", err)
	}

	err = os.WriteFile(s.filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing data to file: %v", err)
	}

	return nil
}

func createEmptyJSONFile(filePath string) error {
	emptyData := []byte("[]") // JSON representation for empty array
	err := os.WriteFile(filePath, emptyData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) PrintTreatments() {
	for id, treatment := range s.treatments {
		fmt.Printf("ID: %s\n", id)
		fmt.Printf("Doctor ID: %s\n", treatment.DoctorID)
		fmt.Printf("Patient ID: %s\n", treatment.PatientID)
		fmt.Printf("Status: %s\n", treatment.Status)
		fmt.Printf("Started At: %s\n", treatment.StartedAt)
		fmt.Printf("Finished At: %s\n", treatment.FinishedAt)
		fmt.Printf("Deleted At: %s\n", treatment.DeletedAt)
		fmt.Printf("Pattern Instance Author ID: %s\n", treatment.PatternInstance.AuthorID)
		fmt.Printf("Pattern Instance Created At: %s\n", treatment.PatternInstance.CreatedAt)
		fmt.Printf("Pattern Instance Updated At: %s\n", treatment.PatternInstance.UpdatedAt)
		fmt.Printf("Pattern Instance Deleted At: %s\n", treatment.PatternInstance.DeletedAt)

		fmt.Println("Tasks:")
		for _, task := range treatment.PatternInstance.Tasks {
			fmt.Printf("Task ID: %d\n", task.ID)
			fmt.Printf("Task Level: %d\n", task.Level)
			fmt.Printf("Task Name: %s\n", task.Name)
			fmt.Printf("Task Status: %s\n", task.Status)
			fmt.Printf("Task Blocked By: %v\n", task.BlockedBy)
			fmt.Printf("Task Responsible: %s\n", task.Responsible)
			fmt.Printf("Task Time Limit: %d\n", task.TimeLimit)
			fmt.Printf("Task Children: %v\n", task.Children)
			fmt.Printf("Task Comment: %s\n", task.Comment.Value)
			fmt.Println("-----------------------")
		}

		fmt.Println("=======================")
	}
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
	fmt.Printf("START TreatmentByID, provider/storage with ID: %s\n", treatmentID)
	// Search treatment in our map
	treatment, ok := s.treatments[treatmentID]
	if !ok {
		// If we don't find it, return error with empty treatment
		return domain.Treatment{}, fmt.Errorf("treatment with ID %s not found", treatmentID)
	}
	fmt.Println("END TreatmentByID, provider/storage")
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
