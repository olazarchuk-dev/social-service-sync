package api

type LoginRequest struct {
	DeviceName string `json:"deviceName" db:"device_name"`
	Password   string `json:"password" db:"password"`
}
