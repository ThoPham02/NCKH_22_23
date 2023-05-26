package common

// topic status
const (
	TOPIC_CANCEL              = 1
	TOPIC_SUGGESTION          = 2
	TOPIC_REGISTATION         = 4
	TOPIC_WAIT_CONFIRM        = 8
	TOPIC_DOING               = 16
	TOPIC_REPORT_STAGE        = 32
	TOPIC_REPORT_SUBCOMMITTEE = 64
	TOPIC_REPORT_SCHOOL       = 128
	TOPIC_DONE                = 256
	TOPIC_DONE_OUT_DATE       = 512
	TOPIC_NOT_DONE            = 1024
)

// user role
const (
	ROLE_STUDENT    = 1
	ROLE_LECTURE    = 2
	ROLE_DEPARTMENT = 3
	ROLE_FACULTY    = 4
	ROLE_ADMIN      = 5
)

// user degree
const (
	DEGREE_STUDENT             = 1
	DEGREE_MASTER              = 2
	DEGREE_DOCTOR              = 3
	DEGREE_ASSOCIATE_PROFESSOR = 4
	DEGREE_PROFESSOR           = 5
)

// LEVEL
const (
	LEVEL_SUBCOMMITTEE = 1
	LEVEL_SCHOOL       = 2
)
