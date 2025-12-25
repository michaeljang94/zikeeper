package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaeljang94/zikeeper/internal/service"
)

type TableSessionsHandler struct {
	Service *service.TableSessionsService
}

func (handler *TableSessionsHandler) AddDealerToTableSession(c *gin.Context) {
	sessionId := c.Param("session_id")

	request := service.AddDealerToTableSessionRequest{
		SessionId: sessionId,
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.Service.AddDealerToTableSession(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *TableSessionsHandler) RemoveDealerFromTableSession(c *gin.Context) {
	sessionId := c.Param("session_id")

	request := service.RemoveDealerFromTableSessionRequest{
		SessionId: sessionId,
	}

	response, err := handler.Service.RemoveDealerFromTableSession(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *TableSessionsHandler) DeleteTableSessionsByTableName(c *gin.Context) {
	tableName := c.Param("table_name")

	request := service.DeleteTableSessionsByTableNameRequest{
		TableName: tableName,
	}

	response, err := handler.Service.DeleteTableSessionsByTableName(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *TableSessionsHandler) GetTableSessions(c *gin.Context) {
	tableName := c.Param("table_name")

	request := service.GetTableSessionsRequest{
		TableName: tableName,
	}

	response, err := handler.Service.GetTableSessions(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *TableSessionsHandler) CreateTableSession(c *gin.Context) {
	tableName := c.Param("table_name")

	request := service.CreateTableSessionRequest{
		TableName: tableName,
	}

	response, err := handler.Service.CreateTableSession(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *TableSessionsHandler) DeleteTableSessionBySessionId(c *gin.Context) {
	tableName := c.Param("table_name")

	request := service.DeleteTableSessionBySessionIdRequest{
		TableName: tableName,
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.Service.DeleteTableSessionBySessionId(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
