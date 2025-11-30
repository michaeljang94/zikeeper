package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaeljang94/zikeeper/internal/service"
)

func GetUser(c *gin.Context) {
	var getUserRequest service.GetUserRequest

	if err := c.ShouldBindJSON(&getUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	getUserResponse := service.GetUser(getUserRequest)

	c.JSON(http.StatusOK, getUserResponse)
}

func CreateUser(c *gin.Context) {
	var request service.CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := service.CreateUser(request)

	c.JSON(http.StatusOK, response)
}
