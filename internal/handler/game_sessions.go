package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaeljang94/zikeeper/internal/service"
)

type GameSessionsHandler struct {
	Service *service.GameSessionsService
}

func (handler *GameSessionsHandler) AddPlayerToGameSession(c *gin.Context) {
	id := c.Param("id")

	request := service.AddPlayerToGameSessionRequest{
		SessionId: id,
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := handler.Service.AddPlayerToTable(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
