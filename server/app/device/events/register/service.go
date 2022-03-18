package register

import (
	"context"
	"database/sql"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"social-service-sync/server/model/api"
)

func Service(db *sql.DB, mongoDb *mongo.Database, ctx context.Context, request api.RegisterRequest) *api.RegisterResponse {

	//tx, err := db.Begin()
	//helper.PanicErr(err)
	//defer helper.RollbackErr(tx)
	//
	//bytes, errHash := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	//helper.PanicErr(errHash)
	//
	//user := entity.Users{
	//	DeviceName: request.DeviceName,
	//	Password:   string(bytes),
	//	Email:      request.Email,
	//	Image:      request.Image,
	//}
	//
	//result, errQuery := Repository(ctx, tx, user)
	var baseResponse api.BaseResponse

	//if errQuery != nil {
	//	if strings.Contains(errQuery.Error(), "duplicate") {
	//		baseResponse = api.BaseResponse{
	//			Success: false,
	//			Code:    400,
	//			Message: "Users was registered",
	//		}
	//
	//		return &api.RegisterResponse{
	//			BaseResponse: &baseResponse,
	//		}
	//
	//	}
	//
	//	baseResponse = api.BaseResponse{
	//		Success: false,
	//		Code:    500,
	//		Message: "Something wrong",
	//	}
	//
	//	return &api.RegisterResponse{
	//		BaseResponse: &baseResponse,
	//	}
	//}

	//
	collection := mongoDb.Collection("users")

	id, err := HandlerCreate(ctx, collection, request)
	user, err := HandlerGet(ctx, collection, id)
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
	//

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

			Id:         rand.Intn(100), // fmt.Sprintf("%v", user.ID.Hex()),
			DeviceName: user.Username,
			Email:      user.Email,
			Image:      "",
		},
	}

}
