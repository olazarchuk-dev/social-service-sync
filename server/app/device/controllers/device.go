package controllers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"

	"github.com/olazarchuk-dev/go-social-service/server/app/device/events/login"
	"github.com/olazarchuk-dev/go-social-service/server/app/device/events/register"
	"github.com/olazarchuk-dev/go-social-service/server/app/device/events/ws"
	"github.com/olazarchuk-dev/go-social-service/server/app/device/events/ws/handler"
	"github.com/olazarchuk-dev/go-social-service/server/app/middleware"
)

func Init(app *fiber.App, db *sql.DB) {

	hub := ws.NewHub()
	go hub.Run()

	app.Post("/register", func(ctx *fiber.Ctx) error {
		return register.Handler(ctx, db)
	})

	app.Post("/login", func(ctx *fiber.Ctx) error {
		return login.Handler(ctx, db)
	})

	app.Post("/ws", middleware.JWTAuth, func(ctx *fiber.Ctx) error {
		return handler.CreateUser(ctx, hub)
	})

	app.Get("/ws/:username", handler.JoinUser(hub))

	app.Get("/ws", func(ctx *fiber.Ctx) error {
		return handler.GetAvailableUsers(ctx, hub)
	})

	app.Get("/ws/users/:username", func(ctx *fiber.Ctx) error {
		return handler.GetDevicesInUser(ctx, hub)
	})

}
