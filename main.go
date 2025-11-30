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

	router.Run()
}
