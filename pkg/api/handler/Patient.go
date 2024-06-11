package handler

import (
	"context"
	"encoding/json"
	"healy-apigateway/pkg/api/response"
	interfaces "healy-apigateway/pkg/client/interface"
	"healy-apigateway/pkg/config"
	models "healy-apigateway/pkg/utils"
	"io"

	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type PatientHandler struct {
	Grpc_client interfaces.PatientClient
	oauthConfig *oauth2.Config
}

func NewPatientHandler(PatientClient interfaces.PatientClient,cfg config.Config) *PatientHandler {
	return &PatientHandler{
		Grpc_client: PatientClient,
		oauthConfig: &oauth2.Config{
			ClientID: cfg.GoogleClientId,
			ClientSecret: cfg.GoogleSecretId,
			RedirectURL: cfg.RedirectURL,
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		},
	}
}

func (p *PatientHandler) GoogleLogin(c *fiber.Ctx) error {
    url := p.oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
    return c.Redirect(url)
}
func (p *PatientHandler) GoogleCallback(c *fiber.Ctx) error {
    code := c.Query("code")
    token, err := p.oauthConfig.Exchange(context.Background(), code)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to exchange token",
        })
    }

    client := p.oauthConfig.Client(context.Background(), token)
    resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to get user info",
        })
    }
    defer resp.Body.Close()

	userInfo, err := io.ReadAll(resp.Body)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to read user info",
        })
    }

    var googleUser models.GoogleUserInfo
    if err := json.Unmarshal(userInfo, &googleUser); err != nil {
        errs := response.ClientResponse("failed to parse user info",nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	
    }

    patient, err := p.Grpc_client.GoogleSignIn(googleUser.ID, googleUser.Email, googleUser.Name)
    if err != nil {
		errs := response.ClientResponse( "error: Failed to authenticate with patient service", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)

    }
	success := response.ClientResponse("Patient created successfully", patient, nil)
	return c.Status(201).JSON(success)
}


func (p *PatientHandler) PatientDetails(c *fiber.Ctx) error {

	patientID := c.Locals("user_id").(int)

	patientDetails, err := p.Grpc_client.PatientDetails(patientID)
	if err != nil {
		errorRes := response.ClientResponse("failed to retrieve patient details", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)

	}

	successRes := response.ClientResponse("patient Details", patientDetails, nil)
	return c.Status(200).JSON(successRes)

}
func (p *PatientHandler) UpdatePatientDetails(c *fiber.Ctx) error {

	user_id := c.Locals("user_id").(int)

	var patient models.PatientDetails

	if err := c.BodyParser(&patient); err != nil {
		errs := response.ClientResponse("details are not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	if err := validator.New().Struct(patient); err != nil {
		errs := response.ClientResponse("Constraints not satisfied", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	updatedDetails, err := p.Grpc_client.UpdatePatientDetails(patient, user_id)
	if err != nil {
		errorRes := response.ClientResponse("failed update user", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)
	}

	successRes := response.ClientResponse("Updated User Details", updatedDetails, nil)
	return c.Status(200).JSON(successRes)

}
func (p *PatientHandler) ListPatients(c *fiber.Ctx) error {
	listedPatients, err := p.Grpc_client.ListPatients()
	if err != nil {
		errorRes := response.ClientResponse("failed retreiving list of patients", nil, err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errorRes)
	}
	successRes := response.ClientResponse("retrived list of patients", listedPatients, nil)
	return c.Status(200).JSON(successRes)
}