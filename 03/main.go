package main

import (
	"fmt"
	"io/ioutil"
)

func sphericalHouses(i *string) int {
	return 0
}

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	input := string(dat)

	fmt.Printf("Houses with more than one present: %d", sphericalHouses(&input))
}
