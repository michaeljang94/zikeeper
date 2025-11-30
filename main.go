package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/michaeljang94/zikeeper/internal/handler"
)

func main() {
	fmt.Println("Hello World")

	router := gin.Default()

	router.GET("/get_user", handler.GetUser)
	router.POST("/create_user", handler.CreateUser)

	router.Run()
}
