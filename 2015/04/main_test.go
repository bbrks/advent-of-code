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
func TestAdventCoin(t *testing.T) {
	file, err := os.Open("input_test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		ret := adventCoin(&line[0], line[1])
		expect, err := strconv.Atoi(line[2])
		if err == nil && ret != expect {
			t.Errorf("Test failed, got: '%d', expected: '%d'", ret, expect)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
