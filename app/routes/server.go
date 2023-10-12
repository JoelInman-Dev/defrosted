package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
)
var (
	store *session.Store
	AUTH_KEY string = "secret_key"
	USER_ID string = "user_id"
	PLAYER string = "none"
)

func Setup() {
	app := fiber.New()

	store = session.New(session.Config{
		CookieHTTPOnly: true,
		Expiration: time.Hour * 120,
	})

	app.Use(NewMiddleware(), cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins: "*", // change this to your production domain
		AllowHeaders: "Access-Control-Allow-Origin, Content-Type, Content-Length, Origin, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))
	// Auth Routes

	app.Post("/auth/login", Login)
	app.Post("/auth/register", Register)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":5000")
}