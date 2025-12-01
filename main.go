package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/michaeljang94/zikeeper/internal/handler"
	"github.com/michaeljang94/zikeeper/internal/repo"
	"github.com/michaeljang94/zikeeper/internal/service"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = "password"
	cfg.Net = "tcp"
	cfg.Addr = os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT")
	cfg.DBName = "zikeeper"

	fmt.Println(cfg)

	db, err := sql.Open("mysql", cfg.FormatDSN())
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
