package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaeljang94/zikeeper/internal/service"
)

type UserHandler struct {
	Service *service.UserService
}

func (handler *UserHandler) GetUsers(c *gin.Context) {
	request := service.GetUsersRequest{}

	response, err := handler.Service.GetUsers(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *UserHandler) GetUserByUsername(c *gin.Context) {
	id := c.Param("id")

	getUserRequest := service.GetUserRequest{
		UserName: id,
	}

	getUserResponse, err := handler.Service.GetUserByUserName(getUserRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, getUserResponse)
}

func (handler *UserHandler) GetUsersScoreboard(c *gin.Context) {
	req := service.GetUsersScoreboardRequest{
		Limit: 10,
	}

	res, err := handler.Service.GetUsersScoreboard(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (handler *UserHandler) UpdateUserByUsername(c *gin.Context) {
	id := c.Param("id")

	req := service.UpdateUserByUsernameRequest{
		Username: id,
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := handler.Service.UpdateUserByUsername(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
