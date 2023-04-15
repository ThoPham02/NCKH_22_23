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
	}
	GetFacultyByIDResponse struct {
		Result  Result  `json:"result"`
		Faculty Faculty `json:"faculty"`
	}
)
