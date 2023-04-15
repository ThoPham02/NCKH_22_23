package types

type UserInfo struct {
	ID          int32  `json:"id"`
	UserID      int32  `json:"userId"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Faculty     string `json:"faculty"`
	Degree      string `json:"degree"`
	YearStart   int32  `json:"yearStart"`
	AvatarUrl   string `json:"avatarUrl"`
	Birthday    string `json:"birthday"`
	BankAccount string `json:"bankAccount"`
}

type (
	GetUserInfoRequest struct {
	}
	GetUserInfoResponse struct {
		Result   Result   `json:"result"`
		UserInfo UserInfo `json:"userInfo"`
	}
)

type (
	UpdateUserInfoRequest struct {
		UserID      int32   `json:"userID"`
		Name        string  `json:"name"`
		Email       string  `json:"email"`
		Phone       string  `json:"phone"`
		FacultyID   int32   `json:"facultyID"`
		Degree      int32   `json:"degree"`
		YearStart   int32   `json:"yearStart"`
		AvatarUrl   *string `json:"avatarUrl"`
		Birthday    *string `json:"birthday"`
		BankAccount *string `json:"bankAccount"`
	}
	UpdateUserInfoResponse struct {
		Result Result `json:"result"`
	}
)
