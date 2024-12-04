package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	val, err := partOne(string(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day one part one: %d\n", val)

	val, err = partTwo(string(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day one part two: %d\n", val)
}

func partOne(input string) (int, error) {
	instructions, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	val := 0
	for _, ins := range instructions {
		switch ins.op {
		case "mul":
			val += ins.a * ins.b
		}
	}
	return val, nil
}

func partTwo(input string) (int, error) {
	instructions, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	val := 0
	do := true
	for _, ins := range instructions {
		switch ins.op {
		case "do":
			do = true
		case "don't":
			do = false
		case "mul":
			if do {
				val += ins.a * ins.b
			}
		}
	}
	return val, nil
}

type instruction struct {
	a, b int
	op   string
}

func parseInput(input string) ([]instruction, error) {
	reports := make([]instruction, 0)

	mulRegexp := regexp.MustCompile(`(mul|do|don't)\((\d*),?(\d*)\)`)
	matches := mulRegexp.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		if len(match) != 4 {
			return nil, fmt.Errorf("invalid input %q", match)
		}
		var a, b int
		if match[2] != "" {
			ai64, err := strconv.ParseInt(match[2], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid input %q: %w", match[2], err)
			}
			a = int(ai64)
		}
		if match[3] != "" {
			bi64, err := strconv.ParseInt(match[3], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid input %q: %w", match[3], err)
			}
			b = int(bi64)
		}
		reports = append(reports, instruction{a: a, b: b, op: match[1]})
	}

	return reports, nil
}
