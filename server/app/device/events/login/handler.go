package login

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber/v2"
	"social-service-sync/server/model/api"
)

func Handler(ctx *fiber.Ctx, mongoDb *mongo.Database) error {
	ctxBg := context.TODO()
	user := new(api.LoginRequest)

	if err := ctx.BodyParser(user); err != nil {
		panic(err)
	}

	res := Service(mongoDb, ctxBg, *user)

	return ctx.JSON(res)
}
