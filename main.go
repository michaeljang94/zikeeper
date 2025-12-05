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

	"github.com/joho/godotenv"

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

func loadEnvironmentVars() {
	env := os.Getenv("ZIKEEPER_ENV")
	if "" == env {
		env = "development"
	}

	godotenv.Load(".env." + env + ".local")
	if "test" != env {
		godotenv.Load(".env.local")
	}

	godotenv.Load(".env." + env)
	godotenv.Load() // The Original .env
}

func main() {
	loadEnvironmentVars()

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("MYSQL_USER")
	cfg.Passwd = os.Getenv("MYSQL_PASSWORD")
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
	authRepo := repo.AuthRepo{
		Db: db,
	}
	tableRepo := repo.TableRepo{
		Db: db,
	}

	userService := service.UserService{
		Repo: &userRepo,
	}
	authService := service.AuthService{
		UserRepo: &userRepo,
		AuthRepo: &authRepo,
	}
	tableService := service.TableService{
		TableRepo: &tableRepo,
	}

	userHandler := handler.UserHandler{
		Service: &userService,
	}
	authHandler := handler.AuthHandler{
		Service: &authService,
	}
	tableHandler := handler.TableHandler{
		Service: &tableService,
	}

	router := gin.New()
	router.Use(CORSMiddleware())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})
	router.GET("/api/user/:id", userHandler.GetUserByUsername)

	router.POST("/api/auth/login", authHandler.AuthenticateUser)
	router.POST("/api/auth/signup", authHandler.CreateNewUser)

	router.GET("/api/table/:id", tableHandler.GetTableByName)

	router.Run()
	// log.Fatal(autotls.Run(router, "api.zikeeper.com", "zikeeper.com"))
}
