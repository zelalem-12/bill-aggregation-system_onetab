package response

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewErrorResponse(message string, code int) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Code:    code,
	}
}
