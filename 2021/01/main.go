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

	depthIncreases, err := dayOnePartOne(string(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number of depth increases: %v\n", depthIncreases)

	slidingWindowDepthIncreases, err := dayOnePartTwo(string(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number of sliding-window depth increases: %v\n", slidingWindowDepthIncreases)
}

func dayOnePartOne(input string) (depthIncreases int, err error) {
	depths, err := parseDepthReadings(input)
	if err != nil {
		return 0, err
	}
	return measureDepthIncreases(depths), nil
}

func dayOnePartTwo(input string) (depthIncreases int, err error) {
	depths, err := parseDepthReadings(input)
	if err != nil {
		return 0, err
	}
	return measureSlidingWindowDepthIncreases(depths), nil
}

func parseDepthReadings(input string) ([]int, error) {
	inputs := strings.Split(input, "\n")
	depths := make([]int, 0, len(inputs))
	for _, depthStr := range inputs {
		depth, err := strconv.Atoi(depthStr)
		if err != nil {
			return nil, err
		}
		depths = append(depths, depth)
	}
	return depths, nil
}

func measureDepthIncreases(depths []int) (depthIncreases int) {
	for i, depth := range depths {
		if i > 0 && depth > depths[i-1] {
			depthIncreases++
		}
	}
	return depthIncreases
}

func measureSlidingWindowDepthIncreases(depths []int) (depthIncreases int) {
	slidingWindows := calculateSlidingWindows(depths)
	return measureDepthIncreases(slidingWindows)
}

func calculateSlidingWindows(depths []int) []int {
	var slidingWindows []int

	const windowWidth = 3

	for i := range depths {
		if i < windowWidth-1 {
			continue
		}
		var windowVal int
		for j := 0; j < windowWidth; j++ {
			windowVal += depths[i-j]
		}
		slidingWindows = append(slidingWindows, windowVal)
	}

	return slidingWindows
}
