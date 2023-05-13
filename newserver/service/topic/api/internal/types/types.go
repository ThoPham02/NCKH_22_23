// Code generated by goctl. DO NOT EDIT.
package types

type Result struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type Department struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	FacultyID int64  `json:"facultyID"`
}

type Faculty struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Topic struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	LectureID       int64  `json:"lectureID"`
	DepartmentID    int64  `json:"departmentID"`
	Status          int64  `json:"status"`
	SubcommitteeID  int64  `json:"subcommitteeID"`
	GroupStudentsID int64  `json:"groupStudentsID"`
	TimeStart       int64  `json:"timeStart"`
	TimeEnd         int64  `json:"timeEnd"`
}

type StudentGroup struct {
	ID        int64 `json:"id"`
	StudentID int64 `json:"studentID"`
	EventID   int64 `json:"eventID"`
	GroupID   int64 `json:"groupID"`
}

type GetDepartmentsReq struct {
	FacultyID int64 `json:"facultyID"`
}

type GetDepartmentsRes struct {
	Result      Result       `json:"result"`
	Departments []Department `json:"departments"`
}

type CreateDepartmentReq struct {
	Name string `json:"name"`
}

type CreateDepartmentRes struct {
	Result Result `json:"result"`
}

type GetFacultiesReq struct {
}

type GetFacultiesRes struct {
	Result    Result    `json:"result"`
	Faculties []Faculty `json:"faculties"`
}

type CreateFacultyReq struct {
	Name      string `json:"name"`
	FacultyID int64  `json:"facultyID"`
}

type CreateFacultyRes struct {
	Result Result `json:"result"`
}

type CreateTopicReq struct {
	Name         string `json:"name"`
	LectureID    int64  `json:"lectureID"`
	DepartmentID int64  `json:"departmentID"`
	EventID      int64  `json:"eventID"`
}

type CreateTopicRes struct {
	Result Result `json:"result"`
}

type UpdateTopicReq struct {
	ID             int64  `path:"id"`
	Name           string `json:"name"`
	LectureID      int64  `json:"lectureID"`
	DepartmentID   int64  `json:"departmentID"`
	Status         int64  `json:"status"`
	EventID        int64  `json:"eventID"`
	SubcommitteeID int64  `json:"subcommitteeID"`
	GroupID        int64  `json:"groupID"`
	TimeStart      int64  `json:"timeStart"`
	TimeEnd        int64  `json:"timeEnd"`
	CashSupport    int64  `json:"cashSupport"`
}

type UpdateTopicRes struct {
	Result Result `json:"result"`
}

type UpdateTopicStatusReq struct {
	ID     int64 `path:"id"`
	Status int64 `json:"status"`
}

type UpdateTopicStatusRes struct {
	Result Result `json:"result"`
}

type UpdateTopicSubcommitteeReq struct {
	ListTopicID    []int64 `json:"listTopicID"`
	SubcommitteeID int64   `json:"subcommitteeID"`
}

type UpdateTopicSubcommitteeRes struct {
	Result Result `json:"result"`
}

type UpdateTopicStudentGroupReq struct {
	TopicID       int64   `path:"topicID"`
	ListStudentID []int64 `path:"listStudentID"`
}

type UpdateTopicStudentGroupRes struct {
	Result Result `json:"result"`
}

type GetTopicReq struct {
	ID int64 `path:"id"`
}

type GetTopicRes struct {
	Result Result `json:"result"`
	Topic  Topic  `json:"topic"`
}

type GetTopicsReq struct {
	Search         string `form:"search"`
	DepartmentID   int64  `form:"departmentID"`
	FacultyID      int64  `form:"facultyID"`
	Status         int64  `form:"status"`
	LectureID      int64  `form:"lectureID"`
	EventID        int64  `form:"eventID"`
	SubcommitteeID int64  `form:"subcommitteeID"`
	TimeStart      int64  `form:"timeStart"`
	TimeEnd        int64  `form:"timeEnd"`
	Limit          int64  `form:"limit"`
	Offset         int64  `form:"offset"`
}

type GetTopicsRes struct {
	Result Result  `json:"result"`
	Total  int64   `json:"total"`
	Topics []Topic `json:"topic"`
}
