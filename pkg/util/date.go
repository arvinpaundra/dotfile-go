package util

import (
	"fmt"
	"time"
)

var LocationTime, _ = time.LoadLocation("Asia/Jakarta")

func FormatStartDate(date string) string {
	return fmt.Sprintf("%s 00:00:00", date)
}

func FormatEndDate(date string) string {
	return fmt.Sprintf("%s 23:59:59", date)
}

// GetStartDayOfWeek return time.Time monday on current week
func GetStartDayOfWeek(t time.Time) time.Time {
	daySinceMonday := int(t.Weekday()) - 1
	return t.AddDate(0, 0, -daySinceMonday)
}

// GetEndDayOfWeek return time.Time sunday on current week
func GetEndDayOfWeek(t time.Time) time.Time {
	dayUntilSunday := 7 - int(t.Weekday())
	return t.AddDate(0, 0, dayUntilSunday)
}

// GetCurrentWeekRange returns time.Time start and end current week range
func GetCurrentWeekRange() (time.Time, time.Time) {
	now := time.Now().In(LocationTime)

	currentWeekStartDay := GetStartDayOfWeek(now)
	currentWeekEndDay := GetEndDayOfWeek(now)

	return currentWeekStartDay, currentWeekEndDay
}

// GetPreviousWeekRange returns time.Time start and end previous week range
func GetPreviousWeekRange() (time.Time, time.Time) {
	currentStartDay, _ := GetCurrentWeekRange()

	previousWeekStartDay := currentStartDay.AddDate(0, 0, -7)
	previousWeekEndDay := currentStartDay.AddDate(0, 0, -1)

	return previousWeekStartDay, previousWeekEndDay
}

// GetPastDays return time.Time from current to past `n` days ago
func GetPastDays(n int) time.Time {
	return time.Now().AddDate(0, 0, n*-1)
}

// GetCurrentMonthRange returns time.Time start and end current month range
func GetCurrentMonthRange() (time.Time, time.Time) {
	now := time.Now().In(LocationTime)

	currentMonthStartDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, LocationTime)
	currentMonthEndDay := currentMonthStartDay.AddDate(0, 1, 0)

	return currentMonthStartDay, currentMonthEndDay
}

// GetPreviousMonthRange returns time.Time start and end previous month range
func GetPreviousMonthRange() (time.Time, time.Time) {
	currentMonthStartDay, _ := GetCurrentMonthRange()

	previousMonthStartDay := currentMonthStartDay.AddDate(0, -1, 0)
	previousMonthEndDay := currentMonthStartDay.AddDate(0, 0, -1)

	return previousMonthStartDay, previousMonthEndDay
}

func GetMonthRangeByDate(date, layout string) (time.Time, time.Time, error) {
	parsedTime, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	firstDay := time.Date(parsedTime.Year(), parsedTime.Month(), 1, 0, 0, 0, 0, LocationTime)
	lastDay := firstDay.AddDate(0, 1, -1)

	return firstDay, lastDay, nil
}
