package types

type Faculty struct {
	Name string `json:"name"`
}

type (
	GetFacultysRequest struct {
	}
	GetFacultysResponse struct {
		Result    Result    `json:"result"`
		Faculties []Faculty `json:"faculties"`
	}
)

type (
	GetFacultyByIDRequest struct {
		FacultyID int32 `uri:"facultyID"`
	}
	GetFacultyByIDResponse struct {
		Result  Result  `json:"result"`
		Faculty Faculty `json:"Faculty"`
	}
)
