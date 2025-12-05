package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaeljang94/zikeeper/internal/service"
)

type TableHandler struct {
	Service *service.TableService
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
