package main

import "testing"

func TestDayTwo(t *testing.T) {
	tests := []struct {
		input                       string
		expectedSafeReports         int
		expectedSafeReportsDampened int
	}{
		{
			input: `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
			expectedSafeReports:         2, // s u u u u s
			expectedSafeReportsDampened: 4, // s u u s s s
		},
	}

	for _, test := range tests {
		t.Run("part 1", func(t *testing.T) {
			numSafe, err := dayTwoPartOne(test.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if numSafe != test.expectedSafeReports {
				t.Errorf("expected %d safe reports but got %d", test.expectedSafeReports, numSafe)
			}
		})
		t.Run("part 2", func(t *testing.T) {
			numSafe, err := dayTwoPartTwo(test.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if numSafe != test.expectedSafeReportsDampened {
				t.Errorf("expected %d safe with dampening but got %d", test.expectedSafeReportsDampened, numSafe)
			}
		})
	}
}
