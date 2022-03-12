package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/olazarchuk-dev/go-social-service/server/app/device/events/ws"
	"github.com/olazarchuk-dev/go-social-service/server/model/api"
)

type WsServices struct {
	*api.BaseResponse
	Data []WsServiceInUser `json:"data"`
}

type WsServiceInUser struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	DeviceName string `json:"deviceName"`
}

func GetDevicesInUser(ctx *fiber.Ctx, hub *ws.Hub) error {

	var wsServices []WsServiceInUser
	username := ctx.Params("username")
	fmt.Println(" ...GetDevicesInUser <<<", "username='"+username+"'")

	if _, isExist := hub.Users[username]; !isExist {
		res := WsServices{
			BaseResponse: &api.BaseResponse{
				Success: true,
				Code:    200,
				Message: "No webSocket-service",
			},
			Data: make([]WsServiceInUser, 0),
		}
		return ctx.JSON(res)
	}

	if len(hub.Users[username].WsServices) == 0 {
		res := WsServices{
			BaseResponse: &api.BaseResponse{
				Success: true,
				Code:    200,
				Message: "No webSocket-service",
			},
			Data: make([]WsServiceInUser, 0),
		}
		return ctx.JSON(res)
	}

	for _, wsService := range hub.Users[username].WsServices {
		fmt.Println("3...isExist hub.Users <<<", wsService)
		wsServices = append(wsServices, WsServiceInUser{
			Id:         wsService.Id,
			Username:   wsService.Username,
			DeviceName: wsService.DeviceName,
		})
	}

	fmt.Println(" ...GetDevicesInUser <<<", wsServices)

	res := WsServices{
		BaseResponse: &api.BaseResponse{
			Success: true,
			Code:    200,
			Message: "Success get webSocket-service this device",
		},
		Data: wsServices,
	}

	return ctx.JSON(res)
}
