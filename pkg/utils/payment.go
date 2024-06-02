package models

type CombinedBookingDetails struct {
	BookingId     uint
	PatientId     uint
	DoctorId      uint
	DoctorName    string
	DoctorEmail   string
	Fees          uint64
	PaymentStatus string
}
