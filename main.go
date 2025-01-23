package main

import (
	"database/sql"
	"fmt"
	"log"
	"rbac/config"
	"rbac/handlers"
	"rbac/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int
	Username string
	Password string
}

type Role struct {
	ID   int
	Name string
}

type Permission struct {
	ID   int
	Name string
}


func getDB() (*sql.DB, error) {
	dbConfig, err := config.LoadDBConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load database config: %w", err)
	}

	db, err := sql.Open("mysql", dbConfig.GetDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}


func setupPublicRoutes(api *gin.RouterGroup, userHandler *handlers.UserHandler, authHandler *handlers.AuthHandler) {
	api.POST("/users", userHandler.CreateUser)
	api.POST("/login", authHandler.Login)
	api.POST("/refresh", authHandler.RefreshToken)

}


func setupProtectedRoutes(protected *gin.RouterGroup, userHandler *handlers.UserHandler, roleHandler *handlers.RoleHandler, authMiddleware *middleware.AuthMiddleware) {

	users := protected.Group("/users")
	{
		users.GET("", userHandler.GetUsers)
		users.GET("/:id", userHandler.GetUser)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}


	roles := protected.Group("/roles")
	{
		roles.GET("", roleHandler.GetRoles)
		roles.POST("", roleHandler.CreateRole)
		roles.GET("/:id", roleHandler.GetRole)
		roles.PUT("/:id", roleHandler.UpdateRole)
		roles.DELETE("/:id", roleHandler.DeleteRole)
	}


	protected.GET("/protected", authMiddleware.RequirePermission("view_post"), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Protected content"})
	})
}


func setupRouter(db *sql.DB) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)


	router := gin.Default()


	userHandler := handlers.NewUserHandler(db)
	roleHandler := handlers.NewRoleHandler(db)
	authHandler := handlers.NewAuthHandler(db)
	authMiddleware := middleware.NewAuthMiddleware(db)


	api := router.Group("/api")
	{
		setupPublicRoutes(api, userHandler, authHandler)
		
		protected := api.Group("")
		protected.Use(authMiddleware.Authenticate())
		setupProtectedRoutes(protected, userHandler, roleHandler, authMiddleware)
	}

	return router
}


func initializeApp() (*gin.Engine, *sql.DB, error) {
	db, err := getDB()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	router := setupRouter(db)

	return router, db, nil
}

func main() {
	router, db, err := initializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}
	defer db.Close()

	serverAddr := ":8080"
	log.Printf("Server starting on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 