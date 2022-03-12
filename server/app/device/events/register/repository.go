package register

import (
	"context"
	"database/sql"

	"github.com/olazarchuk-dev/go-social-service/server/model/entity"
)

func Repository(ctx context.Context, tx *sql.Tx, user entity.Users) (*entity.Users, error) {

	var lastInsertId int
	query := "INSERT INTO users(device_name, password, email, image) VALUES($1, $2, $3, $4) returning id"
	err := tx.QueryRowContext(ctx, query, user.DeviceName, user.Password, user.Email, user.Image).Scan(&lastInsertId)

	if err != nil {
		return &entity.Users{}, err
	}

	user.Id = lastInsertId
	return &user, nil

}
