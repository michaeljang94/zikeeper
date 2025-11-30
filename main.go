package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/michaeljang94/zikeeper/internal/handler"
	"github.com/michaeljang94/zikeeper/internal/repo"
	"github.com/michaeljang94/zikeeper/internal/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@/zikeeper")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	userRepo := repo.UserRepo{
		Db: db,
	}
	userService := service.UserService{
		Repo: &userRepo,
	}
	userHandler := handler.UserHandler{
		Service: &userService,
	}

	router := gin.Default()

	router.GET("/get_user", userHandler.GetUser)
	router.POST("/create_user", userHandler.CreateUser)

	router.Run()
}
