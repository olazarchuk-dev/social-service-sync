package middleware

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/olazarchuk-dev/go-social-service/server/app/config"
	"github.com/olazarchuk-dev/go-social-service/server/model/api"
)

func JWTAuth(ctx *fiber.Ctx) error {

	authHeader := ctx.Get("Authorization", "")

	if !strings.Contains(authHeader, "Bearer") || authHeader == "" {
		return ctx.JSON(api.BaseResponse{
			Success: false,
			Code:    401,
			Message: "invalid headers authorization",
		})
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

	config, errConfig := config.LoadConfig()
	if errConfig != nil {
		return ctx.JSON(api.BaseResponse{
			Success: false,
			Code:    400,
			Message: "error when read config",
		})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Signing method invalid")
		}
		return []byte(config.JwtSecretKey), nil
	})

	if err != nil {
		return ctx.JSON(api.BaseResponse{
			Success: false,
			Code:    400,
			Message: "Invalid jsonwebtoken",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return ctx.JSON(api.BaseResponse{
			Success: false,
			Code:    400,
			Message: "jsonapitoken invalid signature",
		})
	}

	ctx.Locals("jwt", claims)

	return ctx.Next()

}
