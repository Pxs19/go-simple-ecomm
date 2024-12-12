package handlers

// import (
// 	"net/http"

// 	"ecom2.5/internal/api/rest"
// 	"ecom2.5/internal/service"
// 	"github.com/gofiber/fiber/v2"
// )

// type UserHandler struct {
// 	// pass service(svc) here
// 	svc service.UserService
// }

// func SetUserRoutes(rh *rest.RestHandler) {

// 	app := rh.App

// 	// create an instance of user and inject to handler
// 	svc := service.UserService{}
// 	handler := UserHandler{
// 		svc: svc,
// 	}

// 	_ = svc

// 	// Public endpoints
// 	app.Post("/register", handler.Register)
// 	app.Post("/login", handler.Login)

// 	// Private endpoint
// 	app.Get("/verify", handler.GetVerification)
// 	app.Post("/verify", handler.Verify)

// 	app.Get("/profile", handler.GetProfile)
// 	app.Post("/profile", handler.CreateProfile)

// 	app.Get("/cart", handler.GetCart)
// 	app.Post("/cart", handler.AddToCart)

// 	app.Get("/order", handler.GetOrders)
// 	app.Get("/order/:id", handler.GetOrder)

// 	app.Post("/become-seller", handler.BecomeSeller)

// }

// func (h *UserHandler) Register(ctx *fiber.Ctx) error {

// 	h.svc.

// 	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"mmessage": "register",
// 	})
// }

// func (h *UserHandler) Login(ctx *fiber.Ctx) error {

// 	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"mmessage": "login",
// 	})
// }

// func (h *UserHandler) GetVerification(ctx *fiber.Ctx) error {

// 	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"mmessage": "getVerification",
// 	})
// }

// func (h *UserHandler) Verify(ctx *fiber.Ctx) error {

// 	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"mmessage": "verify",
// 	})
// }

// func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {

// 	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"mmessage": "createProfile",
// 	})
// }

// func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {

// 	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"mmessage": "getProfile",
// 	})
// }

// func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {

// 	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"mmessage": "addToCart",
// 	})
// }

// func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {

// 	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"mmessage": "getCart",
// 	})
// }

// func (h *UserHandler) CreateOrder(ctx *fiber.Ctx) error {

// 	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"mmessage": "createOrder",
// 	})
// }

// func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {

// 	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"mmessage": "getOrders",
// 	})
// }

// func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {

// 	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"mmessage": "getOrder",
// 	})
// }

// func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {

// 	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"mmessage": "becomeSeller",
// 	})
// }
