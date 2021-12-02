package main

import (
	"testing"
)

func TestDayTwo(t *testing.T) {
	tests := []struct {
		input                string
		expectedHorizontal   int
		expectedDepth        int
		expectedDepthWithAim int
	}{
		{
			`forward 5
down 5
forward 8
up 3
down 8
forward 2`,
			15,
			10,
			60,
		},
	}

	for _, test := range tests {
		horizontalPartOne, depth, err := dayTwoPartOne(test.input)
		horizontalPartTwo, depthWithAim, err := dayTwoPartTwo(test.input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if horizontalPartOne != horizontalPartTwo {
			t.Errorf("expected horizontal positions to be equal, but got %d and %d", horizontalPartOne, horizontalPartTwo)
		}
		if horizontalPartOne != test.expectedHorizontal {
			t.Errorf("expected horizontal position of %d but got %d", test.expectedHorizontal, horizontalPartOne)
		}
		t.Run("part 1", func(t *testing.T) {
			if depth != test.expectedDepth {
				t.Errorf("expected depth of %d but got %d", test.expectedDepth, depth)
			}
		})
		t.Run("part 2", func(t *testing.T) {
			if depthWithAim != test.expectedDepthWithAim {
				t.Errorf("expected depth with aim of %d but got %d", test.expectedDepthWithAim, depthWithAim)
			}
		})
	}
}
