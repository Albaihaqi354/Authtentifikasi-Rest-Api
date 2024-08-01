package dto

type AuthReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
