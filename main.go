package main

import (
	//"net/http"
	"hellogo/internal/db"
	"hellogo/internal/handler"
	"hellogo/internal/model"
	"hellogo/internal/repository"
	"hellogo/internal/service"
	"hellogo/validator"
	"log"
	"os"

	_ "hellogo/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Blog Aggregator API
// @version 1.0
// @description This is a sample RSS Blog Aggregator API
// @contact.name Rara Labs
// @contact.email hey@raralabs.com
// @host localhost:8080
// @BasePath /v1
func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT not found in the environment")
	}
	PORT = ":" + PORT

	// Connect to database
	db.Connect()
	// Run migrations
	db.DB.AutoMigrate(&model.Feed{})
	if err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}

	e := echo.New()
	//CORS ko middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// Allow all origins (use exact domains in production for security)
		AllowOrigins: []string{"*"},
		// Allowed HTTP methods
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// Allow any headers in requests
		AllowHeaders: []string{"*"},
		// Expose headers to the browser if needed
		ExposeHeaders: []string{"Link"},
		// Set to true if you need cookies/auth headers across domains
		AllowCredentials: false,
		// Preflight cache duration (in seconds)
		//MaxAge: 300,
	}))
	e.Validator = validator.NewValidator()

	feedrepo := &repository.FeedRepositoryGorm{DB: db.DB}     // Initialize repository
	feedservice := &service.FeedService{Repo: feedrepo}       // Initialize service
	feedHandler := &handler.FeedHandler{Service: feedservice} // Initialize handler

	v1 := e.Group("/v1")
	feeds := v1.Group(("/feeds"))

	// Routes
	feeds.GET("", feedHandler.ListFeed)
	feeds.GET("/:id", feedHandler.GetFeed)
	feeds.POST("", feedHandler.CreateFeed)
	feeds.PATCH("/:id", feedHandler.UpdateFeed)
	//feeds.PUT("/:id", handler.ReplaceFeed)
	feeds.DELETE("/:id", feedHandler.DeleteFeed)

	//Swagger documentation route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(PORT))
}
