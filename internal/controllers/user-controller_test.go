package controllers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	userservicemocks "github.com/aryannr97/iam-server/internal/services/mocks"
	"github.com/aryannr97/iam-server/models"
	"github.com/gin-gonic/gin"
	testifymock "github.com/stretchr/testify/mock"
)

func TestUserController_AddUser(t *testing.T) {
	userServiceMock := new(userservicemocks.MockUserService)
	uc := UserController{UserService: userServiceMock}
	res := &models.AddUserResponse{
		ID: "qwertyuisdfg",
	}
	reqBody1 := `{
		"firstname":  "John",
		"lastname":   "Doe",
		"email":      "john@mail.com",
		"phone":      "4567891234",
		"country":    "US"
	}`

	userServiceMock.On("AddUser", testifymock.Anything).Return(res, nil).Once()
	userServiceMock.On("AddUser", testifymock.Anything).Return(nil, errors.New("add_error")).Once()

	w := httptest.NewRecorder()
	ctx1, _ := gin.CreateTestContext(w)
	req1, _ := http.NewRequest("POST", "/users/add", bytes.NewBuffer([]byte(reqBody1)))
	ctx1.Request = req1

	ctx2, _ := gin.CreateTestContext(w)
	req2, _ := http.NewRequest("POST", "/users/add", bytes.NewBuffer([]byte("")))
	ctx2.Request = req2

	ctx3, _ := gin.CreateTestContext(w)
	req3, _ := http.NewRequest("POST", "/users/add", bytes.NewBuffer([]byte(reqBody1)))
	ctx3.Request = req3

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		uc   *UserController
		args args
	}{
		{
			"Positive case",
			&uc,
			args{
				ctx1,
			},
		},
		{
			"Negative case 1",
			&uc,
			args{
				ctx2,
			},
		},
		{
			"Negative case 2",
			&uc,
			args{
				ctx3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.uc.AddUser(tt.args.c)
		})
	}
}

func TestUserController_GetUser(t *testing.T) {
	userServiceMock := new(userservicemocks.MockUserService)
	uc := UserController{UserService: userServiceMock}
	res := &models.UserDetailsResponse{}

	w := httptest.NewRecorder()
	ctx1, _ := gin.CreateTestContext(w)
	req1, _ := http.NewRequest("GET", "/users/qwertyuiog", nil)
	ctx1.Request = req1

	userServiceMock.On("GetUser", testifymock.Anything).Return(res, nil).Once()
	userServiceMock.On("GetUser", testifymock.Anything).Return(nil, errors.New("get-error")).Once()

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		uc   *UserController
		args args
	}{
		{
			"Positive case",
			&uc,
			args{ctx1},
		},
		{
			"Negative case",
			&uc,
			args{ctx1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.uc.GetUser(tt.args.c)
		})
	}
}

func TestUserController_DeleteUser(t *testing.T) {
	userServiceMock := new(userservicemocks.MockUserService)
	uc := UserController{UserService: userServiceMock}

	w := httptest.NewRecorder()
	ctx1, _ := gin.CreateTestContext(w)
	req1, _ := http.NewRequest("GET", "/users/qwertyuiog", nil)
	ctx1.Request = req1

	userServiceMock.On("DeleteUser", testifymock.Anything).Return(nil).Once()
	userServiceMock.On("DeleteUser", testifymock.Anything).Return(errors.New("delete-error")).Once()

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		uc   *UserController
		args args
	}{
		{
			"Positive case",
			&uc,
			args{ctx1},
		},
		{
			"Negative case",
			&uc,
			args{ctx1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.uc.DeleteUser(tt.args.c)
		})
	}
}

func TestUserController_UpdateUser(t *testing.T) {
	userServiceMock := new(userservicemocks.MockUserService)
	uc := UserController{UserService: userServiceMock}
	reqBody1 := `{
		"firstname":  "John",
		"lastname":   "Doe",
		"email":      "john@mail.com",
		"phone":      "4567891234",
		"country":    "US"
	}`

	userServiceMock.On("UpdateUser", testifymock.Anything, testifymock.Anything).Return(nil).Once()
	userServiceMock.On("UpdateUser", testifymock.Anything, testifymock.Anything).Return(errors.New("update_error")).Once()

	w := httptest.NewRecorder()
	ctx1, _ := gin.CreateTestContext(w)
	req1, _ := http.NewRequest("PUT", "/users/qwertyuiog", bytes.NewBuffer([]byte(reqBody1)))
	ctx1.Request = req1

	ctx2, _ := gin.CreateTestContext(w)
	req2, _ := http.NewRequest("PUT", "/users/qwertyuiog", bytes.NewBuffer([]byte("")))
	ctx2.Request = req2

	ctx3, _ := gin.CreateTestContext(w)
	req3, _ := http.NewRequest("PUT", "/users/qwertyuiog", bytes.NewBuffer([]byte(reqBody1)))
	ctx3.Request = req3

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		uc   *UserController
		args args
	}{
		{
			"Positive case",
			&uc,
			args{
				ctx1,
			},
		},
		{
			"Negative case 1",
			&uc,
			args{
				ctx2,
			},
		},
		{
			"Negative case 2",
			&uc,
			args{
				ctx3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.uc.UpdateUser(tt.args.c)
		})
	}
}
