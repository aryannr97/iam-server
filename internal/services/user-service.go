package services

import (
	"context"

	"github.com/aryannr97/data-server/pkg/grpc/user"
	"github.com/aryannr97/iam-server/models"
)

// IUserService defines all user specific service layer actions
type IUserService interface {
	AddUser(models.AddUserRequest) (*models.AddUserResponse, error)
	GetUser(string) (*models.UserDetailsResponse, error)
	DeleteUser(string) error
	UpdateUser(string, models.UpdateUserRequest) error
}

// UserService handles all user related service requests
type UserService struct {
	UserServiceClient user.UserServiceClient
}

// AddUser process requests for adding new user
func (uc *UserService) AddUser(userInfo models.AddUserRequest) (*models.AddUserResponse, error) {
	res, err := uc.UserServiceClient.AddUser(context.Background(), &user.UserAddRequest{
		Firstname: userInfo.Firstname,
		Lastname:  userInfo.Laststname,
		Email:     userInfo.Email,
		Phone:     userInfo.Phone,
		Country:   userInfo.Country,
	})

	if err != nil {
		return nil, err
	}

	return &models.AddUserResponse{ID: res.Id}, nil
}

// GetUser process requests for retrieving user details
func (uc *UserService) GetUser(userid string) (*models.UserDetailsResponse, error) {
	res, err := uc.UserServiceClient.GetUser(context.Background(), &user.UserGetReuqest{
		Id: userid,
	})

	if err != nil {
		return nil, err
	}

	return &models.UserDetailsResponse{
		Firstname:  res.Firstname,
		Laststname: res.Lastname,
		Email:      res.Email,
		Phone:      res.Phone,
		Country:    res.Country,
	}, nil
}

// DeleteUser process requests for deleting user
func (uc *UserService) DeleteUser(userid string) error {
	_, err := uc.UserServiceClient.DeleteUser(context.Background(), &user.UserDeleteRequest{Id: userid})

	return err
}

// UpdateUser process requests for adding new user
func (uc *UserService) UpdateUser(userid string, request models.UpdateUserRequest) error {
	update := &user.UserUpdateRequest{
		Id:        userid,
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Email:     request.Email,
		Phone:     request.Phone,
		Country:   request.Country,
	}
	_, err := uc.UserServiceClient.UpdateUser(context.Background(), update)
	return err
}
