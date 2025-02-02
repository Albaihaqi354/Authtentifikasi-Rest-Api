package service

import (
	"context"
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
	"shellrean.id/golang_auth/domain"
	"shellrean.id/golang_auth/dto"
	"shellrean.id/golang_auth/internal/util"
)

type userService struct {
	userRepository  domain.UserRepository
	cacheRepository domain.CacheRepository
}

func NewUser(userRepository domain.UserRepository, cacheRepository domain.CacheRepository) domain.UserService {
	return &userService{
		userRepository:  userRepository,
		cacheRepository: cacheRepository,
	}

}

func (u *userService) Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthRes, error) {
	user, err := u.userRepository.FindByUserName(ctx, req.UserName)
	if err != nil {
		return dto.AuthRes{}, err
	}
	if user == (domain.User{}) {
		return dto.AuthRes{}, domain.ErrAuthFailed
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthRes{}, domain.ErrAuthFailed
	}

	token := util.GenerateRandomString(16)

	userJson, _ := json.Marshal(user)
	_ = u.cacheRepository.Set("user:"+token, userJson)

	return dto.AuthRes{
		Token: token,
	}, nil
}

func (u *userService) ValidateToken(ctx context.Context, token string) (dto.UserData, error) {
	data, err := u.cacheRepository.Get("user:" + token)
	if err != nil {
		return dto.UserData{}, domain.ErrAuthFailed
	}
	var user domain.User
	_ = json.Unmarshal(data, &user)

	return dto.UserData{
		ID:       user.ID,
		FullName: user.FullName,
		Phone:    user.Phone,
		UserName: user.UserName,
	}, nil
}
