package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaeljang94/zikeeper/internal/service"
)

type UserHandler struct {
	Service *service.UserService
}

func (handler *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	getUserRequest := service.GetUserRequest{
		UserName: id,
	}

	getUserResponse, err := handler.Service.GetUserByUserName(getUserRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, getUserResponse)
}

func (handler *UserHandler) CreateUser(c *gin.Context) {
	var request service.CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.Service.CreateUser(request)

	if err != nil {
		if err.Error() == "duplicate entry" {
			c.JSON(http.StatusConflict, "duplicate entry")
		}

		return
	}

	c.JSON(http.StatusOK, response)
}
