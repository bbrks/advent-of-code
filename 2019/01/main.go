package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// FuelRequirement calculates the required fuel for a single module mass.
func FuelRequirement(moduleMass int) (fuelRequired int) {
	return moduleMass/3 - 2
}

// TotalFuelRequirements calculates the sum of the required fuel for the given module masses.
func TotalFuelRequirements(moduleMasses ...int) (totalFuelRequired int) {
	for _, moduleMass := range moduleMasses {
		totalFuelRequired += FuelRequirement(moduleMass)
	}
	return totalFuelRequired
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var moduleMasses []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		moduleMass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		moduleMasses = append(moduleMasses, moduleMass)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total fuel required: %v", TotalFuelRequirements(moduleMasses...))

}
