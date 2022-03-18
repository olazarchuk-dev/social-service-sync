package login

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber/v2"
	"social-service-sync/server/model/api"
)

func Handler(ctx *fiber.Ctx, db *mongo.Database) error {
	ctxBg := context.TODO()
	user := new(api.LoginRequest)

	if err := ctx.BodyParser(user); err != nil {
		panic(err)
	}

	res := Service(db, ctxBg, *user)

	return ctx.JSON(res)
}
