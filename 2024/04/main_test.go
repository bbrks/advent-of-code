package main

import (
	"fmt"
	"testing"
)

func TestDayFour(t *testing.T) {
	tests := []struct {
		input           string
		expectedPartOne int
		expectedPartTwo int
	}{
		{
			input: `..X...
.SAMX.
.A..A.
XMAS.S
.X....
`,
			expectedPartOne: 4,
		},
		{
			input: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			expectedPartOne: 18,
			expectedPartTwo: 9,
		},
		{
			input: `M.S
.A.
M.S`,
			expectedPartTwo: 1,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("part 1 %d", i), func(t *testing.T) {
			result, err := partOne(test.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != test.expectedPartOne {
				t.Errorf("expected %d but got %d", test.expectedPartOne, result)
			}
		})
		t.Run(fmt.Sprintf("part 2 %d", i), func(t *testing.T) {
			result, err := partTwo(test.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != test.expectedPartTwo {
				t.Errorf("expected %d but got %d", test.expectedPartTwo, result)
			}
		})
	}
}
