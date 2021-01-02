package main

import (
	"strings"
	"testing"
)

func TestASDF(t *testing.T) {
	tests := []struct {
		input                 string
		expectedDistance      int
		expectedLength        int
		expectedIntersections []xyCoord
	}{
		{
			`R8,U5,L5,D3
U7,R6,D4,L4`,
			6,
			30,
			[]xyCoord{{0, 0}, {3, 3}, {6, 5}},
		},
		{
			`R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`,
			159,
			610,
			[]xyCoord{{0, 0}, {146, 46}, {155, 4}, {155, 11}, {158, -12}},
		},
		{
			`R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`,
			135,
			410,
			[]xyCoord{{0, 0}, {107, 47}, {107, 51}, {107, 71}, {124, 11}, {157, 18}},
		},
	}

	for _, test := range tests {
		t.Run("intersections", func(t *testing.T) {
			wires := strings.Split(test.input, "\n")
			if len(wires) < 2 {
				panic("len(wires) < 2")
			}
			path1 := markVisited(strings.Split(wires[0], ","))
			path2 := markVisited(strings.Split(wires[1], ","))
			i := intersections(path1, path2)
			if len(i) != len(test.expectedIntersections) {
				t.Errorf("Expected %d intersections but got %d", len(test.expectedIntersections), len(i))
			}
			for _, val := range test.expectedIntersections {
				var found bool
				for val2 := range i {
					if val2 == val {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected intersection %v but couldn't find it in %v", val, i)
				}
			}
		})

		t.Run("FindClosestCrossedWire", func(t *testing.T) {
			dist := FindClosestCrossedWire(test.input)
			if dist != test.expectedDistance {
				t.Errorf("Expected distance %d but got %d", test.expectedDistance, dist)
			}
		})

		t.Run("FindClosestCrossedWire", func(t *testing.T) {
			dist := FindClosestCrossedWireByLength(test.input)
			if dist != test.expectedLength {
				t.Errorf("Expected length %d but got %d", test.expectedLength, dist)
			}
		})
	}
}
