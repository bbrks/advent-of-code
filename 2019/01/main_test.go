package main

import (
	"strconv"
	"testing"
)

func TestFuelRequirement(t *testing.T) {
	tests := []struct {
		input          int
		expectedOutput int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.input), func(t *testing.T) {
			output := FuelRequirement(test.input)
			if output != test.expectedOutput {
				t.Errorf("output %d did not match expected %d", output, test.expectedOutput)
			}
		})
	}
}
