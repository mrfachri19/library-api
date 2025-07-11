package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mrfachri19/digital-library-backend/internal/config"
	"github.com/mrfachri19/digital-library-backend/internal/handler"
	"github.com/mrfachri19/digital-library-backend/internal/middleware"
	"github.com/mrfachri19/digital-library-backend/internal/repository"
	"github.com/mrfachri19/digital-library-backend/internal/usecase"
)

func main() {
	// Load env & connect to DB
	config.LoadEnv()
	config.ConnectDB()
	config.RunMigration()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // frontend URL
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	// ========= Init Layer =========
	// Book
	bookRepo := repository.NewBookRepository()
	bookUC := usecase.NewBookUsecase(bookRepo)
	bookHandler := handler.NewBookHandler(bookUC)

	// Lending
	lendingRepo := repository.NewLendingRepository()
	lendingUC := usecase.NewLendingUsecase(lendingRepo)
	lendingHandler := handler.NewLendingHandler(lendingUC)

	// Analytics
	analyticsRepo := repository.NewAnalyticsRepository()
	analyticsUC := usecase.NewAnalyticsUsecase(analyticsRepo)
	analyticsHandler := handler.NewAnalyticsHandler(analyticsUC)

	// Auth
	authRepo := repository.NewAuthRepository()
	authUC := usecase.NewAuthUsecase(authRepo)
	authHandler := handler.NewAuthHandler(authUC)

	// ========= Public Routes =========
	app.Post("/auth/register", authHandler.Register)
	app.Post("/auth/login", authHandler.Login)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("📚 Digital Library Backend is running!")
	})

	// ========= Protected Routes =========
	api := app.Group("/api", middleware.Protected())

	// Books
	api.Get("/books", bookHandler.GetAll)
	api.Get("/books/:id", bookHandler.GetByID)
	api.Post("/books", bookHandler.Create)
	api.Put("/books/:id", bookHandler.Update)
	api.Delete("/books/:id", bookHandler.Delete)

	// Lendings
	api.Get("/lendings", lendingHandler.GetAll)
	api.Post("/lendings", lendingHandler.Create)
	api.Put("/lendings/:id/return", lendingHandler.MarkAsReturned)

	// Analytics
	api.Get("/analytics/summary", analyticsHandler.GetSummary)

	// Start server
	port := os.Getenv("APP_PORT")
	log.Fatal(app.Listen(":" + port))
}
