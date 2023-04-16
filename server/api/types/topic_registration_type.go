package types

type TopicRegistration struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	Lecture string `json:"lecture"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Faculty string `json:"faculty"`
	Status  string `json:"status"`
}

type (
	GetTopicRegistrationsRequest struct {
		Search    string `form:"search"`
		Lecture   string `form:"lecture"`
		FacultyID int32  `form:"facultyId"`
		Status    int32  `form:"status"`
	}
	GetTopicRegistrationsResponse struct {
		Result             Result              `json:"result"`
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
		Status int32 `json:"status"`
	}
	UpdateTopicRegistrationResponse struct {
		Result Result `json:"result"`
	}
)
