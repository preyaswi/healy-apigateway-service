package models

type DoctorSignUp struct {
	FullName          string `json:"full_name" binding:"required" validate:"required"`
	Email             string `json:"email"  binding:"required" validate:"required"`
	PhoneNumber       string `json:"phone_number"  binding:"required" validate:"required"`
	Password          string `json:"password"  binding:"required" validate:"required"`
	ConfirmPassword   string `json:"confirm_password"  binding:"required" validate:"required"`
	Specialization    string `json:"specialization"  binding:"required" validate:"required"`
	YearsOfExperience int32  `json:"years_of_experience"  binding:"required" validate:"required"`
	LicenseNumber     string `json:"license_number"  binding:"required" validate:"required"`
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