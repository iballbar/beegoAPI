package byte_test

import (
	"testing"

	byte_utils "github.com/iballbar/beegoAPI/utils/byte"
)

func TestByteToBinaryString(t *testing.T) {
	tt := []struct {
		name   string
		value  byte
		expect string
		err    string
	}{
		{name: "zero", value: 0, expect: "00000000"},
		{name: "one", value: 1, expect: "00000001"},
		{name: "two", value: 2, expect: "00000010"},
		{name: "not send data", expect: "00000000"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := byte_utils.ByteToBinaryString(tc.value)
			if result != tc.expect {
				t.Errorf("Expected %s got: %v", tc.expect, result)
			}
		})
	}
}
