package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaeljang94/zikeeper/internal/service"
)

type UserHandler struct {
	Service *service.UserService
}

func (handler *UserHandler) GetUserByUsername(c *gin.Context) {
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
