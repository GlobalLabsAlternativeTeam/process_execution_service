package domain

type Task struct {
	ID          int    `json:"id"`
	Level       int    `json:"level"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	BlockedBy   []any  `json:"blocked_by"`
	Responsible int    `json:"responsible"`
	TimeLimit   int    `json:"time_limit"`
	Children    []any  `json:"children"`
	Comment     struct {
		Value string `json:"value"`
	} `json:"comment"`
}

type SchemaInstance struct {
	SchemaInstanceID      string `json:"schema_instance_id"`
	SchemaID              string `json:"schema_id"`
	AuthorID              string `json:"author_id"`
	SchemaName            string `json:"schema_name"`
	PatternInstanceStatus string `json:"PatternInstanceStatus"`
	CreatedAt             struct {
		Seconds int `json:"seconds"`
		Nanos   int `json:"nanos"`
	} `json:"created_at"`
	UpdatedAt struct {
		Seconds int `json:"seconds"`
		Nanos   int `json:"nanos"`
	} `json:"updated_at"`
	DeletedAt struct {
		Seconds int `json:"seconds"`
		Nanos   int `json:"nanos"`
	} `json:"deleted_at"`
	Tasks []Task `json:"tasks"`
}
type Treatment struct {
	TreatmentdID    string `json:"treatmentd_id"`
	DoctorID        int    `json:"doctor_id"`
	PatientID       int    `json:"patient_id"`
	Status          string `json:"status"`
	StartedAt       string `json:"started_at"`
	FinishedAt      string `json:"finished_at"`
	DeletedAt       string `json:"deleted_at"`
	PatternInstance struct {
	} `json:"pattern_instance"`
}
