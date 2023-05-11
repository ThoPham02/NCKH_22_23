package common

const (
	SUCCESS_CODE = 0
	SUCCESS_MESS = "Success"

	DB_ERR_CODE = 1
	DB_ERR_MESS = "DB Error"

	INVALID_INPUT_ERR_CODE = 2
	INVALID_INPUT_ERR_MESS = "Invalid Input Error"
)

// User error code 100-199
const (
	LOGIN_ERR_CODE = 100
	LOGIn_ERR_MESS = "Login Error"
)

// JWT key
const (
	JWT_IAT_KEY        = "iat"
	JWT_EXP_KEY        = "exp"
	JWT_USER_ID_KEY    = "user_id"
	JWT_SESSION_ID_KEY = "session_id"
)