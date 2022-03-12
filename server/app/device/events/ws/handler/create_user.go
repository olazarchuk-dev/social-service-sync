package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/olazarchuk-dev/go-social-service/server/app/device/events/ws"
)

type User struct { // TODO: JoinUser
	Username string
}

func CreateUser(ctx *fiber.Ctx, hub *ws.Hub) error {

	user := new(User)

	if err := ctx.BodyParser(user); err != nil {
		panic(err)
	}

	hub.Users[user.Username] = &ws.User{
		Username:   user.Username,
		WsServices: make(map[string]*ws.WsService),
	}
	fmt.Println(" ...CreateUser <<<", "Username='"+user.Username+"'")

	return ctx.JSON(user)

}

//
