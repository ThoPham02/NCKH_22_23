package types

type Department struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	FaculityID int64  `json:"faculity_id"`
}

type (
	GetDepartmentByIDRequest struct {
		ID int64 `uri:"id"`
	}

	GetDepartmentByIDResponse struct {
		ID         int64  `json:"id"`
		Name       string `json:"name"`
		FaculityID int64  `json:"faculity_id"`
	}
)

type (
	GetListDepartmentByFaculityRequest struct {
		FaculityID int64 `form:"faculity_id"`
	}

	GetListDepartmentByFaculityResponse struct {
		ListDepartment []*Department `json:"list_department"`
	}
)
