package models

// Struct response API
type APIResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(message string, data interface{}) APIResponse {
	return APIResponse{
		Status:  true,
		Message: message,
		Data:    data,
	}
}
func ErrorResponse(message string) APIResponse {
	return APIResponse{
		Status:  false,
		Message: message,
		Data:    nil,
	}
}
