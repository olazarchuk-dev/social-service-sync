package ws

import (
	"fmt"
	"strconv"
)

type Hub struct {
	Register   chan *WsService
	Unregister chan *WsService
	Broadcast  chan *Something
	Users      map[string]*User // TODO: JoinUser
}

func (hub *Hub) Run() {
	for {
		select {
		// Register case
		case wsService := <-hub.Register:
			if _, isUsernameExist := hub.Users[wsService.Username]; !isUsernameExist {
				fmt.Println(" ...Hub.Register: isUsernameExist <<<", wsService)
				hub.Users[wsService.Username] = &User{
					Username:   wsService.Username,
					WsServices: make(map[string]*WsService),
				}
			}
			user := hub.Users[wsService.Username]
			if _, isIdExist := user.WsServices[wsService.Id]; !isIdExist {
				fmt.Println(" ...Hub.Register: isIdExist <<<", wsService)
				user.WsServices[wsService.Id] = wsService
			}

		// Unregister case
		case wsService := <-hub.Unregister:
			if _, isWsServiceExist := hub.Users[wsService.Username].WsServices[wsService.Id]; isWsServiceExist {
				fmt.Println(" ...Hub.Unregister: delete Connection")
				if len(hub.Users[wsService.Username].WsServices) != 0 {
					hub.Broadcast <- &Something{
						Id:          wsService.Id,
						Username:    wsService.Username,
						DeviceName:  wsService.DeviceName,
						AppUsername: "disjoined_device", // TODO: [special] sync disjoined device(s) by user
					}
				}
				delete(hub.Users[wsService.Username].WsServices, wsService.Id)
				close(wsService.Something)
			}

			// remove device if no one wsService
			wsServices := hub.Users[wsService.Username].WsServices
			if len(wsServices) == 0 {
				delete(hub.Users, wsService.Username)
			}

		// Broadcast case
		case something := <-hub.Broadcast:
			if _, exist := hub.Users[something.Username]; exist {
				for _, wsService := range hub.Users[something.Username].WsServices {
					if wsService.Username == something.Username {
						wsService.Something <- something // TODO: Websocket.Connect: Id, Username, DeviceName;  ||  App (settings): AppUsername, AppEmailAddress, AppAlignedCb, AppBillingPeriod, AppSalary;
						fmt.Println(" ...Hub.Broadcast something <<<",
							"(Conn) Id='"+something.Id+"'",
							"(Conn) Username='"+something.Username+"'",
							"(Conn) DeviceName='"+something.DeviceName+"'",
							"(App) Username='"+something.AppUsername+"'",
							"(App) EmailAddress='"+something.AppEmailAddress+"'",
							"(App) AlignedCb='"+strconv.FormatBool(something.AppAlignedCb)+"'",
							"(App) BillingPeriod='"+strconv.Itoa(something.AppBillingPeriod)+"'",
							"(App) Salary='"+strconv.Itoa(something.AppSalary)+"'")
					}
				}
			}
		}
	}
}

func NewHub() *Hub {
	return &Hub{
		Register:   make(chan *WsService),
		Unregister: make(chan *WsService),
		Broadcast:  make(chan *Something, 5),
		Users:      make(map[string]*User),
	}
}
