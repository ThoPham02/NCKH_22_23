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

func CompareTimeStringWithNow(timeString string) (bool, error) {
	// Parse the time string into a time.Time object
	t, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		return false, err
	}

	// Get the current time
	now := time.Now()

	// Compare the two times
	return t.Before(now), nil
}
