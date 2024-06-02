package models

type Prescription struct {
	DoctorName string    `json:"doctor_name" validate:"required"`
	Medicine   string    `json:"medicine" validate:"required"`
	Dosage     string    `json:"dosage" validate:"required"`
	Notes      string    `json:"notes"`
}

type PrescriptionRequest struct {
    DoctorID   int    `json:"doctor_id"`
    PatientID  int `json:"patient_id"`
    DoctorName string `json:"doctor_name" validate:"required"`
    Medicine   string `json:"medicine" validate:"required"`
    Dosage     string `json:"dosage" validate:"required"`
    Notes      string `json:"notes"`
}