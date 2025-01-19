package repetitionrule

import (
	"fmt"
	"strings"
	"time"
)

func RepetitionRule(now time.Time, date string, repeat string) (string, error) {
	if repeat == "" {
		return "", fmt.Errorf("empty repeat field")
	}

	firstDate, err := time.Parse("20060102", date)
	if err != nil {
		return "", fmt.Errorf("wrong date format")
	}

	parseRepeat := strings.Split(repeat, " ")

	switch parseRepeat[0] {
	case "d":
		return repeatDay(now, firstDate, parseRepeat)
	case "y":
		return repeatYear(now, firstDate, parseRepeat)
	case "w":
		return repeatWeekday(now, firstDate, parseRepeat)
	case "m":
		return repeatMonth(now, firstDate, parseRepeat)
	}

	return "", fmt.Errorf("invalid repeat format")
}
