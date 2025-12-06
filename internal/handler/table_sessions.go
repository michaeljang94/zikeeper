package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaeljang94/zikeeper/internal/service"
)

type TableSessionsHandler struct {
	Service *service.TableSessionsService
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
