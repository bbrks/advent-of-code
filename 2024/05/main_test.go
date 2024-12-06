package main

import (
	"fmt"
	"testing"
)

func TestDayFive(t *testing.T) {
	tests := []struct {
		input           string
		expectedPartOne int
		expectedPartTwo int
	}{
		{
			input: `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`,
			expectedPartOne: 143,
			expectedPartTwo: 123,
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
