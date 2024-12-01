package main

import "testing"

func TestDayOne(t *testing.T) {
	tests := []struct {
		input              string
		expectedDistance   int
		expectedSimilarity int
	}{
		{
			input: `3   4
4   3
2   5
1   3
3   9
3   3`,
			expectedDistance:   11, // 2 + 1 + 0 + 1 + 2 + 5
			expectedSimilarity: 31, // 9 + 4 + 0 + 0 + 9 + 9
		},
	}

	for _, test := range tests {
		t.Run("part 1", func(t *testing.T) {
			distance, err := dayOnePartOne(test.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if distance != test.expectedDistance {
				t.Errorf("expected distance %d but got %d", test.expectedDistance, distance)
			}
		})
		t.Run("part 2", func(t *testing.T) {
			similarity, err := dayOnePartTwo(test.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if similarity != test.expectedSimilarity {
				t.Errorf("expected similarity %d but got %d", test.expectedSimilarity, similarity)
			}
		})
	}
}
