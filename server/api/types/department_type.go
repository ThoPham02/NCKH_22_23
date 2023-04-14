package types

type Department struct {
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
		DepartmentID int32 `uri:"DepartmentID"`
	}
	GetDepartmentByIDResponse struct {
		Result     Result     `json:"result"`
		Department Department `json:"department"`
	}
)
