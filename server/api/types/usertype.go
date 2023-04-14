package types

type User struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccountType int64  `json:"account_type"`
}

type (
	UserRegisterRequest struct {
		Name          string `json:"name"`
		Password      string `json:"password"`
		Email         string `json:"email"`
		TypeAccountID int64  `json:"type_account_id"`
	}
	UserRegisterResponse struct {
	}
)

type (
	UserLoginRequest struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	UserLoginResponse struct {
		AccessToken      string `json:"access_token"`
		AccessExpiredAt  string `json:"access_expired_at"`
		RefreshToken     string `json:"refresh_token"`
		RefreshExpiredAt string `json:"refresh_expired_at"`
		User             User   `json:"user"`
	}
)

type (
	RefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token"`
	}

	AccessTokenResponse struct {
		AccessToken string `json:"access_token"`
	}
)
