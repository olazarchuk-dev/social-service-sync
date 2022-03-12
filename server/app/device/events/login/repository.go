package login

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/olazarchuk-dev/go-social-service/server/model/api"
	"log"

	"github.com/olazarchuk-dev/go-social-service/server/app/helper"
	"github.com/olazarchuk-dev/go-social-service/server/model/entity"
)

func Repository(ctx context.Context, tx *sql.Tx, request api.LoginRequest) (*entity.Users, error) {

	//	fmt.Println(deviceName)
	query := "SELECT * FROM users WHERE device_name = $1"
	rows, err := tx.QueryContext(ctx, query, request.DeviceName)

	helper.PanicErr(err)

	var user entity.Users

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.DeviceName, &user.Password, &user.Email, &user.Image)
		if err != nil {
			fmt.Println(err)
			return new(entity.Users), err
		}
		log.Print(user)
	}
	return &user, nil

}
