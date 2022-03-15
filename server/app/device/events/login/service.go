package login

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"social-service-sync/server/app/config"
	"social-service-sync/server/app/helper"
	"social-service-sync/server/model/api"
)

type JwtClaims struct { // TODO: JwtClaims User
	jwt.StandardClaims
	Id         int    `json:"id"`
	Email      string `json:"email"`
	DeviceName string `json:"deviceName"`
}

var (
	APP_NAME                 = "Social-Service"
	JWT_SIGNING_METHOD       = jwt.SigningMethodHS256
	JWT_ACESS_TOKEN_EXPIRED  = time.Duration(1) * time.Hour
	JWT_ACCESS_TOKEN_EXPIRED = time.Duration(30) * (time.Hour * 24)
)

func Service(db *sql.DB, ctx context.Context, request api.LoginRequest) *api.LoginResponse {

	tx, err := db.Begin()
	helper.PanicErr(err)
	defer helper.RollbackErr(tx)

	var baseResponse api.BaseResponse
	//result, errQuery := Repository(ctx, tx, request)
	//
	//if errQuery != nil {
	//	if strings.Contains(errQuery.Error(), "found") {
	//		baseResponse = api.BaseResponse{
	//			Success: false,
	//			Code:    401,
	//			Message: errQuery.Error(),
	//		}
	//
	//		return &api.LoginResponse{
	//			BaseResponse: &baseResponse,
	//		}
	//	}
	//
	//	baseResponse = api.BaseResponse {
	//		Success: false,
	//		Code:    401,
	//		Message: "error when query to database",
	//	}
	//
	//	return &api.LoginResponse{
	//		BaseResponse: &baseResponse,
	//	}
	//}

	//
	user, err := HandlerGet(request)
	if err != nil {
		baseResponse = api.BaseResponse{
			Success: false,
			Code:    401,
			Message: err.Error(),
		}

		return &api.LoginResponse{
			BaseResponse: &baseResponse,
		}
	}

	errComparePass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	//
	//errComparePass := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(request.Password))
	if errComparePass != nil {
		return &api.LoginResponse{
			BaseResponse: &api.BaseResponse{
				Success: false,
				Code:    401,
				Message: "Password are invalid",
			},
		}
	}

	baseResponse = api.BaseResponse{
		Success: true,
		Code:    200,
		Message: "Login is success",
	}

	accessTokenClaims := JwtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APP_NAME,
			ExpiresAt: time.Now().Add(JWT_ACCESS_TOKEN_EXPIRED).Unix(),
		},
		//DeviceName: result.DeviceName,
		//Email:      result.Email,
		//Id:         result.Id,
		DeviceName: user.Username,
		Email:      user.Email,
		Id:         rand.Intn(100), // fmt.Sprintf("%v", user.ID.Hex()),
	}

	refreshTokenClaims := JwtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APP_NAME,
			ExpiresAt: time.Now().Add(JWT_ACCESS_TOKEN_EXPIRED).Unix(),
		},
		//DeviceName: result.DeviceName,
		//Email:      result.Email,
		//Id:         result.Id,
		DeviceName: user.Username,
		Email:      user.Email,
		Id:         rand.Intn(100), // fmt.Sprintf("%v", user.ID.Hex()),
	}

	accessToken := jwt.NewWithClaims(JWT_SIGNING_METHOD, accessTokenClaims)
	refreshToken := jwt.NewWithClaims(JWT_SIGNING_METHOD, refreshTokenClaims)

	config, errConfig := config.LoadConfig()
	if errConfig != nil {
		return &api.LoginResponse{
			BaseResponse: &api.BaseResponse{
				Success: false,
				Code:    401,
				Message: err.Error(),
			},
		}
	}

	signedAccessToken, errSignedToken := accessToken.SignedString([]byte(config.JwtSecretKey))
	signedrefreshToken, errSignedToken := refreshToken.SignedString([]byte(config.JwtSecretKey))

	if errSignedToken != nil {
		fmt.Println(errSignedToken)
		panic(errSignedToken)
	}

	return &api.LoginResponse{
		BaseResponse: &baseResponse,
		Data: &api.LoginResponseData{
			AccessToken:  signedAccessToken,
			RefreshToken: signedrefreshToken,
		},
	}

}
