package types

type Department struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	FacultyID int32  `json:"facultyId"`
}

type (
	GetDepartmentsRequest struct {
		FacultyID int32 `form:"facultyId"`
	}
	GetDepartmentsResponse struct {
		Result      Result       `json:"result"`
		Departments []Department `json:"departments"`
	}
)

type (
	GetDepartmentByIDRequest struct {
	}
	GetDepartmentByIDResponse struct {
		Result     Result     `json:"result"`
		Department Department `json:"department"`
	}
)

type (
	CreateDepartmentRequest struct {
		Name      string `json:"name"`
		FacultyID int32  `json:"facultyId"`
	}
	CreateDepartmentResponse struct {
		Result Result `json:"result"`
	}
)
