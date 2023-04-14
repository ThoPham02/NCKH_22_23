package types

type (
	GetUserInfoRequest struct {
	}
	GetUserInfoResponse struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		AvatarUrl   string `json:"avatar_url"`
		Birthday    string `json:"birthday"`
		FaculityID  int64  `json:"faculity_id"`
		YearStart   int64  `json:"year_start"`
		BankAccount string `json:"bank_account"`
		Phone       string `json:"phone"`
		Sex         int64  `json:"sex"`
	}
)

type (
	UpdateUserInfoRequest struct {
		Name        string  `json:"name"`
		Description *string `json:"description"`
		AvatarUrl   *string `json:"avatar_url"`
		Birthday    *string `json:"birthday"`
		FaculityID  int64   `json:"faculity_id"`
		YearStart   int64   `json:"year_start"`
		BankAccount *string `json:"bank_account"`
		Phone       *string `json:"phone"`
		Sex         *int64  `json:"sex"`
	}
	UpdateUserInfoResponse struct {
	}
)
