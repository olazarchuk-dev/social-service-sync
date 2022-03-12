package register

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/olazarchuk-dev/go-social-service/server/model/api"
)

func Handler(ctx *fiber.Ctx, db *sql.DB) error {

	ctxBg := context.Background()
	user := new(api.RegisterRequest)

	if err := ctx.BodyParser(user); err != nil {
		panic(err)
	}

	res := Service(db, ctxBg, *user)

	return ctx.JSON(res)

}
