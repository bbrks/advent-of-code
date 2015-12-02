package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ribbonLength(b []int) int {
	length := 2*b[0] + 2*b[1]
	bow := b[0] * b[1] * b[2]
	return length + bow
}

// Returns the surface area of a rectangular cuboid
// plus the area of the smallest side as extra
func surfaceArea(b []int) int {
	surfaceArea := 2*b[0]*b[1] + 2*b[0]*b[2] + 2*b[1]*b[2]
	extra := b[0] * b[1] // smallest side
	return surfaceArea + extra
}

func main() {
	sqFt := 0
	ribbon := 0

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var b []int
		ds := strings.Split(scanner.Text(), "x")
		for _, d := range ds {
			i, err := strconv.Atoi(d)
			if err != nil {
				panic(err)
			}
			b = append(b, i)
		}
		sort.Ints(b) // so we can get the smallest side by [0]*[1]
		sqFt = sqFt + surfaceArea(b)
		ribbon = ribbon + ribbonLength(b)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total sqft of wrapping paper needed:", sqFt)
	fmt.Println("Total length of ribbon needed:", ribbon)
}
