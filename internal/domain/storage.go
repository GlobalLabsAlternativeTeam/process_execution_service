package domain

type StorageServer struct {
	BaseURL string
}

type TreatmentLight struct {
	Treatment_ID       string `json:"treatment_id"`
	Treatment_Name     string `json:"treatment_name"`
	Treatment_Status   string `json:"treatment_status"`
	Treatment_Progress int    `json:"treatment_progress"`
}
