package utils

import (
	"database/sql"
	"time"
)

func GetString(input *string) sql.NullString {
	if input == nil {
		return sql.NullString{
			Valid: false,
		}
	}
	return sql.NullString{
		Valid:  input != nil,
		String: *input,
	}
}

func GetInt32(input *int32) sql.NullInt32 {
	if input == nil {
		return sql.NullInt32{
			Valid: false,
		}
	}
	return sql.NullInt32{
		Valid: input != nil,
		Int32: *input,
	}
}

func GetTime(input *time.Time) sql.NullTime {
	if input == nil {
		return sql.NullTime{
			Valid: false,
		}
	}
	return sql.NullTime{
		Valid: input != nil,
		Time:  *input,
	}
}
