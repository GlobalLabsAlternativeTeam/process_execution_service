package domain

import "time"

type Task struct {
	ID          int    `json:"id"`
	Level       int    `json:"level"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	BlockedBy   []any  `json:"blocked_by"`
	Responsible string
	TimeLimit   int64  `json:"time_limit"`
	Children    []Task `json:"children"`
	Comment     struct {
		Value string `json:"value"`
	} `json:"comment"`
}

type PatternInstance struct {
	SchemaInstanceID      string    `json:"schema_instance_id"`
	SchemaID              string    `json:"schema_id"`
	AuthorID              string    `json:"author_id"`
	SchemaName            string    `json:"schema_name"`
	PatternInstanceStatus string    `json:"PatternInstanceStatus"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	DeletedAt             time.Time `json:"deleted_at"`
	Tasks                 []Task    `json:"tasks"`
}
type Treatment struct {
	TreatmentID     string          `json:"treatmentd_id"`
	DoctorID        string          `json:"doctor_id"`
	PatientID       string          `json:"patient_id"`
	Status          string          `json:"status"`
	StartedAt       string          `json:"started_at"`
	FinishedAt      string          `json:"finished_at"`
	DeletedAt       string          `json:"deleted_at"`
	PatternInstance PatternInstance `json:"pattern_instance"`
}
