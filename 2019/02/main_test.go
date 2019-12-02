package main

import (
	"testing"
)

func TestFuelRequirement(t *testing.T) {
	tests := []struct {
		input          []int
		expectedOutput []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			output := RunIntcode(test.input)
			for i, val := range output {
				if val != test.expectedOutput[i] {
					t.Errorf("output at idx %d in %v did not match expected %v", i, test.input, test.expectedOutput)
				}
			}
		})
	}
}
