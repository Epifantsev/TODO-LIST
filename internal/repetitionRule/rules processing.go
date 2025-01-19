package repetitionrule

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func repeatMonth(now time.Time, firstDate time.Time, parseRepeat []string) (string, error) {
	if len(parseRepeat) > 3 || len(parseRepeat) < 2 {
		return "", fmt.Errorf("wrong length parseRepeat")
	}

	parseRuleInDays := strings.Split(parseRepeat[1], ",")

	days, err := convertingInInt(parseRuleInDays)
	if err != nil {
		return "", err
	}

	for _, v := range days {
		if v < -2 || v > 31 || v == 0 {
			return "", fmt.Errorf("wrong day")
		}
	}

	nextDateRepeat := firstDate
	if firstDate.Before(now) {
		nextDateRepeat = now
	}

	var nextMonth int

	if len(parseRepeat) == 3 {
		parseRuleInMonths := strings.Split(parseRepeat[2], ",")
		months, err := convertingInInt(parseRuleInMonths)
		if err != nil {
			return "", err
		}

		for _, v := range months {
			if v < 1 || v > 12 {
				return "", fmt.Errorf("wrong month")
			}
		}

		nextMonth = nearestDate(months, int(nextDateRepeat.Month()))

		for {
			if int(nextDateRepeat.Month()) == nextMonth {
				break
			}
			nextDateRepeat = nextDateRepeat.AddDate(0, 1, 0)
			nextDateRepeat = time.Date(nextDateRepeat.Year(), nextDateRepeat.Month(), 1, 0, 0, 0, 0, time.UTC)
		}
	}

	dayNow := nextDateRepeat.Day()
	totalDays := daysInMonth(nextDateRepeat)
	days = getPositiveNumber(days, totalDays)
	nextDay := nearestDate(days, dayNow)
	dateForComparison := nextDateRepeat

	for {
		if nextDateRepeat.Day() == nextDay {
			break
		}

		nextDateRepeat = nextDateRepeat.AddDate(0, 0, 1)

		if dateForComparison.Month() < nextDateRepeat.Month() {
			totalDays = daysInMonth(nextDateRepeat)
			days = getPositiveNumber(days, totalDays)
			nextDay = nearestDate(days, dayNow)
			dateForComparison = nextDateRepeat
		}
	}

	if nextDateRepeat.Before(now) {
		return "", fmt.Errorf("next date is before current date")
	}

	return nextDateRepeat.Format("20060102"), nil

}

func repeatWeekday(now time.Time, firstDate time.Time, parseRepeat []string) (string, error) {
	if len(parseRepeat) < 2 {
		return "", fmt.Errorf("parseRepeat length is least 2 elements")
	}

	parseRepeat = strings.Split(parseRepeat[1], ",")

	if len(parseRepeat) > 7 {
		return "", fmt.Errorf("length parseRepeat exceeded")
	}

	weekDays, err := convertingInInt(parseRepeat)
	if err != nil {
		return "", err
	}

	for _, val := range weekDays {
		if val > 7 || val < 1 {
			return "", fmt.Errorf("wrong day of the week")
		}
	}

	nextDateRepeat := firstDate
	if firstDate.Before(now) {
		nextDateRepeat = now
	}

	nowWeekDay := int(nextDateRepeat.Weekday())

	nextWeekDay := nearestDate(weekDays, nowWeekDay)
	if nextWeekDay == 7 {
		nextWeekDay = 0
	}

	if nowWeekDay == nextWeekDay {
		nextDateRepeat = nextDateRepeat.AddDate(0, 0, 7)
	}

	for {
		nextDateRepeat = nextDateRepeat.AddDate(0, 0, 1)
		if int(nextDateRepeat.Weekday()) == nextWeekDay {
			break
		}
	}

	if nextDateRepeat.Before(now) {
		return "", fmt.Errorf("next date is before current date")
	}

	return nextDateRepeat.Format("20060102"), nil
}

func repeatDay(now time.Time, firstDate time.Time, parseRepeat []string) (string, error) {
	if len(parseRepeat) != 2 {
		return "", fmt.Errorf("error repeat d")
	}

	d, err := strconv.Atoi(parseRepeat[1])
	if err != nil || d < 1 || d > 400 {
		return "", fmt.Errorf("invalid day interval")
	}

	var nextDateRepeat time.Time

	daysSinceFirst := int(now.Sub(firstDate).Hours() / 24)

	switch {
	case daysSinceFirst <= 1:
		nextDateRepeat = firstDate.AddDate(0, 0, d)
	default:
		surplus := daysSinceFirst % d
		nextDateRepeat = firstDate.AddDate(0, 0, daysSinceFirst-surplus+d)
	}

	if nextDateRepeat.Before(now) {
		return "", fmt.Errorf("next date is before current date")
	}

	return nextDateRepeat.Format("20060102"), nil
}

func repeatYear(now time.Time, firstDate time.Time, parseRepeat []string) (string, error) {
	if len(parseRepeat) != 1 {
		return "", fmt.Errorf("error repeat y")
	}

	var nextDateRepeat time.Time

	between := now.Year() - firstDate.Year()

	if between <= 0 {
		nextDateRepeat = firstDate.AddDate(1, 0, 0)
	} else {
		nextDateRepeat = firstDate.AddDate(between, 0, 0)
	}

	if nextDateRepeat.Before(now) {
		nextDateRepeat = nextDateRepeat.AddDate(1, 0, 0)
	}

	return nextDateRepeat.Format("20060102"), nil
}
