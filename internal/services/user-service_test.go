package services

import (
	"errors"
	"reflect"
	"testing"

	"github.com/aryannr97/data-server/pkg/grpc/user"
	clientmocks "github.com/aryannr97/data-server/pkg/grpc/user/mocks"
	"github.com/aryannr97/iam-server/models"
	testifymock "github.com/stretchr/testify/mock"
)

func TestUserService_AddUser(t *testing.T) {
	userClientMock := new(clientmocks.MockUserServiceClient)
	uc := &UserService{UserServiceClient: userClientMock}
	res := &user.UserAddResponse{Id: "qwertyuiog"}

	userClientMock.On("AddUser", testifymock.Anything, testifymock.Anything).Return(res, nil).Once()
	userClientMock.On("AddUser", testifymock.Anything, testifymock.Anything).Return(&user.UserAddResponse{}, errors.New("add-error")).Once()

	type args struct {
		userInfo models.AddUserRequest
	}
	tests := []struct {
		name    string
		uc      *UserService
		args    args
		want    *models.AddUserResponse
		wantErr bool
	}{
		{
			"Positive case",
			uc,
			args{
				models.AddUserRequest{},
			},
			&models.AddUserResponse{ID: "qwertyuiog"},
			false,
		},
		{
			"Negative case",
			uc,
			args{
				models.AddUserRequest{},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.AddUser(tt.args.userInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.AddUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_GetUser(t *testing.T) {
	userClientMock := new(clientmocks.MockUserServiceClient)
	uc := &UserService{UserServiceClient: userClientMock}
	res := &user.UserGetResponse{
		Id:        "qwertyuiog",
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john@mail.com",
		Phone:     "4567891234",
		Country:   "US",
	}

	userClientMock.On("GetUser", testifymock.Anything, testifymock.Anything).Return(res, nil).Once()
	userClientMock.On("GetUser", testifymock.Anything, testifymock.Anything).Return(&user.UserGetResponse{}, errors.New("get-error")).Once()

	type args struct {
		userid string
	}
	tests := []struct {
		name    string
		uc      *UserService
		args    args
		want    *models.UserDetailsResponse
		wantErr bool
	}{
		{
			"Positive case",
			uc,
			args{
				"qwertyuiog",
			},
			&models.UserDetailsResponse{
				Firstname:  "John",
				Laststname: "Doe",
				Email:      "john@mail.com",
				Phone:      "4567891234",
				Country:    "US",
			},
			false,
		},
		{
			"Negative case",
			uc,
			args{
				"qwertyuiog",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetUser(tt.args.userid)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	userClientMock := new(clientmocks.MockUserServiceClient)
	uc := &UserService{UserServiceClient: userClientMock}

	userClientMock.On("DeleteUser", testifymock.Anything, testifymock.Anything).Return(&user.Empty{}, nil).Once()

	type args struct {
		userid string
	}
	tests := []struct {
		name    string
		uc      *UserService
		args    args
		wantErr bool
	}{
		{
			"Case 1",
			uc,
			args{
				userid: "qwertyuiog",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.uc.DeleteUser(tt.args.userid); (err != nil) != tt.wantErr {
				t.Errorf("UserService.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	userClientMock := new(clientmocks.MockUserServiceClient)
	uc := &UserService{UserServiceClient: userClientMock}

	userClientMock.On("UpdateUser", testifymock.Anything, testifymock.Anything).Return(&user.Empty{}, nil).Once()

	type args struct {
		userid  string
		request models.UpdateUserRequest
	}
	tests := []struct {
		name    string
		uc      *UserService
		args    args
		wantErr bool
	}{
		{
			"Positive case",
			uc,
			args{
				"qwertyuiog",
				models.UpdateUserRequest{
					Firstname: "John",
					Lastname:  "Doe",
					Email:     "john@mail.com",
					Phone:     "4567891245",
					Country:   "US",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.uc.UpdateUser(tt.args.userid, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("UserService.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
