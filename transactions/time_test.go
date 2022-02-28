package transactions

import "testing"

var j = []string{"January", "january", "Jan", "jan", "1"}
var d = []string{"December", "december", "Dec", "dec", "12"}
var months = Months()

func TestMonths(t *testing.T) {
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
