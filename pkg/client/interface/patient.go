package interfaces

import models "healy-apigateway/pkg/utils"

type PatientClient interface {
	PatientsSignUp(patient models.PatientSignUp) (models.TokenPatient, error)
	PatientLogin(patient models.PatientLogin)(models.TokenPatient,error)
}
