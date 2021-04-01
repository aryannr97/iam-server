package controllers

import (
	"net/http"

	"github.com/aryannr97/iam-server/internal/services"
	"github.com/aryannr97/iam-server/models"
	resterrors "github.com/aryannr97/iam-server/pkg/rest-errors"
	"github.com/gin-gonic/gin"
)

// IUserController defines all user specific controller layer actions
type IUserController interface {
	AddUser(*gin.Context)
	GetUser(*gin.Context)
	DeleteUser(*gin.Context)
	UpdateUser(*gin.Context)
}

// UserController handles all user related requests
type UserController struct {
	UserService services.IUserService
}

// AddUser process requests for adding new user
// @Summary Add user API
// @Tags user-api
// @Param request body models.AddUserRequest true "Request Body"
// @Success 201 {object} models.AddUserResponse
// @Failure 500 {object} models.RestError
// @Router /add [post]
func (uc *UserController) AddUser(c *gin.Context) {
	var requestPayload models.AddUserRequest
	if err := c.ShouldBindJSON(&requestPayload); err != nil {
		c.JSON(http.StatusBadRequest, resterrors.NewBadRequest(err.Error()))

		return
	}

	res, err := uc.UserService.AddUser(requestPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resterrors.NewInternalServerError(err.Error()))

		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetUser process requests for fetching user details
// @Summary Get user API
// @Tags user-api
// @Param userid path string true "userid"
// @Success 200 {object} models.UserDetailsResponse
// @Failure 500 {object} models.RestError
// @Router /{userid} [get]
func (uc *UserController) GetUser(c *gin.Context) {
	userid := c.Param("id")
	res, err := uc.UserService.GetUser(userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resterrors.NewInternalServerError(err.Error()))

		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteUser process requests for deleting user
// @Summary Delete user API
// @Tags user-api
// @Param userid path string true "userid"
// @Success 204
// @Failure 500 {object} models.RestError
// @Router /{userid} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	userid := c.Param("id")
	err := uc.UserService.DeleteUser(userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resterrors.NewInternalServerError(err.Error()))

		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// UpdateUser process requests for updating user details
// @Summary Update user API
// @Tags user-api
// @Param userid path string true "userid"
// @Param request body models.UpdateUserRequest true "Request Body"
// @Success 204
// @Failure 500 {object} models.RestError
// @Router /{userid} [put]
func (uc *UserController) UpdateUser(c *gin.Context) {
	userid := c.Param("id")
	var requestPayload models.UpdateUserRequest

	if err := c.ShouldBindJSON(&requestPayload); err != nil {
		c.JSON(http.StatusBadRequest, resterrors.NewBadRequest(err.Error()))

		return
	}

	err := uc.UserService.UpdateUser(userid, requestPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, resterrors.NewInternalServerError(err.Error()))

		return
	}

	c.JSON(http.StatusNoContent, nil)
}
