package models

// AddUserRequest request payload for adding user
type AddUserRequest struct {
	Firstname  string `json:"firstname"`
	Laststname string `json:"lastname"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Country    string `json:"country"`
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

// UpdateUserRequest request payload for updating user details
type UpdateUserRequest struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Country   string `json:"country" binding:"required"`
}
