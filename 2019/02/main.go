package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func RunIntcode(input []int) {
	for i := 0; i < len(input); i += 4 {
		switch input[i] {
		case 1:
			// add
			input[input[i+3]] = input[input[i+1]] + input[input[i+2]]

		case 2:
			// multiply
			input[input[i+3]] = input[input[i+1]] * input[input[i+2]]
		case 99:
			// halt
			break
		}
	}
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

	RunIntcode(input)
	fmt.Printf("Position 0: %v\n", input[0])

}
