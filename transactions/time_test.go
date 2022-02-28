package transactions

import (
	"testing"
	"time"
)

var months = Months()

func TestMonths(t *testing.T) {
	j := []string{"January", "january", "Jan", "jan", "1"}
	d := []string{"December", "december", "Dec", "dec", "12"}

	for i := 0; i < len(j); i++ {
		if j[i] != months[i] {
			t.Fatalf("%s is not in the valid Month arguments", j[i])
		}
		if d[i] != months[len(months)-(len(d)-i)] {
			t.Fatalf("%s is not in the valid Month arguments", d[i])
		}
	}
}

func TestStringToMonth(t *testing.T) {
	for _, month := range months {
		_, err := StringToMonth(month)
		if err != nil {
			t.Fatalf("StringToMonth() failed with err: %v", err)
		}
	}
}

func TestYears(t *testing.T) {
	needYears := []string{"2022", "2023", "2024", "2025", "2026", "2027", "2028", "2029", "2030"}
	haveYears := Years()

	for i, needYear := range needYears {
		if needYear != haveYears[i] {
			t.Fatalf("%s is not in the valid Year arguments", needYear)
		}
	}
}

func TestGetStartEndDates(t *testing.T) {
	GetStartEndDates(time.January)
}
