package models
type PatientSignUp struct {
	Fullname        string
	Email           string
	Password        string
	Confirmpassword string
	Gender          string
	Contactnumber   string
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