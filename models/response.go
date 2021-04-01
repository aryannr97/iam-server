package models

// RestError forms common error object
type RestError struct {
	StatusCode int64  `json:"statusCode"`
	Status     string `json:"status"`
	ErrMessage string `json:"errorMessage"`
}

// AddUserResponse response body after adding user
type AddUserResponse struct {
	ID string `json:"id"`
}

// UserDetailsResponse response payload for user info
type UserDetailsResponse struct {
	Firstname  string `json:"firstname"`
	Laststname string `json:"lastname"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Country    string `json:"country"`
}
