package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
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

	d, err := dayOnePartOne(string(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day one part one: %d\n", d)

	s, err := dayOnePartTwo(string(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day one part two: %d\n", s)
}

func dayOnePartOne(input string) (int, error) {
	l1, l2, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	total := totalDist(l1, l2)
	return total, nil
}

func totalDist(l1, l2 []int) (totalDist int) {
	slices.Sort(l1)
	slices.Sort(l2)
	for i, val1 := range l1 {
		val2 := l2[i]
		d := dist(val1, val2)
		totalDist += d
	}
	return totalDist
}

func dist(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func dayOnePartTwo(input string) (int, error) {
	l1, l2, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	total := totalSimilarity(l1, l2)
	return total, nil
}

func totalSimilarity(l1, l2 []int) (totalSimilarity int) {
	l2Appearances := make(map[int]int)
	for _, val := range l2 {
		l2Appearances[val]++
	}
	for _, val1 := range l1 {
		similarity := val1 * l2Appearances[val1]
		totalSimilarity += similarity
	}
	return totalSimilarity
}

const listSeparator = "   "

func parseInput(input string) ([]int, []int, error) {
	l1, l2 := make([]int, 0), make([]int, 0)
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		listParts := strings.SplitN(s.Text(), listSeparator, 2)
		if len(listParts) != 2 {
			return nil, nil, fmt.Errorf("invalid input: %s (expected 2 distances, had %d)", s.Text(), len(listParts))
		}
		i1, err := strconv.ParseInt(listParts[0], 10, 64)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid input %q: %w", listParts[0], err)
		}
		i2, err := strconv.ParseInt(listParts[1], 10, 64)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid input %q: %w", listParts[1], err)
		}
		l1 = append(l1, int(i1))
		l2 = append(l2, int(i2))
	}
	return l1, l2, nil
}
