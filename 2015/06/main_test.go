package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

// Read input_test line by line, with expected result and input and check
func TestFireHazard(t *testing.T) {

	file, err := os.Open("input_test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	part1Lights := lights{}
	part2Lights := lights{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		x1, y1, x2, y2, cmd := parseInput(line[0])
		part1Lights.flickLights(x1, y1, x2, y2, cmd)
		part2Lights.twiddleLights(x1, y1, x2, y2, cmd)
		expectedPart1, err := strconv.Atoi(line[1])
		if err == nil && part1Lights.totalLitCount != expectedPart1 {
			t.Errorf("Test failed, got: '%d', expected: '%d'", part1Lights.totalLitCount, expectedPart1)
		}
		expectedPart2, err := strconv.Atoi(line[2])
		if err == nil && part2Lights.totalBrightness != expectedPart2 {
			t.Errorf("Test failed, got: '%d', expected: '%d'", part2Lights.totalBrightness, expectedPart2)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
