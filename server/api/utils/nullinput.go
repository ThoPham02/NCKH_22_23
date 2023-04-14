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

func GetInt64(input *int64) sql.NullInt64 {
	if input == nil {
		return sql.NullInt64{
			Valid: false,
		}
	}
	return sql.NullInt64{
		Valid: input != nil,
		Int64: *input,
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