package register

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"social-service-sync/server/model/api"
	"social-service-sync/server/model/entity"
	"strconv"
	"time"
)

func HandlerCreate(request api.RegisterRequest) (string, error) {
	Setup()

	newUser := entity.NewUser(
		request.DeviceName,
		request.Email,
		request.Password,
		time.Now(),
		AddDate(0, 0, 7),
	)

	strNewUserId, err := CreateUser(newUser)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("New UserID:", strNewUserId)
	return strNewUserId, err
}

func HandlerGet(id string) (*entity.User, error) {
	Setup()

	user, err := GetUser(id) // TODO: Repository
	if err != nil {
		log.Fatal(err)
	}
	Print(user)
	return &user, nil
}

//func HandlerGets() {
//	Setup()
//
//	users, err := GetUsers()
//	if err != nil {
//		log.Fatal(err)
//	}
//	for u, user := range users {
//		PrintList(u, user)
//	}
//}

func ToString(date primitive.Timestamp, layout string) string {
	uintDate := strconv.FormatUint(uint64(date.T), 10)
	intDate, err := strconv.ParseInt(uintDate, 10, 64)

	if err != nil {
		panic(err)
	}
	unixDate := time.Unix(intDate, 0)

	return unixDate.Format(layout)
}

func Print(user entity.User) {
	ID := user.ID.Hex()
	Username := user.Username
	Email := user.Email
	Password := user.Password
	CreatedAt := ToString(user.CreatedAt, time.Stamp)
	DeactivatedAt := ToString(user.DeactivatedAt, time.Stamp)
	fmt.Printf("\nID='%v'; Username='%v'; Email='%v'; Password='%v'; CreatedAt='%v'; DeactivatedAt='%v'; \n\n",
		ID, Username, Email, Password, CreatedAt, DeactivatedAt)
}

func PrintList(u int, user entity.User) {
	ID := user.ID.Hex()
	Username := user.Username
	Email := user.Email
	Password := user.Password
	CreatedAt := ToString(user.CreatedAt, time.Stamp)
	DeactivatedAt := ToString(user.DeactivatedAt, time.Stamp)
	fmt.Printf("%v. ID='%v'; Username='%v'; Email='%v'; Password='%v'; CreatedAt='%v'; DeactivatedAt='%v'; \n",
		u, ID, Username, Email, Password, CreatedAt, DeactivatedAt)
}

func AddDate(years int, months int, days int) time.Time {
	return time.Now().AddDate(years, months, days)
}
