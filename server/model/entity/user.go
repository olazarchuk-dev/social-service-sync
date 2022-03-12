package entity

type Users struct {
	Id         int    `json:"id" db:"id"`
	DeviceName string `json:"deviceName" db:"device_name"`
	Password   string `json:"password" db:"password"`
	Email      string `json:"email" db:"email"`
	Image      string `json:"image" db:"image"`
}
