package models

// RestError forms common error object
type RestError struct {
	StatusCode int64  `json:"statusCode"`
	Status     string `json:"status"`
	ErrMessage string `json:"errorMessage"`
}
