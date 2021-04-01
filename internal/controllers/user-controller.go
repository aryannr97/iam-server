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

	c.JSON(http.StatusOK, res)
}

// GetUser process requests for fetching user details
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
