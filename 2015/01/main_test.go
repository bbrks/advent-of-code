package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestDayOne(t *testing.T) {
	file, err := os.Open("input_test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		floor, pos, _ := dayOne(&line[0])
		expectFloor, floorErr := strconv.Atoi(line[1])
		expectPos, posErr := strconv.Atoi(line[2])
		if floorErr == nil && floor != expectFloor {
			t.Errorf("Test failed, got: '%d', expected: '%d'", floor, expectFloor)
		}
		if posErr == nil && pos != expectPos {
			t.Errorf("Test failed, got: '%d', expected: '%d'", pos, expectPos)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
