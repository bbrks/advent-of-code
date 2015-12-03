package main

import (
	"fmt"
	"io/ioutil"
)

func sphericalHouses(inpt *string) int {
	x, y := 0, 0
	houses := make(map[string]int)

	// Initial house gets a present!
	houses[fmt.Sprintf("%d,%d", x, y)]++

	for _, c := range *inpt {
		if string(c) == ">" {
			x++
		} else if string(c) == "<" {
			x--
		} else if string(c) == "^" {
			y++
		} else if string(c) == "v" {
			y--
		}
		houses[fmt.Sprintf("%d,%d", x, y)]++
	}
	return len(houses)
}

func roboSanta(inpt *string) int {
	return 0
}

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	input := string(dat)

	houses := sphericalHouses(&input)
	fmt.Printf("Houses with more than one present: %d\n", houses)

	houses2 := roboSanta(&input)
	fmt.Printf("Houses with the help of Robo-Santa: %d\n", houses2)
}
