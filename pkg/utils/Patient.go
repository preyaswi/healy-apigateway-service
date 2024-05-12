package models
type PatientSignUp struct {
	FullName        string `json:"full_name" binding:"required" validate:"required"`
	Email           string `json:"email" binding:"required" validate:"required"`
	Password        string `json:"password" binding:"required" validate:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required" validate:"required"`
	Gender          string `json:"gender" binding:"required" validate:"required"`
	ContactNumber   string `json:"contact_number" binding:"required" validate:"required"`
}


type SignupdetailResponse struct {
	Id            uint
	Fullname      string
	Email         string
	Gender        string
	Contactnumber string
}
type TokenPatient struct {
	Patient         SignupdetailResponse
	AccessToken  string
	RefreshToken string
}

type PatientLogin struct{
	Email string
	Password string
}
