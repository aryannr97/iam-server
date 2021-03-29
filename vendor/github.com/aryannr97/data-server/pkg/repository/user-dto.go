package repository

// UserEntity is user object for document
type UserEntity struct {
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
	Phone     string `json:"phone" bson:"phone"`
	Country   string `json:"country" bson:"country"`
}

// UserDocument is user object for mongo document retrieval
type UserDocument struct {
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
	Phone     string `json:"phone" bson:"phone"`
	Country   string `json:"country" bson:"country"`
}
