package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"social-service-sync/server/app/device/events/ws"
)

func JoinUser(hub *ws.Hub) fiber.Handler {
	return websocket.New(func(conn *websocket.Conn) {
		id := conn.Query("id")
		username := conn.Params("username")
		deviceName := conn.Query("deviceName")
		fmt.Println(" ...JoinUser <<<", "id='"+id+"'", "username='"+username+"'", "deviceName='"+deviceName+"'")

		wsService := &ws.WsService{
			Conn:       conn,
			Id:         id,
			Username:   username,
			DeviceName: deviceName,
			Something:  make(chan *ws.Something, 10),
		}

		something := ws.Something{
			Id:               wsService.Id,
			Username:         wsService.Username,
			DeviceName:       wsService.DeviceName,
			SyncDeviceJoined: "joined_device", // TODO: [one special] sync joined device(s) by user
			//CurrentDevice:    nil,
			//LastModifiedAt:    0,
			AppUsername:      "",
			AppEmailAddress:  "",
			AppAlignedCb:     false,
			AppBillingPeriod: 3,
			AppSalary:        2500,
		}

		hub.Register <- wsService
		hub.Broadcast <- &something

		go wsService.WriteSomething()
		wsService.ReadSomething(hub)

	})
}
