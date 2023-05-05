package types

type Topic struct {
	ID           int32    `json:"id"`
	Name         string   `json:"name"`
	Lecture      string   `json:"lecture"` // lecture name
	Faculty      string   `json:"faculty"` // faculty name
	Status       string   `json:"status"`
	ResultUrl    string   `json:"resultUrl"`
	ListStudents []string `json:"listStudents"` // list of students
	TimeStart    string   `json:"timeStart"`
	TimeEnd      string   `json:"timeEnd"`
}

type (
	GetTopicByIdRequest struct {
	}

	GetTopicByIdResponse struct {
		Result Result `json:"result"`
		Topic  Topic  `json:"topic"`
	}
)

type (
	GetTopicRequest struct {
		Search       string `form:"search"`
		ConferenceID int32  `form:"conferenceId"`
		LectureID    int32  `form:"lectureId"`
		FacultyID    int32  `form:"facultyId"`
		StudentID    int32  `form:"studentId"`
		Status       int32  `form:"status"`
		TimeStart    string `form:"timeStart"`
		TimeEnd      string `form:"timeEnd"`
		Limit        int32  `form:"limit"`
		Offset       int32  `form:"offset"`
	}

	GetTopicResponse struct {
		Result     Result  `json:"result"`
		Total      int     `json:"total"`
		ListTopics []Topic `json:"listTopics"`
	}
)

type (
	UpdateTopicRequest struct {
		ListTopicID string `json:"listTopicId"`
		GroupID     *int32 `json:"groupId"`
	}
	UpdateTopicResponse struct {
		Result Result `json:"result"`
	}
)

type (
	CreateTopicRequest struct {
		ListStudent  string `json:"listStudent"`
		TopicRegisID int32  `json:"topicRegisId"`
		ConferenceID int32  `json:"conferenceId"`
	}
	CreateTopicResponse struct {
		Result Result `json:"result"`
	}
)

type (
	AcceptTopicRequest struct {
		ListTopicID string `json:"listTopicId"`
		TimeStart   string `json:"timeStart"`
		TimeEnd     string `json:"timeEnd"`
	}
	AcceptTopicResponse struct {
		Result Result `json:"result"`
	}
)
