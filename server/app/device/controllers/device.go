package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber/v2"

	"social-service-sync/server/app/device/events/login"
	"social-service-sync/server/app/device/events/register"
	"social-service-sync/server/app/device/events/ws"
	"social-service-sync/server/app/device/events/ws/handler"
	"social-service-sync/server/app/middleware"
)

func Init(app *fiber.App, db *mongo.Database) {

	m := ws.NewMongo(db)
	hub := ws.NewHub()
	go hub.Run(m)

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
