package main

import (
	"testing"
)

func TestDayOne(t *testing.T) {
	tests := []struct {
		input                 string
		expectedDepthIncreases      int
	}{
		{
			`199
200
208
210
200
207
240
269
260
263`,
			7,
		},
	}

	for _, test := range tests {
		t.Run("intersections", func(t *testing.T) {
			output, err := dayOne(test.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if output != test.expectedDepthIncreases {
				t.Errorf("expected %d but got %d", test.expectedDepthIncreases, output)
			}
		})
	}
}
