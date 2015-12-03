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
func TestSphericalHouses(t *testing.T) {

	file, err := os.Open("input_test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		ret := sphericalHouses(&line[1])
		expect, _ := strconv.Atoi(line[0])
		if ret != expect {
			t.Errorf("Test failed, got: '%d', expected: '%d'", ret, expect)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
