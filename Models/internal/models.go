package models

type Provider struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Name	 string `json:"name"`
}

type Patient struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Name	 string `json:"name"`
}

type Appointment struct {
	ProviderID int    `json:"provider_id"`
	PatientID  int    `json:"patient_id"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}


