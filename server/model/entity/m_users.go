package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"social-service-sync/server/app/helper"
	"time"
)

//User is the structure we work with
type User struct {
	ID            primitive.ObjectID  `bson:"_id" json:"_id"`
	Username      string              `json:"username" gorm:"unique" bson:"username,omitempty"`
	Email         string              `json:"email" gorm:"unique" bson:"email,omitempty"`
	Password      string              `json:"password" bson:"password"`
	CreatedAt     primitive.Timestamp `json:"createdat" bson:"createat"`
	DeactivatedAt primitive.Timestamp `json:"updatedat" bson:"updatedat"`
}

// NewUser create new instance of User
func NewUser(username string, email string, password string, createdAt time.Time, deactivatedAt time.Time) User {
	bytes, errHash := bcrypt.GenerateFromPassword([]byte(password), 10)
	helper.PanicErr(errHash)

	user := User{}
	user.ID = primitive.NewObjectID()
	user.Username = username
	user.Email = email
	user.Password = string(bytes)
	user.CreatedAt = primitive.Timestamp{T: uint32(createdAt.Unix())}
	user.DeactivatedAt = primitive.Timestamp{T: uint32(deactivatedAt.Unix())}
	return user
}
