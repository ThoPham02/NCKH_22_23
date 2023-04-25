package constant

const (
	WrongPasswordErrMsg   = "wrong password"
	UserIsNotExistErrMsg  = "user is not exist"
	InputValidationErrMsg = "input validation error"
)

const (
	AuthorizationHeaderKey  = "Authorization"
	AuthorizationTypeBearer = "Bearer"
	PayloadKey              = "payload_key"
	AdminType               = 5
	StudentType             = 1
	LectureType             = 2
	DepartmentType          = 3
	FaculityType            = 4
)

type strings string

const (
	TraceIDKey strings = "trace-id"
)

var MapDegree = map[int32]string{
	1: "Sinh viên",
	2: "Cử nhân",
	3: "Thạc sĩ",
	4: "Tiến sĩ",
	5: "Phó giáo sư",
	6: "Giáo sư",
}

var MapTypeAccount = map[int32]string{
	1: "Sinh Viên",
	2: "Giảng Viên",
	3: "Bộ Môn",
	4: "Khoa",
	5: "Admin",
}

var MapStatusTopicRegistration = map[int32]string{
	1: "Chờ phê duyệt",
	2: "Đã duyệt",
	3: "Đã hủy",
}

var MapStatusTopic = map[int32]string{
	1: "Chờ phê duyệt",
	2: "Đang thực hiện",
	3: "Sắp hết hạn",
	4: "Hoàn thành",
	5: "Hoàn thành quá hạn",
	6: "Quá hạn",
}

const Layout = "2006-01-02 15:04:05"
