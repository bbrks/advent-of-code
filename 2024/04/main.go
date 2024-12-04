package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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

	val, err := partOne(string(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day one part one: %d\n", val)

	val, err = partTwo(string(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day one part two: %d\n", val)
}

func partOne(input string) (int, error) {
	matrix, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	val := 0
	for y, row := range matrix {
		for x, _ := range row {
			// iterate from this position in every direction for a continuation of the word
			directions := dirs(N, E, S, W, NE, SE, SW, NW)
			val += foundWords(matrix, x, y, "XMAS", directions)
		}
	}
	return val, nil
}

func boundsSafeRune(matrix [][]rune, x, y int) rune {
	if y < 0 || y >= len(matrix) {
		return -1
	}
	if x < 0 || x >= len(matrix[y]) {
		return -1
	}
	char := matrix[y][x]
	return char
}

type directions uint8

const (
	N = 1 << iota
	E
	S
	W
	NE
	SE
	SW
	NW
)

func dirs(directions ...int) (d directions) {
	for _, direction := range directions {
		d.set(direction)
	}
	return d
}

func (d *directions) set(direction int) {
	*d |= directions(direction)
}

func (d *directions) toggle(direction int) {
	*d ^= directions(direction)
}

func (d *directions) enabled(direction int) bool {
	return *d&directions(direction) != 0
}

func foundWords(matrix [][]rune, originX, originY int, word string, directions directions) (results int) {
	// early return
	if r := boundsSafeRune(matrix, originX, originY); r != rune(word[0]) {
		return 0
	}

	for i, char := range word {
		if i == 0 {
			// we already checked the first char in the early return
			continue
		}
		endOfWord := i == len(word)-1

		// LTR
		if directions.enabled(E) {
			if boundsSafeRune(matrix, originX+i, originY) != char {
				directions.toggle(E)
			} else if endOfWord {
				results++
			}
		}
		// RTL
		if directions.enabled(W) {
			if boundsSafeRune(matrix, originX-i, originY) != char {
				directions.toggle(W)
			} else if endOfWord {
				results++
			}
		}
		if directions.enabled(S) {
			if boundsSafeRune(matrix, originX, originY+i) != char {
				directions.toggle(S)
			} else if endOfWord {
				results++
			}
		}
		if directions.enabled(N) {
			if boundsSafeRune(matrix, originX, originY-i) != char {
				directions.toggle(N)
			} else if endOfWord {
				results++
			}
		}

		// ordinals (diagonals)
		if directions.enabled(SE) {
			if boundsSafeRune(matrix, originX+i, originY+i) != char {
				directions.toggle(SE)
			} else if endOfWord {
				results++
			}
		}
		if directions.enabled(SW) {
			if boundsSafeRune(matrix, originX-i, originY+i) != char {
				directions.toggle(SW)
			} else if endOfWord {
				results++
			}
		}
		if directions.enabled(NE) {
			if boundsSafeRune(matrix, originX+i, originY-i) != char {
				directions.toggle(NE)
			} else if endOfWord {
				results++
			}
		}
		if directions.enabled(NW) {
			if boundsSafeRune(matrix, originX-i, originY-i) != char {
				directions.toggle(NW)
			} else if endOfWord {
				results++
			}
		}
	}

	return results
}

func partTwo(input string) (int, error) {
	matrix, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	val := 0
	for y, row := range matrix {
		for x, _ := range row {
			// find candidates for middle of the X-MAS
			if boundsSafeRune(matrix, x, y) != 'A' {
				continue
			}
			// back up and left, and look SE for MAS/SAM
			if foundWords(matrix, x-1, y-1, "MAS", dirs(SE)) > 0 ||
				foundWords(matrix, x-1, y-1, "SAM", dirs(SE)) > 0 {
				// found half - now down and left and look NE
				if foundWords(matrix, x-1, y+1, "MAS", dirs(NE)) > 0 ||
					foundWords(matrix, x-1, y+1, "SAM", dirs(NE)) > 0 {
					val++
				}
			}
		}
	}
	return val, nil
}

func parseInput(input string) ([][]rune, error) {
	matrix := make([][]rune, 0)

	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		matrix = append(matrix, []rune(s.Text()))
	}

	return matrix, nil
}
