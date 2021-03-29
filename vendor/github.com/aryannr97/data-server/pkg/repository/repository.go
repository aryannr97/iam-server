package repository

import "context"

// IUserRepository defines user operations
type IUserRepository interface {
	AddUser(context.Context, UserEntity) (string, error)
	GetUser(context.Context, string) (UserDocument, error)
	UpdateUser(context.Context, string, UserEntity) error
	DeleteUser(context.Context, string) error
}
