package models

type DoctorSignUp struct {
	FullName          string `json:"full_name"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phone_number"`
	Password          string `json:"password"`
	ConfirmPassword   string `json:"confirm_password"`
	Specialization    string `json:"specialization"`
	YearsOfExperience int32  `json:"years_of_experience"`
	LicenseNumber     string `json:"license_number"`
}
type DoctorDetail struct {
	Id                uint
	FullName          string
	Email             string
	PhoneNumber       string
	Specialization    string
	YearsOfExperience int32
	LicenseNumber     string
}
type DoctorSignUpResponse struct {
	DoctorDetail DoctorDetail
	AccessToken  string
	RefreshToken string
}
type DoctorLogin struct {
	Email    string
	Password string
}
type DoctorsDetails struct{
	DoctorDetail DoctorDetail
	Rating int32
}