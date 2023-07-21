package utils

import (
	"time"
)

func StringToDate(text string) (time.Time, error) {
	layout := "2006-01-02"
	date, err := time.Parse(layout, text)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
