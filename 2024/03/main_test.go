package main

import "testing"

func TestDayThree(t *testing.T) {
	tests := []struct {
		input           string
		expectedPartOne int
		expectedPartTwo int
	}{
		{
			input:           `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`,
			expectedPartOne: 161,
			expectedPartTwo: 161,
		},
		{
			input:           `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`,
			expectedPartOne: 161,
			expectedPartTwo: 48,
		},
	}

	for _, test := range tests {
		t.Run("part 1", func(t *testing.T) {
			result, err := partOne(test.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != test.expectedPartOne {
				t.Errorf("expected %d but got %d", test.expectedPartOne, result)
			}
		})
		t.Run("part 2", func(t *testing.T) {
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
