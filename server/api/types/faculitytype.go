package types

type Faculity struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type (
	GetFaculityByIDRequest struct {
		ID int64 `uri:"id"`
	}

	GetFaculityByIDResponse struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
)

type (
	GetListFaculityRequest struct {
	}

	GetListFaculityResponse struct {
		ListFaculity []*Faculity `json:"list_faculity"`
	}
)
