package main

import (
	"testing"
)

func TestDayOne(t *testing.T) {
	tests := []struct {
		input                  string
		expectedDepthIncreases int
		expectedWindows        []int
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
			[]int{607, 618, 618, 617, 647, 716, 769, 792},
		},
	}

	for _, test := range tests {
		t.Run("part 1", func(t *testing.T) {
			output, err := dayOnePartOne(test.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if output != test.expectedDepthIncreases {
				t.Errorf("expected %d but got %d", test.expectedDepthIncreases, output)
			}
		})
		t.Run("part 2", func(t *testing.T) {
			depths, err := parseDepthReadings(test.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			output := calculateSlidingWindows(depths)
			for i, oVal := range output {
				if test.expectedWindows[i] != oVal {
					t.Errorf("unknown sliding window val for idx %d - %v expected %v", i, oVal, test.expectedWindows[i])
				}
			}
		})
	}
}
