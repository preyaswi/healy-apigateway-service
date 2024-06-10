package interfaces

import (
	models "healy-apigateway/pkg/utils"
)

type PatientClient interface {
	GoogleSignIn(googleID, email, name string) (models.TokenPatient, error)
	PatientDetails(user_id int) (models.SignupdetailResponse, error)
	UpdatePatientDetails(patient models.PatientDetails, patient_id int) (models.PatientDetails,error)
	ListPatients()([]models.SignupdetailResponse,error)
	

}
