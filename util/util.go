package util

// ErrorResponse represents a standardized error response structure
type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
