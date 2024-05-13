package models

type DoctorSignUp struct {
    FullName          string `json:"full_name" binding:"required" validate:"required,min=3,max=50"`
    Email             string `json:"email"  binding:"required" validate:"required,email"`
    PhoneNumber       string `json:"phone_number"  binding:"required" validate:"required,min=10,max=15"`
    Password          string `json:"password"  binding:"required" validate:"required,min=6,max=20"`
    ConfirmPassword   string `json:"confirm_password"  binding:"required" validate:"required,min=6,max=20"`
    Specialization    string `json:"specialization"  binding:"required" validate:"required"`
    YearsOfExperience int32  `json:"years_of_experience"  binding:"required" validate:"required,min=0,max=50"`
    LicenseNumber     string `json:"license_number"  binding:"required" validate:"required,min=6,max=20"`
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
type IndDoctorDetail struct {
	Id                uint
	FullName          string
	Email             string
	PhoneNumber       string
	Specialization    string
	YearsOfExperience int32
	LicenseNumber     string
	Rating int32
}