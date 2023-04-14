package types

type TopicRegistration struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	Lecture string `json:"lecture"`
	Faculty string `json:"faculty"`
	Status  string `json:"status"`
}

type (
	GetTopicRegistrationRequest struct {
		Search    string `form:"search"`
		LectureID int32  `form:"lectureId"`
		FacultyID int32  `form:"facultyId"`
		Status    int32  `form:"status"`
	}
	GetTopicRegistrationResponse struct {
		Result             Result              `json:"result"`
		TopicRegistrations []TopicRegistration `json:"topicRegistrations"`
	}
)

type (
	CreateTopicRegistrationRequest struct {
		Name      string `json:"name"`
		LectureID int32  `json:"lectureId"`
		FacultyID int32  `json:"facultyId"`
	}
	CreateTopicRegistrationResponse struct {
		Result Result `json:"result"`
	}
)

type (
	UpdateTopicRegistrationRequest struct {
		Status int32 `json:"status"`
	}
	UpdateTopicRegistrationResponse struct {
		Result Result `json:"result"`
	}
)
