package entity

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"social-service-sync/server/app/helper"
	"strconv"
	"time"
)

type User struct {
	ID            primitive.ObjectID  `bson:"_id" json:"_id"`
	Username      string              `json:"username" gorm:"unique" bson:"username,omitempty"`
	Email         string              `json:"email" gorm:"unique" bson:"email,omitempty"`
	Password      string              `json:"password" bson:"password"`
	CreatedAt     primitive.Timestamp `json:"createdat" bson:"createat"`
	DeactivatedAt primitive.Timestamp `json:"updatedat" bson:"updatedat"`
}

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

func AddDate(years int, months int, days int) time.Time {
	return time.Now().AddDate(years, months, days)
}

func PrintUser(user User) {
	ID := user.ID.Hex()
	Username := user.Username
	Email := user.Email
	Password := user.Password
	CreatedAt := ToString(user.CreatedAt, time.Stamp)
	DeactivatedAt := ToString(user.DeactivatedAt, time.Stamp)
	fmt.Printf("\nID='%v'; Username='%v'; Email='%v'; Password='%v'; CreatedAt='%v'; DeactivatedAt='%v'; \n\n",
		ID, Username, Email, Password, CreatedAt, DeactivatedAt)
}

func PrintUserList(u int, user User) {
	ID := user.ID.Hex()
	Username := user.Username
	Email := user.Email
	Password := user.Password
	CreatedAt := ToString(user.CreatedAt, time.Stamp)
	DeactivatedAt := ToString(user.DeactivatedAt, time.Stamp)
	fmt.Printf("%v. ID='%v'; Username='%v'; Email='%v'; Password='%v'; CreatedAt='%v'; DeactivatedAt='%v'; \n",
		u, ID, Username, Email, Password, CreatedAt, DeactivatedAt)
}

func ToString(date primitive.Timestamp, layout string) string {
	uintDate := strconv.FormatUint(uint64(date.T), 10)
	intDate, err := strconv.ParseInt(uintDate, 10, 64)

	if err != nil {
		panic(err)
	}
	unixDate := time.Unix(intDate, 0)

	return unixDate.Format(layout)
}
