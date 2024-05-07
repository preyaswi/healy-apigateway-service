package response

type Response struct {
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}
func ClientResponse(message string, data interface{}, err interface{}) Response {

	return Response{
		Message:    message,
		Data:       data,
		Error:      err,
	}

}
