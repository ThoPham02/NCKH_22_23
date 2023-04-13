package types

type TopicRegistation struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	DescriptionURL string `json:"description_url"`
	LectureID      int64  `json:"lecture_id"`
	FaculityID     int64  `json:"faculity_id"`
	CreatedAt      string `json:"created_at"`
}

type (
	GetListTopicRegistationRequest struct {
		Search     string `form:"search"`
		LectureID  int64  `form:"lecture_id"`
		FaculityID int64  `form:"faculity_id"`
		Limit      int64  `form:"limit"`
		Offset     int64  `form:"offset"`
	}

	GetListTopicRegistationResponse struct {
		Total                int64              `json:"total"`
		ListTopicRegistation []TopicRegistation `json:"list_topic_registation"`
	}
)

type (
	GetListTopicRegistationByLectureIDRequest struct {
		LectureID int64 `json:"lecture_id"`
	}
	GetListTopicRegistationByLectureIDResponse struct {
		ListTopicRegistation []TopicRegistation `json:"list_topic_registation"`
	}
)

type (
	UpdateTopicRegistationRequest struct {
		ID             int64  `json:"id"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		DescriptionUrl string `json:"description_url"`
		LectureID      int64  `json:"lecture_id"`
		FaculityID     int64  `json:"faculity_id"`
	}
	UpdateTopicRegistationResponse struct {
		ID             int64  `json:"id"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		DescriptionUrl string `json:"description_url"`
		LectureID      int64  `json:"lecture_id"`
		FaculityID     int64  `json:"faculity_id"`
	}
)

type (
	CreateTopicRegistationRequest struct {
		Name           string `json:"name"`
		Description    string `json:"description"`
		DescriptionUrl string `json:"description_url"`
		LectureID      int64  `json:"lecture_id"`
		FaculityID     int64  `json:"faculity_id"`
	}
	CreateTopicRegistationResponse struct {
		ID             int64  `json:"id"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		DescriptionUrl string `json:"description_url"`
		LectureID      int64  `json:"lecture_id"`
		FaculityID     int64  `json:"faculity_id"`
	}
)

type (
	DeleteTopicRegistationRequest struct {
		ID int64 `uri:"id"`
	}
)

type (
	GetTopicRegistationByIDRequest struct {
		ID int64 `uri:"id"`
	}
	GetTopicRegistationByIdResponse struct {
		ID             int64  `json:"id"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		DescriptionUrl string `json:"description_url"`
		LectureID      int64  `json:"lecture_id"`
		FaculityID     int64  `json:"faculity_id"`
	}
)
