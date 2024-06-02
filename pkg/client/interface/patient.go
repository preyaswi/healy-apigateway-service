package interfaces

import (
	"context"
	models "healy-apigateway/pkg/utils"
)

type PatientClient interface {
	PatientsSignUp(patient models.PatientSignUp) (models.TokenPatient, error)
	PatientLogin(patient models.PatientLogin) (models.TokenPatient, error)
	PatientDetails(user_id int) (models.SignupdetailResponse, error)
	UpdatePatientDetails(patient models.PatientDetails, patient_id int) (models.PatientDetails,error)
	UpdatePassword(ctx context.Context,userid int,body models.UpdatePassword)error
	ListPatients()([]models.SignupdetailResponse,error)
	CreatePrescription(prescription models.PrescriptionRequest)(models.PrescriptionRequest,error)

}
