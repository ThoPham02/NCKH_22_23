package types

type UserRegisterRequest struct {
	Name          string `json:"name"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	TypeAccountID int64  `json:"type_account_id"`
}
