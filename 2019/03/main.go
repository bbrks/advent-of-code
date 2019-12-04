package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	dist := FindClosestCrossedWire(string(b))
	fmt.Printf("Manhattan distance of closest intersection: %v\n", dist)

	dist = FindClosestCrossedWireByLength(string(b))
	fmt.Printf("Shortest length to intersection: %v\n", dist)
}

func FindClosestCrossedWire(input string) (distance int) {
	wires := strings.Split(input, "\n")
	if len(wires) < 2 {
		panic("len(wires) < 2")
	}

	path1 := markVisited(strings.Split(wires[0], ","))
	path2 := markVisited(strings.Split(wires[1], ","))

	i := intersections(path1, path2)
	closest := closestIntersectionManhattanDistance(i)
	return manhattanDistance(closest)
}

func FindClosestCrossedWireByLength(input string) (distance int) {
	wires := strings.Split(input, "\n")
	if len(wires) < 2 {
		panic("len(wires) < 2")
	}

	path1 := markVisited(strings.Split(wires[0], ","))
	path2 := markVisited(strings.Split(wires[1], ","))
	i := intersections(path1, path2)
	return closestIntersectionLength(i)
}

func intersections(path1, path2 map[xyCoord]int) map[xyCoord][2]int {
	intersections := make(map[xyCoord][2]int)
	for coord, path2Length := range path2 {
		if path1Length, ok := path1[coord]; ok {
			// Path 1 already visited here
			intersections[coord] = [2]int{path1Length, path2Length}
		}
	}
	return intersections
}

func closestIntersectionLength(i map[xyCoord][2]int) (closest int) {
	for intersection, lengths := range i {
		// skip 0,0
		if intersection.x == 0 && intersection.y == 0 {
			continue
		}
		if closest == 0 || lengths[0]+lengths[1] < closest {
			closest = lengths[0] + lengths[1]
		}
	}
	return closest
}

func closestIntersectionManhattanDistance(i map[xyCoord][2]int) (closest xyCoord) {
	for intersection := range i {
		// skip 0,0
		if intersection.x == 0 && intersection.y == 0 {
			continue
		}
		if manhattanDistance(closest) == 0 || manhattanDistance(intersection) < manhattanDistance(closest) {
			closest = intersection
		}
	}
	return closest
}

func markVisited(directions []string) map[xyCoord]int {
	// start in the middle
	x, y := 0, 0
	visited := map[xyCoord]int{
		xyCoord{x, y}: 0,
	}

	var totalDistance int

	for _, direction := range directions {
		distance, err := strconv.Atoi(direction[1:])
		if err != nil {
			panic(err)
		}
		for ; distance > 0; distance-- {
			switch direction[0] {
			case 'U':
				y++
			case 'D':
				y--
			case 'L':
				x--
			case 'R':
				x++
			}
			totalDistance++
			visited[xyCoord{x, y}] = totalDistance
		}
	}

	return visited
}

func manhattanDistance(coord xyCoord) int {
	return modInt(coord.x) + modInt(coord.y)
}

func modInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type xyCoord struct {
	x int
	y int
}
