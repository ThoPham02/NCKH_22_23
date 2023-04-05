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
