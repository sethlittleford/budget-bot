package transactions

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// Months returns all 12 months in various string formats
// that are valid arguments in the budget-bot run command
func Months() []string {
	m := []time.Month{
		time.January,
		time.February,
		time.March,
		time.April,
		time.May,
		time.June,
		time.July,
		time.August,
		time.September,
		time.October,
		time.November,
		time.December,
	}
	months := make([]string, 0)
	for _, month := range m {
		months = append(months, month.String())                      // e.g. January
		months = append(months, strings.ToLower(month.String()))     // e.g. january
		months = append(months, month.String()[:3])                  // e.g. Jan
		months = append(months, strings.ToLower(month.String()[:3])) // e.g. jan
		months = append(months, strconv.Itoa(int(month)))            // e.g. 1
	}
	return months
}

// StringToMonth converts the string representation of a month
// into a time.Month, if possible
func StringToMonth(s string) (time.Month, error) {
	m := Months()
	switch s {
	case m[0], m[1], m[2], m[3], m[4]:
		return time.January, nil
	case m[5], m[6], m[7], m[8], m[9]:
		return time.February, nil
	case m[10], m[11], m[12], m[13], m[14]:
		return time.March, nil
	case m[15], m[16], m[17], m[18], m[19]:
		return time.April, nil
	case m[20], m[21], m[22], m[23], m[24]:
		return time.May, nil
	case m[25], m[26], m[27], m[28], m[29]:
		return time.June, nil
	case m[30], m[31], m[32], m[33], m[34]:
		return time.July, nil
	case m[35], m[36], m[37], m[38], m[39]:
		return time.August, nil
	case m[40], m[41], m[42], m[43], m[44]:
		return time.September, nil
	case m[45], m[46], m[47], m[48], m[49]:
		return time.October, nil
	case m[50], m[51], m[52], m[53], m[54]:
		return time.November, nil
	case m[55], m[56], m[57], m[58], m[59]:
		return time.December, nil
	default:
		return 0, errors.New("Could not recognize month string")
	}
}

// Years returns all years in YYYY format
// that are valid arguments in the budget-bot run command
func Years() []string {
	base := "20"
	years := make([]string, 0)
	for i := 22; i < 100; i++ {
		years = append(years, base+strconv.Itoa(i))
	}
	return years
}

// GetStartEndDates takes in a month and returns first and last days
// of that month in Plaid string format YYYY-MM-DD
func GetStartEndDates(m time.Month) (startDate, endDate string) {
	firstOfMonth := time.Date(2022, m, 1, 0, 0, 0, 0, time.UTC)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	startDate = firstOfMonth.Format("2020-01-01")
	endDate = lastOfMonth.Format("2020-01-01")
	return
}
