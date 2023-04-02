package types

import "time"

type UserLoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserInfo struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccountType int64  `json:"account_type"`
}

type UserLoginResponse struct {
	AccessToken      string    `json:"access_token"`
	AccessExpiredAt  time.Time `json:"access_expired_at"`
	RefreshToken     string    `json:"refresh_token"`
	RefreshExpiredAt time.Time `json:"refresh_expired_at"`
	UserInfo         UserInfo  `json:"user_info"`
}
