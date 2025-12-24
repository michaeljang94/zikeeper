package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaeljang94/zikeeper/internal/service"
)

type UserHandler struct {
	Service               *service.UserService
	PlayerSessionsService *service.PlayerSessionsService
}

func (handler *UserHandler) GetSessionInfoByUsername(c *gin.Context) {
	username := c.Param("id")

	tokenUsername := c.GetString("username")
	role := c.GetString("role")

	if username != tokenUsername && role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"error":  "Unauthorized for user",
		})
		return
	}

	req := service.GetPlayerSessionByUsernameRequest{
		Username: username,
	}

	res, err := handler.PlayerSessionsService.GetPlayerSessionByUsername(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
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

func (handler *UserHandler) GetPlayerRankingByUsername(c *gin.Context) {
	id := c.Param("id")

	req := service.GetPlayerRankingByUsernameRequest{
		Username: id,
	}

	res, err := handler.Service.GetPlayerRankingByUsername(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
