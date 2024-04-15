package utils

import "time"

const layoutData = "2006-01-02T15:04:05"

func StringToDate(value string) time.Time {

	convertedTime, _ := time.Parse(layoutData, value)
	return convertedTime

}
