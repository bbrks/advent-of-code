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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		startX, startY, endX, endY, cmd := parseInput(line[0])
		ret := flickLights(startX, startY, endX, endY, cmd)
		expect, err := strconv.Atoi(line[1])
		if err == nil && ret != expect {
			t.Errorf("Test failed, got: '%d', expected: '%d'", ret, expect)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
