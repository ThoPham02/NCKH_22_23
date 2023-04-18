package types

type Conference struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	CashSupport int32  `json:"cashSupport"`
	SchoolYear  string `json:"schoolYear"`
}

type (
	UpdateConferenceRequest struct {
		ID          int32 `json:"id"`
		CashSupport int32 `json:"cashSupport"`
	}
	UpdateConferenceResponse struct {
		Result Result `json:"result"`
	}
)

type (
	ListConferenceRequest struct {
		Name string `form:"name"`
	}
	ListConferenceResponse struct {
		Result      Result       `json:"result"`
		Conferences []Conference `json:"conference"`
	}
)

type (
	CreateConferenceRequest struct {
		Name         string `json:"name"`
		CrashSupport *int32 `json:"crashSupport"`
		SchoolYear   string `json:"schoolYear"`
	}
	CreateConferenceResponse struct {
		Result Result `json:"result"`
	}
)

type (
	GetConferenceByIdRequest struct {
	}
	GetConferenceByIdResponse struct {
		Result     Result     `json:"result"`
		Conference Conference `json:"conference"`
	}
)
