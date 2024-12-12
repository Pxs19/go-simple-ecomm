package handlers

import (
	"net/http"

	"ecom2.5/internal/api/rest"
	"ecom2.5/internal/dto"
	"ecom2.5/internal/helper"
	"ecom2.5/internal/repository"
	"ecom2.5/internal/service"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	// pass service(svc) here
	svc service.UserService

	Auth helper.Auth
}

func SetUserRoutes(rh *rest.RestHandler) {

	app := rh.App

	// create an instance of user and inject to handler
	svc := service.UserService{
		Repo: repository.NewUserRepository(rh.DB),
		Auth: rh.Auth,
	}

	handler := UserHandler{
		svc: svc,
	}

	// _ = svc

	// ** pub = public, pvt = private
	pubRoute := app.Group("/api/v1")
	// Public endpoints
	pubRoute.Post("/register", handler.Register)
	pubRoute.Post("/login", handler.Login)

	pvtRoute := pubRoute.Group("/users", rh.Auth.Authorization)
	// Private endpoint
	pvtRoute.Get("/verify", handler.GetVerification)
	pvtRoute.Post("/verify", handler.Verify)

	pvtRoute.Get("/profile", handler.GetProfile)
	pvtRoute.Post("/profile", handler.CreateProfile)

	pvtRoute.Get("/cart", handler.GetCart)
	pvtRoute.Post("/cart", handler.AddToCart)

	pvtRoute.Get("/order", handler.GetOrders)
	pvtRoute.Get("/order/:id", handler.GetOrder)

	pvtRoute.Post("/become-seller", handler.BecomeSeller)

}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	// create user
	user := dto.UserSignup{}
	err := ctx.BodyParser(&user)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"mmessage": "please provide valids input",
		})
	}
	token, err := h.svc.SignUp(user)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"mmessage": "error on signup",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message":      "register successful",
		"status":       200,
		"access_token": token,
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {

	loginInput := dto.UsersLogin{}
	err := ctx.BodyParser(&loginInput)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"mmessage": "please provide valids input",
		})
	}

	token, err := h.svc.Login(loginInput.Email, loginInput.Password)

	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"mmessage": "please provide correct user id password",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mmessage": "login",
		"token":    token,
	})
}

func (h *UserHandler) GetVerification(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mmessage": "getVerification",
	})
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mmessage": "verify",
	})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mmessage": "createProfile",
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {

	user := h.svc.Auth.GetCurrentUser(ctx)

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "getProfile",
		"user":    user,
	})
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mmessage": "addToCart",
	})
}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mmessage": "getCart",
	})
}

func (h *UserHandler) CreateOrder(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mmessage": "createOrder",
	})
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mmessage": "getOrders",
	})
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mmessage": "getOrder",
	})
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"mmessage": "becomeSeller",
	})
}
