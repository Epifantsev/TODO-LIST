package repetitionrule

import (
	"fmt"
	"slices"
	"strconv"
	"time"
)

func convertingInInt(sl []string) ([]int, error) {
	var slInt []int

	for _, val := range sl {
		n, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("converting error")
		}
		slInt = append(slInt, n)
	}

	return slInt, nil
}

func nearestDate(sl []int, now int) int {
	slices.Sort(sl)
	nextDate := sl[0]

	for _, v := range sl {
		if now < v {
			nextDate = v
			break
		}
	}

	return nextDate
}

func getPositiveNumber(sl []int, total int) []int {
	for i := range sl {
		switch sl[i] {
		case -1:
			sl[i] = total
		case -2:
			sl[i] = total - 1
		}
	}
	return sl
}

func daysInMonth(now time.Time) int {
	nextMonth := 1
	year := now.Year()

	if nextMonth > 12 {
		nextMonth = 1
		year++
	}

	allDaysOfMonth := time.Date(year, now.Month()+time.Month(nextMonth), 0, 0, 0, 0, 0, time.UTC)

	return allDaysOfMonth.Day()
}
