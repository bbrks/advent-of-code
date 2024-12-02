package main

import (
	"bufio"
	"fmt"
	"io"
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

	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	val, err := dayTwoPartOne(string(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day one part one: %d\n", val)

	val, err = dayTwoPartTwo(string(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day one part two: %d\n", val)
}

func dayTwoPartOne(input string) (int, error) {
	reports, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	numSafe := 0
	for _, report := range reports {
		if isReportSafe(report) {
			numSafe++
		}
	}
	return numSafe, nil
}

const (
	safeLevelJumpMin = 1
	safeLevelJumpMax = 3
)

// reports true if the report is safe. The report is considered safe if:
// - The levels are either all increasing or all decreasing.
// - Any two adjacent levels differ by at least one and at most three.
func isReportSafe(report []int) bool {
	return reportSafe(report)
}

// reports true if the report is safe. The report is considered safe if:
// - The levels are either all increasing or all decreasing.
// - Any two adjacent levels differ by at least one and at most three.
func isReportSafeWithDampening(report []int) bool {
	if reportSafe(report) {
		return true
	}
	// brute force - remove elements one at a time until we're safe
	for i := range report {
		newReport := make([]int, 0, len(report)-1)
		newReport = append(newReport, report[:i]...)
		newReport = append(newReport, report[i+1:]...)
		if isReportSafe(newReport) {
			return true
		}
	}
	return false
}

// reportSafe returns the index of the first bad level change and a boolean indicating if the report is safe.
func reportSafe(report []int) bool {
	var isIncreasing bool
	for i, curr := range report {
		if i == 0 {
			continue
		}

		prev := report[i-1]

		// distance check
		if d := dist(prev, curr); d < safeLevelJumpMin || d > safeLevelJumpMax {
			return false
		}

		// increasing/decreasing
		if curr > prev {
			if !isIncreasing {
				if i > 1 {
					return false
				}
				isIncreasing = true
			}
		} else if isIncreasing {
			return false
		}
	}

	return true
}

func dist(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func dayTwoPartTwo(input string) (int, error) {
	reports, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	numSafe := 0
	for _, report := range reports {
		if isReportSafeWithDampening(report) {
			numSafe++
		}
	}
	return numSafe, nil
}

func parseInput(input string) ([][]int, error) {
	reports := make([][]int, 0)
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		levels := strings.Split(s.Text(), " ")
		parsedLevels := make([]int, 0, len(levels))
		for _, level := range levels {
			val, err := strconv.ParseInt(level, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid input %q: %w", level, err)
			}
			parsedLevels = append(parsedLevels, int(val))
		}
		reports = append(reports, parsedLevels)
	}
	return reports, nil
}
