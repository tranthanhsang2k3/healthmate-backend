package utils


const(
	ErrorNotFound = "not_found"
	ErrorInternal = "internal_error"
	ErrorBadRequest = "bad_request"
	ErrorUnauthorized = "unauthorized"
	ErrorForbidden    = "forbidden"
)

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func ErrorResponseFull(status bool, message string) *ErrorResponse {
	return &ErrorResponse{
		Status:  status,
		Message: message,
	}
}