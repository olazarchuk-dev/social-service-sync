package login

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"social-service-sync/server/model/api"
)

func Handler(ctx *fiber.Ctx, db *sql.DB) error {
	ctxBg := context.Background()
	user := new(api.LoginRequest)

	if err := ctx.BodyParser(user); err != nil {
		panic(err)
	}

	res := Service(db, ctxBg, *user)
	//
	HandlerGet(*user)
	//

	return ctx.JSON(res)
}
