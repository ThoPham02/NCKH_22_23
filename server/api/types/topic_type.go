package types

type Topic struct {
	ID           int32   `json:"id"`
	Name         string  `json:"name"`
	Lecture      string  `json:"lecture"` // lecture name
	FacultyID    int32   `json:"facultyId"`
	Status       string  `json:"status"`
	ResultUrl    string  `json:"resultUrl"`
	ListStudents []int32 `json:"listStudents"`
	GroupId      int32   `json:"groupId"`
	TimeStart    string  `json:"timeStart"`
	TimeEnd      string  `json:"timeEnd"`
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
		LectureID    int32  `form:"lectureId"`
		ConferenceID int32  `form:"conferenceId"`
		GroupID      int32  `form:"groupId"`
		Status       int32  `form:"status"`
		TimeStart    string `form:"timeStart"`
		TimeEnd      string `form:"timeEnd"`
		Limit        int32  `form:"limit"`
		Offset       int32  `form:"offset"`
	}

	GetTopicResponse struct {
		Result     Result  `json:"result"`
		ListTopics []Topic `json:"listTopics"`
	}
)

type (
	UpdateTopicRequest struct {
	}
	UpdateTopicResponse struct {
	}
)

type (
	CreateTopicRequest struct {
	}
	CreateTopicResponse struct {
	}
)

type (
	AcceptTopicRequest struct {
	}
	AcceptTopicResponse struct {
	}
)
