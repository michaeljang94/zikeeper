package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/michaeljang94/zikeeper/internal/handler"
	"github.com/michaeljang94/zikeeper/internal/repo"
	"github.com/michaeljang94/zikeeper/internal/service"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

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

	authService := service.AuthService{
		UserRepo: &userRepo,
	}
	authHandler := handler.AuthHandler{
		Service: &authService,
	}

	router := gin.New()
	router.Use(CORSMiddleware())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})
	router.GET("/get_user/:id", userHandler.GetUser)
	router.POST("/create_user", userHandler.CreateUser)

	router.POST("/auth", authHandler.AuthenticateUser)

	router.Run()

	// autotls.Run(router, "zikeeper.com")
}
