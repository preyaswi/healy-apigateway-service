package models

type GoogleUserInfo struct {
	ID    string `json:"id"` // Google's unique user ID
	Email string `json:"email"`
	Name  string `json:"name"`
	AccessToken  string   `json:"accesstoken"`
    RefreshToken string   `json:"refreshtoken"`
    TokenExpiry string `json:"tokenexpiry"`
}
type GoogleSignupdetailResponse struct {
	Id       string
	GoogleId string
	FullName string
	Email    string
}

type SignupdetailResponse struct {
	Id            string
	Fullname      string
	Email         string
	Gender        string
	Contactnumber string
}
type TokenPatient struct {
	Patient      GoogleSignupdetailResponse
	AccessToken  string
	RefreshToken string
}

type PatientDetails struct {
	Fullname      string
	Email         string
	Gender        string
	Contactnumber string
}

type Patient struct {
	BookingId     uint
	PaymentStatus string
	PatientId     uint
	Fullname      string
	Email         string
	Gender        string
	Contactnumber string
}
