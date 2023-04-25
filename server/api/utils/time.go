package utils

import (
	"fmt"
	"github/ThoPham02/research_management/api/constant"
	"time"
)

func ConvertStringToTime(s string) (*time.Time, error) {
	t, err := time.Parse(constant.Layout, s)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return &t, nil
}

func CompareTimeStringWithNow(timeString string) (bool, error) {
	// Parse the time string into a time.Time object
	t, err := time.Parse(constant.Layout, timeString)
	if err != nil {
		return false, err
	}

	// Get the current time
	now := time.Now()

	// Compare the two times
	return t.After(now), nil
}
