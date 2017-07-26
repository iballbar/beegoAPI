package utils_test

import (
	"testing"
	"time"

	"github.com/iballbar/beegoAPI/utils"
)

func TestInTimeSpan(t *testing.T) {
	tt := []struct {
		name   string
		start  string
		end    string
		check  string
		expect bool
		err    string
	}{
		{name: "Test 2017/01/01 00:00:00 in 2017/01/01 00:00:00 and 2017/01/01 01:00:00", start: "2017-01-01T00:00:00+00:00", end: "2017-01-01T01:00:00+00:00", check: "2017-01-01T00:00:00+00:00", expect: false},
		{name: "Test 2017/01/01 00:00:01 in 2017/01/01 00:00:00 and 2017/01/01 01:00:00", start: "2017-01-01T00:00:00+00:00", end: "2017-01-01T01:00:00+00:00", check: "2017-01-01T00:00:01+00:00", expect: true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			startTime, err := time.Parse(time.RFC3339, tc.start)
			if err != nil {
				t.Fatal("cloud not parse start time:", err)
			}
			endTime, err := time.Parse(time.RFC3339, tc.end)
			if err != nil {
				t.Fatal("cloud not parse end time:", err)
			}
			checkTime, err := time.Parse(time.RFC3339, tc.check)
			if err != nil {
				t.Fatal("cloud not parse check time:", err)
			}
			result := utils.InTimeSpan(startTime, endTime, checkTime)
			if result != tc.expect {
				t.Errorf("expected %v got: %v", tc.expect, result)
			}
		})
	}
}
