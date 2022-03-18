package login

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"math/rand"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"social-service-sync/server/app/config"
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

func Service(db *mongo.Database, ctx context.Context, request api.LoginRequest) *api.LoginResponse {

	collection := db.Collection("users")

	result, err := RepositoryGet(ctx, collection, request.DeviceName) // TODO: Repository
	//results, err := RepositoryFindByName(ctx, collection, request.DeviceName) // TODO: Repository
	if err != nil {
		log.Fatal(err)
	}
	//for n, user := range results {
	//	entity.PrintUserList(n, user)
	//}

	var baseResponse api.BaseResponse
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

	errComparePass := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(request.Password))
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
		DeviceName: result.Username,
		Email:      result.Email,
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
		DeviceName: result.Username,
		Email:      result.Email,
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
