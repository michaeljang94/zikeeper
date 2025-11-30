package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/michaeljang94/zikeeper/internal"
)

func main() {
	fmt.Println("Hello World")

	router := gin.Default()

	router.GET("/getUser", internal.GetUser)

	router.Run()
}
