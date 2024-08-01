package domain

import (
	"context"

	"shellrean.id/golang_auth/dto"
)

type User struct {
	ID       int64  `db:"id"`
	FullName string `db:"full_name"`
	Phone    string `db:"phone"`
	UserName string `db:"user_name"`
	Password string `db:"password"`
}

type UserRepository interface {
	FindById(ctx context.Context, id int64) (User, error)
	FindByUserName(ctx context.Context, userName string) (User, error)
}

type UserService interface {
	Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthRes, error)
	ValidateToken(ctx context.Context, token string) (dto.UserData, error)
}
