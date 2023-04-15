package constant

const (
	SUCCESS_CODE    = 0
	SUCCESS_MESSAGE = "Success"

	DB_ERR_CODE    = 1
	DB_ERR_MESSAGE = "Database error"

	TOKEN_ERR_CODE         = 2
	TOKEN_ERR_CODE_MESSAGE = "Token error"

	TOKEN_EXPIRES_CODE    = 3
	TOKEN_EXPIRES_MESSAGE = "Token expires"

	HASH_PASSWORD_ERR_CODE    = 4
	HASH_PASSWORD_ERR_MESSAGE = "Hash password error"
)

const (
	DEPARTMENT_NOT_FOUND_CODE    = 404
	DEPARTMENT_NOT_FOUND_MESSAGE = "Department not found"
)

const (
	USER_NOT_FOUND_CODE    = 10000
	USER_NOT_FOUND_MESSAGE = "User not found"

	WRONG_PASSWORD_CODE    = 10001
	WRONG_PASSWORD_MESSAGE = "Wrong password"
)
