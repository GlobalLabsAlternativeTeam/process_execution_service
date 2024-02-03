package domain

type StorageServer struct {
	BaseURL string
}

type TreatmentLight struct {
	Treatment_ID       string
	Treatment_Name     string
	Treatment_Status   string
	Treatment_Progress float64
}
