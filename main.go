package main

import (
	"fmt"
	"time"
)

var weekStartDay time.Weekday = 1
var startDate time.Time = time.Date(time.Now().Year(), 2, 1, 0, 0, 0, 0, time.Local)
var sampleSize int = 100

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
func repetativePatternSearch(source string) string {
	pattern := ""
	for i := 0; i < len(source); {
		if len(pattern) == 0 || source[i] != pattern[i%len(pattern)] {
			pattern = source[0 : len(pattern)+1]
			i = 0
		} else {
			i++
		}
	}
	return pattern
}

func main() {
	previousStartDate := startDate
	secuencial := ""
	for i := 0; i < sampleSize; {
		startDate = startDate.AddDate(1, 0, 0)
		if startDate.Weekday() == weekStartDay && !isLeapYear(startDate.Year()) {
			yearsDiff := startDate.Year() - previousStartDate.Year()
			secuencial += fmt.Sprint(yearsDiff) + ","
			fmt.Println(yearsDiff)
			previousStartDate = startDate
			i++
		}
	}
	fmt.Println(repetativePatternSearch(secuencial))
}
