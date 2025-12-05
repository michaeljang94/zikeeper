package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaeljang94/zikeeper/internal/service"
)

type AuthHandler struct {
	Service *service.AuthService
}

func (handler *AuthHandler) CreateNewUser(c *gin.Context) {
	var request service.CreateAuthUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.Service.CreateNewUser(request)

	if err != nil {
		if err.Error() == "duplicate entry" {
			c.JSON(http.StatusConflict, "duplicate entry")
		}

		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *AuthHandler) AuthenticateUser(c *gin.Context) {
	authUser := service.AuthUser{}

	if err := c.ShouldBindJSON(&authUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request := service.AuthenticateUserRequest{
		AuthUser: authUser,
	}

	response, err := handler.Service.AuthenticateUser(request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
