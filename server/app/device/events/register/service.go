package register

import (
	"context"
	"database/sql"
	"strings"

	"github.com/olazarchuk-dev/go-social-service/server/app/helper"
	"github.com/olazarchuk-dev/go-social-service/server/model/api"
	"github.com/olazarchuk-dev/go-social-service/server/model/entity"
	"golang.org/x/crypto/bcrypt"
)

func Service(db *sql.DB, ctx context.Context, request api.RegisterRequest) *api.RegisterResponse {

	tx, err := db.Begin()
	helper.PanicErr(err)
	defer helper.RollbackErr(tx)

	bytes, errHash := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	helper.PanicErr(errHash)

	user := entity.Users{
		DeviceName: request.DeviceName,
		Password:   string(bytes),
		Email:      request.Email,
		Image:      request.Image,
	}

	result, errQuery := Repository(ctx, tx, user)
	var baseResponse api.BaseResponse

	if errQuery != nil {
		if strings.Contains(errQuery.Error(), "duplicate") {

			baseResponse = api.BaseResponse{
				Success: false,
				Code:    400,
				Message: "Users was registered",
			}

			return &api.RegisterResponse{
				BaseResponse: &baseResponse,
			}

		}

		baseResponse = api.BaseResponse{
			Success: false,
			Code:    500,
			Message: "Something wrong",
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
			Id:         result.Id,
			DeviceName: result.DeviceName,
			Email:      request.Email,
			Image:      result.Image,
		},
	}

}
