package api

import (
	"log"
	"net/http"

	"ecom2.5/config"
	"ecom2.5/internal/api/rest"
	handlers "ecom2.5/internal/api/rest/handlers"
	"ecom2.5/internal/domain"
	"ecom2.5/internal/helper"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()
	log.Println("Configred")

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Database not connetion error %v\n", err)
	}

	log.Printf("Data base connected")
	// log.Print(db)

	// run migration
	db.AutoMigrate(&domain.User{})

	auth := helper.SetUpAuth(config.AppSecret)

	restHandler := &rest.RestHandler{
		App:  app,
		DB:   db,
		Auth: auth,
	}

	setUpRoutes(restHandler)

	app.Get("/health", HealthCheck)
	app.Listen(config.ServerPort)
}

func setUpRoutes(rh *rest.RestHandler) {
	// user handler
	handlers.SetUserRoutes(rh)
	//transaction

	// catalog
}

func HealthCheck(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "yay",
	})
}
