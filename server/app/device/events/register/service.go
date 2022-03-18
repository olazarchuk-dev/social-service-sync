package register

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"math/rand"
	"social-service-sync/server/model/api"
	"social-service-sync/server/model/entity"
	"time"
)

func Service(mongoDb *mongo.Database, ctx context.Context, request api.RegisterRequest) *api.RegisterResponse {

	collection := mongoDb.Collection("users")

	newUser := entity.NewUser(
		request.DeviceName,
		request.Email,
		request.Password,
		time.Now(),
		entity.AddDate(0, 0, 7),
	)

	id, err := RepositoryCreate(ctx, collection, newUser)
	if err != nil {
		log.Fatal(err)
	}

	result, err := RepositoryGet(ctx, collection, id) // TODO: Repository
	if err != nil {
		log.Fatal(err)
	}

	var baseResponse api.BaseResponse
	if err != nil {
		baseResponse = api.BaseResponse{
			Success: false,
			Code:    401,
			Message: err.Error(),
		}

		return &api.RegisterResponse{
			BaseResponse: &baseResponse,
		}
	}

	baseResponse = api.BaseResponse{
		Success: true,
		Code:    201,
		Message: "Register is success",
	}

	return &api.RegisterResponse{
		BaseResponse: &baseResponse,
		Data: &api.RegisterResponseData{
			//Id:         result.Id,
			//DeviceName: result.DeviceName,
			//Email:      request.Email,
			//Image:      result.Image,
			Id:         rand.Intn(100), // fmt.Sprintf("%v", user.ID.Hex()), // id
			DeviceName: result.Username,
			Email:      result.Email,
			Image:      "",
		},
	}
}
