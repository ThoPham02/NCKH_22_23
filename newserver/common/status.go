package common

// topic status
const (
	TOPIC_WAIT_CONFIRM  = 1
	TOPIC_CONFIRMED     = 2
	TOPIC_REFUSED       = 3
	TOPIC_DOING         = 4
	TOPIC_DONE          = 5
	TOPIC_OUT_DATE      = 6
	TOPIC_DONE_OUT_DATE = 7
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
