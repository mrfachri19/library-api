package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mrfachri19/digital-library-backend/internal/config"
	"github.com/mrfachri19/digital-library-backend/internal/handler"
	"github.com/mrfachri19/digital-library-backend/internal/repository"
	"github.com/mrfachri19/digital-library-backend/internal/usecase"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	config.RunMigration()

	app := fiber.New()

	// Init layer
	bookRepo := repository.NewBookRepository()
	bookUC := usecase.NewBookUsecase(bookRepo)
	bookHandler := handler.NewBookHandler(bookUC)

	lendingRepo := repository.NewLendingRepository()
	lendingUC := usecase.NewLendingUsecase(lendingRepo)
	lendingHandler := handler.NewLendingHandler(lendingUC)

	// Lending routes
	app.Get("/lendings", lendingHandler.GetAll)
	app.Post("/lendings", lendingHandler.Create)
	app.Put("/lendings/:id/return", lendingHandler.MarkAsReturned)

	// Book routes
	app.Get("/books", bookHandler.GetAll)
	app.Get("/books/:id", bookHandler.GetByID)
	app.Post("/books", bookHandler.Create)
	app.Put("/books/:id", bookHandler.Update)
	app.Delete("/books/:id", bookHandler.Delete)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Digital Library Backend is running!")
	})

	port := os.Getenv("APP_PORT")
	log.Fatal(app.Listen(":" + port))
}
