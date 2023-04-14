package utils

import (
	"fmt"
	"time"
)

func ConvertStringToTime(s string) (*time.Time, error) {
	layout := "2006-01-02 15:04:05" // the format of your string

	t, err := time.Parse(layout, s)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return &t, nil
}
