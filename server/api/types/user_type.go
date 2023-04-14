package types

type User struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	TypeAccount int32  `json:"typeAccount"`
}

type Token struct {
	AccessToken      string `json:"accessToken"`
	AccessExpiresAt  string `json:"accessExpiresAt"`
	RefreshToken     string `json:"refreshToken"`
	RefreshExpiresAt string `json:"refreshExpiresAt"`
}

type (
	UserLoginRequest struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	UserLoginResponse struct {
		Result Result `json:"result"`
		Token  Token  `json:"token"`
		User   User   `json:"user"`
	}
)

type (
	UserRegisterRequest struct {
		Name        string `json:"name"`
		Password    string `json:"password"`
		Email       string `json:"email"`
		TypeAccount int32  `json:"typeAccount"`
	}
	UserRegisterResponse struct {
		Result Result `json:"result"`
		User   User   `json:"user"`
		Token  Token  `json:"token"`
	}
)
