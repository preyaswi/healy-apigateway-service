package models

type Prescription struct {
	Medicine   string    `json:"medicine" validate:"required"`
	Dosage     string    `json:"dosage" validate:"required"`
	Notes      string    `json:"notes"`
}

type PrescriptionRequest struct {
    DoctorID   int    `json:"doctor_id"`
    PatientID  int `json:"patient_id"`
    Medicine   string `json:"medicine" validate:"required"`
    Dosage     string `json:"dosage" validate:"required"`
    Notes      string `json:"notes"`
}