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
	sessionId := c.Param("session_id")

	request := service.GetPlayersForSessionIdRequest{
		SessionId: sessionId,
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
	sessionId := c.Param("session_id")
	tableName := c.Param("table_name")

	request := service.AddPlayerToPlayerSessionRequest{
		SessionId: sessionId,
		TableName: tableName,
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

func (handler *PlayerSessionsHandler) DeletePlayerFromPlayerSession(c *gin.Context) {
	sessionId := c.Param("session_id")
	tableName := c.Param("table_name")

	request := service.DeletePlayerFromPlayerSessionRequest{
		SessionId: sessionId,
		TableName: tableName,
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := handler.Service.DeletePlayerFromPlayerSession(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
