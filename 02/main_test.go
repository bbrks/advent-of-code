package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestSurfaceArea(t *testing.T) {
	file, err := os.Open("input_test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		var b []int
		ds := strings.Split(line[0], "x")
		for _, d := range ds {
			i, err := strconv.Atoi(d)
			if err != nil {
				panic(err)
			}
			b = append(b, i)
		}
		ret := surfaceArea(b)
		expect, _ := strconv.Atoi(line[1])
		if ret != expect {
			t.Errorf("Test failed, got: '%d', expected: '%d'", ret, expect)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func TestRibbonLength(t *testing.T) {
	file, err := os.Open("input_test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		var b []int
		ds := strings.Split(line[0], "x")
		for _, d := range ds {
			i, err := strconv.Atoi(d)
			if err != nil {
				panic(err)
			}
			b = append(b, i)
		}
		ret := ribbonLength(b)
		expect, _ := strconv.Atoi(line[2])
		if ret != expect {
			t.Errorf("Test failed, got: '%d', expected: '%d'", ret, expect)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
