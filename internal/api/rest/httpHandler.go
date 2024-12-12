package rest

import (
	"ecom2.5/internal/helper"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestHandler struct {
	App  *fiber.App
	DB   *gorm.DB
	Auth helper.Auth
}
