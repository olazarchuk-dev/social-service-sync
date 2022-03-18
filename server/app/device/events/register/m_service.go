package register

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"social-service-sync/server/model/api"
	"social-service-sync/server/model/entity"
	"strconv"
	"time"
)

func HandlerCreate(ctx context.Context, collection *mongo.Collection, request api.RegisterRequest) (string, error) {
	newUser := entity.NewUser(
		request.DeviceName,
		request.Email,
		request.Password,
		time.Now(),
		AddDate(0, 0, 7),
	)

	strNewUserId, err := CreateUser(ctx, collection, newUser)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("New UserID:", strNewUserId)
	return strNewUserId, err
}

func HandlerGet(ctx context.Context, collection *mongo.Collection, id string) (*entity.User, error) {
	user, err := GetUser(ctx, collection, id) // TODO: Repository
	if err != nil {
		log.Fatal(err)
	}
	Print(user)
	return &user, nil
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

func AddDate(years int, months int, days int) time.Time {
	return time.Now().AddDate(years, months, days)
}
