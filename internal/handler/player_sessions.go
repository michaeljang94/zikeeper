package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaeljang94/zikeeper/internal/service"
)

type PlayerSessionsHandler struct {
	Service *service.PlayerSessionsService
}

func (handler *PlayerSessionsHandler) GetPlayersForSessionId(c *gin.Context) {
	id := c.Param("id")

	request := service.GetPlayersForSessionIdRequest{
		SessionId: id,
	}

	response, err := handler.Service.GetPlayersForSessionId(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *PlayerSessionsHandler) AddPlayerToPlayerSession(c *gin.Context) {
	id := c.Param("id")

	request := service.AddPlayerToPlayerSessionRequest{
		SessionId: id,
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := handler.Service.AddPlayerToPlayerSession(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
