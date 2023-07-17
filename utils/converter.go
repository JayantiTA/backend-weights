package utils

import (
	"strconv"
	"time"
)

func StringToInt64(text string) (int64, error) {
	num, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func StringToDate(text string) (time.Time, error) {
	layout := "2006-01-02"
	date, err := time.Parse(layout, text)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func UnixToTime(unix string) (time.Time, error) {
	unixInt, err := StringToInt64(unix)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(unixInt, 0), nil
}
