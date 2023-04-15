package utils

import "regexp"

func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	matched, err := regexp.MatchString(pattern, email)

	if err != nil || !matched {
		return false
	}

	return true
}
