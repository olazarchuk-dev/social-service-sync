package register

import (
	"context"
	"database/sql"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber/v2"
	"social-service-sync/server/model/api"
)

func Handler(ctx *fiber.Ctx, db *sql.DB, mongoDb *mongo.Database) error {
	//ctxBg := context.Background()
	ctxBg := context.TODO()
	user := new(api.RegisterRequest)

	if err := ctx.BodyParser(user); err != nil {
		panic(err)
	}

	res := Service(db, mongoDb, ctxBg, *user)

	return ctx.JSON(res)
}
