package types

type TopicRegistration struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	Lecture string `json:"lecture"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type (
	GetTopicRegistrationsRequest struct {
		Search    string `form:"search"`
		Lecture   string `form:"lecture"`
		FacultyID int32  `form:"facultyId"`
		Status    int32  `form:"status"`
		Limit     int32  `form:"limit"`
		Offset    int32  `form:"offset"`
	}
	GetTopicRegistrationsResponse struct {
		Result             Result              `json:"result"`
		Total              int                 `json:"total"`
		TopicRegistrations []TopicRegistration `json:"topicRegistrations"`
	}
)

type (
	GetTopicRegistrationByIdRequest struct {
	}
	GetTopicRegistrationByIdResponse struct {
		Result            Result            `json:"result"`
		TopicRegistration TopicRegistration `json:"topicRegistration"`
	}
)

type (
	CreateTopicRegistrationRequest struct {
		Name string `json:"name"`
	}
	CreateTopicRegistrationResponse struct {
		Result Result `json:"result"`
	}
)

type (
	UpdateTopicRegistrationRequest struct {
		Name      string `json:"name"`
		LectureID int32  `json:"lectureId"`
		FacultyID int32  `json:"facultyId"`
		Status    int32  `json:"status"`
	}
	UpdateTopicRegistrationResponse struct {
		Result Result `json:"result"`
	}
)

type (
	AccceptTopicRegistrationRequest struct {
		ListTopicRegistrationId string `json:"listTopicRegistrationId"`
		Status                  int32  `json:"status"`
	}
	AcceptTopicRegistrationResponse struct {
		Result Result `json:"result"`
	}
)
