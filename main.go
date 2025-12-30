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
	playerSessionsRepo := repo.PlayerSessionsRepo{
		Db: db,
	}
	tableSessionsRepo := repo.TableSessionsRepo{
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
	playerSessionsService := service.PlayerSessionsService{
		Repo:              &playerSessionsRepo,
		UserRepo:          &userRepo,
		TableSessionsRepo: &tableSessionsRepo,
	}
	tableSessionsService := service.TableSessionsService{
		Repo:                  &tableSessionsRepo,
		UserRepo:              &userRepo,
		PlayerSessionsRepo:    &playerSessionsRepo,
		PlayerSessionsService: &playerSessionsService,
	}

	userHandler := handler.UserHandler{
		Service:               &userService,
		PlayerSessionsService: &playerSessionsService,
	}
	authHandler := handler.AuthHandler{
		Service: &authService,
	}
	tableHandler := handler.TableHandler{
		Service: &tableService,
	}
	playerSessionsHandler := handler.PlayerSessionsHandler{
		Service: &playerSessionsService,
	}
	tableSessionsHandler := handler.TableSessionsHandler{
		Service: &tableSessionsService,
	}

	router := gin.New()
	router.Use(CORSMiddleware())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	router.POST("/api/auth/login", authHandler.AuthenticateUser)
	router.POST("/api/auth/signup", authHandler.CreateNewUser)

	protected := router.Group("/", handler.AuthMiddleware())
	roleProtected := router.Group("/", handler.AuthMiddlewareWithRoles([]string{"admin"}))
	dealerProtected := router.Group("/", handler.AuthMiddlewareWithRoles([]string{"dealer", "admin"}))

	// Update player bet
	// Update total pool
	// Update player turn

	// Users
	protected.GET("/api/user/:id", userHandler.GetUserByUsername)
	protected.GET("/api/user/:id/session", userHandler.GetSessionInfoByUsername)
	protected.GET("/api/user/:id/rank", userHandler.GetPlayerRankingByUsername)
	roleProtected.GET("/api/users", userHandler.GetUsers)
	roleProtected.POST("/api/user/:id/update", userHandler.UpdateUserByUsername)

	dealerProtected.POST("/api/user/:id/send", userHandler.TransferScoreByUsername)

	// Unprotected
	router.GET("/api/public/scoreboard", userHandler.WSUpdateScoreboard)
	router.GET("/api/users/scoreboard", userHandler.GetUsersScoreboard)

	// Tables
	protected.GET("/api/table/:table_name", tableHandler.GetTableByName)
	roleProtected.GET("/api/tables", tableHandler.GetTables)
	roleProtected.GET("/api/table/:table_name/sessions", tableSessionsHandler.GetTableSessions)
	roleProtected.POST("/api/table/:table_name/sessions/delete", tableSessionsHandler.DeleteTableSessionsByTableName)
	roleProtected.POST("/api/table/:table_name/session/create", tableSessionsHandler.CreateTableSession)
	roleProtected.POST("/api/table/:table_name/session/delete", tableSessionsHandler.DeleteTableSessionBySessionId)

	roleProtected.POST("/api/table/:table_name/session/:session_id/dealer/add", tableSessionsHandler.AddDealerToTableSession)
	roleProtected.POST("/api/table/:table_name/session/:session_id/dealer/remove", tableSessionsHandler.RemoveDealerFromTableSession)

	roleProtected.POST("/api/table/create", tableHandler.CreateTable)
	roleProtected.POST("/api/table/delete", tableHandler.DeleteTable)

	// Sessions
	dealerProtected.GET("/api/dealer/:id/session", tableSessionsHandler.GetTableSessionByDealer)
	dealerProtected.POST("/api/table/:table_name/session/:session_id/player/add", playerSessionsHandler.AddPlayerToPlayerSession)
	dealerProtected.POST("/api/table/:table_name/session/:session_id/status/update", tableSessionsHandler.UpdateTableSessionStatusBySessionId)
	dealerProtected.POST("/api/table/:table_name/session/:session_id/player/delete", playerSessionsHandler.DeletePlayerFromPlayerSession)

	roleProtected.GET("/api/table/:table_name/session/:session_id/players", playerSessionsHandler.GetPlayersForSessionId)

	router.Run()
	// log.Fatal(autotls.Run(router, "api.zikeeper.com", "zikeeper.com"))
}
