package domain

type LightTreatment struct {
	TreatmentID       string `json:"treatment_id"`
	TreatmentName     string `json:"treatment_name"`
	TreatmentStatus   string `json:"treatment_status"`
	TreatmentProgress int    `json:"treatment_progress"`
}
