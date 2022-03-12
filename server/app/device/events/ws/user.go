package ws

type User struct { // TODO: JoinUser
	Username   string                `json:"username"`
	WsServices map[string]*WsService `json:"wsServices"`
}
