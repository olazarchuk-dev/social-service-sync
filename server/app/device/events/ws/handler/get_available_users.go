package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"social-service-sync/server/app/device/events/ws"
	"social-service-sync/server/model/api"
)

type Res struct {
	*api.BaseResponse
	Data *[]Users `json:"data"`
}

type Users struct { // TODO: JoinUsers
	Username string `json:"username"`
}

func GetAvailableUsers(ctx *fiber.Ctx, hub *ws.Hub) error {
	users := make([]Users, 0)

	for _, user := range hub.Users {
		fmt.Println("<<<", "Username='"+user.Username+"'")
		users = append(users, Users{
			Username: user.Username,
		})
	}
	fmt.Println(" ...GetAvailableUsers <<<", users)

	res := Res{
		BaseResponse: &api.BaseResponse{
			Success: true,
			Code:    200,
			Message: "Success get users",
		},
		Data: &users,
	}

	return ctx.JSON(&res)
}
