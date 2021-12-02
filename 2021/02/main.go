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

	horizontal, depth, err := dayTwoPartOne(string(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 1: Horizontal position (%v) * Final Depth (%v) = %v\n", horizontal, depth, horizontal*depth)

	horizontal, depth, err = dayTwoPartTwo(string(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 2: Horizontal position (%v) * Final Depth (%v) = %v\n", horizontal, depth, horizontal*depth)
}

type direction uint8

const (
	forward direction = iota
	down
	up
)

func getDirection(direction string) (direction, error) {
	switch direction {
	case "forward":
		return forward, nil
	case "down":
		return down, nil
	case "up":
		return up, nil
	}
	return 0, fmt.Errorf("unknown direction: %q", direction)
}

type move struct {
	direction direction
	distance  int
}

func parseCourseReadings(input string) (moves []move, err error) {
	inputs := strings.Split(input, "\n")
	moves = make([]move, 0, len(inputs))
	for _, courseStr := range inputs {
		course := strings.Split(courseStr, " ")
		if len(course) != 2 {
			return nil, fmt.Errorf("invalid course format (expecting 2 elements space separated): %q", courseStr)
		}
		distance, err := strconv.Atoi(course[1])
		if err != nil {
			return nil, err
		}
		dir, err := getDirection(course[0])
		if err != nil {
			return nil, err
		}
		moves = append(moves, move{dir, distance})
	}
	return moves, nil
}

func dayTwoPartOne(input string) (horizontal, depth int, err error) {
	moves, err := parseCourseReadings(input)
	if err != nil {
		return 0, 0, err
	}
	for _, v := range moves {
		switch v.direction {
		case forward:
			horizontal += v.distance
		case up:
			depth -= v.distance
		case down:
			depth += v.distance
		}
	}
	return horizontal, depth, nil
}

func dayTwoPartTwo(input string) (horizontal, depth int, err error) {
	moves, err := parseCourseReadings(input)
	if err != nil {
		return 0, 0, err
	}
	aim := 0
	for _, v := range moves {
		switch v.direction {
		case forward:
			horizontal += v.distance
			depth += v.distance * aim
		case up:
			aim -= v.distance
		case down:
			aim += v.distance
		}
	}
	return horizontal, depth, nil
}
