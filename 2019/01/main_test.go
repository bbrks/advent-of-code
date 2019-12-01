package main

import (
	"strconv"
	"testing"
)

func TestFuelRequirement(t *testing.T) {
	tests := []struct {
		input           int
		expectedOutput1 int
		expectedOutput2 int
	}{
		{12, 2, 2},
		{14, 2, 2},
		{1969, 654, 966},
		{100756, 33583, 50346},
	}

	const (
		totalOutput1 = 34241
		totalOutput2 = 51316
	)

	for _, test := range tests {
		t.Run(strconv.Itoa(test.input), func(t *testing.T) {
			output := FuelRequirement1(test.input)
			if output != test.expectedOutput1 {
				t.Errorf("output 1 %d did not match expected %d", output, test.expectedOutput1)
			}
			output = FuelRequirement2(test.input)
			if output != test.expectedOutput2 {
				t.Errorf("output 2 %d did not match expected %d", output, test.expectedOutput2)
			}
		})
	}

	t.Run("TotalFuelRequirements", func(t *testing.T) {
		var input []int
		for _, test := range tests {
			input = append(input, test.input)
		}
		output := TotalFuelRequirements(FuelRequirement1, input...)
		if output != totalOutput1 {
			t.Errorf("total output 1 %d did not match expected %d", output, totalOutput1)
		}
		output = TotalFuelRequirements(FuelRequirement2, input...)
		if output != totalOutput2 {
			t.Errorf("total output 2 %d did not match expected %d", output, totalOutput2)
		}
	})
}
