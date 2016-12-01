package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestIsNice(t *testing.T) {
	file, err := os.Open("input_test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		ret := isNice(&line[0])
		expect, err := strconv.ParseBool(line[1])
		if err == nil && ret != expect {
			t.Errorf("Test failed, got: '%t', expected: '%t'", ret, expect)
		}
		ret = isNice2(&line[0])
		expect, err = strconv.ParseBool(line[2])
		if err == nil && ret != expect {
			t.Errorf("Test failed, got: '%t', expected: '%t'", ret, expect)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
