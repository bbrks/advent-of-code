package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func RunIntcode(input []int) []int {
	output := make([]int, len(input))
	for i, val := range input {
		output[i] = val
	}

	for i := 0; i < len(input); i += 4 {
		switch output[i] {
		case 1:
			// add
			output[output[i+3]] = output[output[i+1]] + output[output[i+2]]
		case 2:
			// multiply
			output[output[i+3]] = output[output[i+1]] * output[output[i+2]]
		case 99:
			// halt
			break
		}
	}
	return output
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		intsStr := strings.Split(scanner.Text(), ",")
		for _, intStr := range intsStr {
			i, err := strconv.Atoi(intStr)
			if err != nil {
				log.Fatal(err)
			}
			input = append(input, i)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// modify input
	input[1] = 12
	input[2] = 2
	output := RunIntcode(input)
	fmt.Printf("Part 1: Position 0: %v\n", output[0])

}
