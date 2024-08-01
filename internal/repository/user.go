package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"shellrean.id/golang_auth/domain"
)

type userRepository struct {
	db *goqu.Database
}

func NewUser(con *sql.DB) domain.UserRepository {
	return &userRepository{
		db: goqu.New("postgres", con),
	}
}

func (u *userRepository) FindById(ctx context.Context, id int64) (user domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.Ex{"id": id})
	_, err = dataset.ScanStructContext(ctx, &user)
	return
}

func (u *userRepository) FindByUserName(ctx context.Context, userName string) (user domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.Ex{"user_name": userName})
	_, err = dataset.ScanStructContext(ctx, &user)
	return
}
