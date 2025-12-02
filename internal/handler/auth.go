package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaeljang94/zikeeper/internal/service"
)

type AuthHandler struct {
	Service *service.AuthService
}

func (handler *AuthHandler) AuthenticateUser(c *gin.Context) {
	authUser := service.AuthUser{}

	if err := c.ShouldBindJSON(&authUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request := service.AuthenticateUserRequest{
		User: service.AuthUser{},
	}

	response, err := handler.Service.AuthenticateUser(request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
	}

	c.JSON(http.StatusOK, response)
}
