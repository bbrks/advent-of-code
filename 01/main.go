package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

var errNoBasement = errors.New("Basement never reached!")

// Returns the final floor, and the position of the first time you reach the basement
func dayOne(inpt *string) (int, int, error) {
	var (
		floor int
		pos   = -1
	)
	for i, c := range *inpt {
		if string(c) == "(" {
			floor++
		} else if string(c) == ")" {
			floor--
		}

		if floor == -1 && pos == -1 {
			pos = i + 1
		}
	}
	return floor, pos, nil
}

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	input := string(dat)

	finalFloor, basementPos, err := dayOne(&input)
	if err == errNoBasement {
		fmt.Println(err)
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("Basement Floor Position:", basementPos)
	}
	fmt.Println("Final Floor:", finalFloor)
}
