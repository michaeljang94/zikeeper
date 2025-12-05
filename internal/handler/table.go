package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaeljang94/zikeeper/internal/service"
)

type TableHandler struct {
	Service *service.TableService
}

func (handler *TableHandler) CreateTable(c *gin.Context) {
	var request service.CreateTableRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := handler.Service.CreateTable(request)

	if err != nil {
		if err.Error() == "table already exists" {
			c.JSON(http.StatusConflict, "table already exists")
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (handler *TableHandler) GetTables(c *gin.Context) {
	request := service.GetTablesRequest{}

	response, err := handler.Service.GetTables(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *TableHandler) GetTableByName(c *gin.Context) {
	id := c.Param("id")

	request := service.GetTableByNameRequest{
		TableName: id,
	}

	response, err := handler.Service.GetTableByName(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *TableHandler) AddPlayerToTable(c *gin.Context) {
	request := service.AddPlayerToTableRequest{}

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
