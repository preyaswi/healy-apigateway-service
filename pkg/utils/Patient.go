package models

type PatientSignUp struct {
	FullName        string `json:"full_name" binding:"required" validate:"required,min=3,max=50"`
	Email           string `json:"email" binding:"required" validate:"required"`
	Password        string `json:"password" binding:"required" validate:"required,min=3"`
	ConfirmPassword string `json:"confirm_password" binding:"required" validate:"required"`
	Gender          string `json:"gender" binding:"required" validate:"required"`
	ContactNumber   string `json:"contact_number" binding:"required" validate:"required,min=10,max=15"`
}

type SignupdetailResponse struct {
	Id            uint
	Fullname      string
	Email         string
	Gender        string
	Contactnumber string
}
type TokenPatient struct {
	Patient      SignupdetailResponse
	AccessToken  string
	RefreshToken string
}

type PatientLogin struct {
	Email    string
	Password string
}
type PatientDetails struct {
	Fullname      string
	Email         string
	Gender        string
	Contactnumber string
}

type UpdatePassword struct{
	OldPassword        string `json:"old_password" binding:"required"`
    NewPassword        string `json:"new_password" binding:"required"`
    ConfirmNewPassword string `json:"confirm_new_password" binding:"required"`
}