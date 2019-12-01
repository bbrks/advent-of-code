package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// FuelRequirement1 calculates the required fuel for a single module mass for part 1.
func FuelRequirement1(moduleMass int) (fuelRequired int) {
	return moduleMass/3 - 2
}

// FuelRequirement2 calculates the required fuel for a single module mass for part 2.
func FuelRequirement2(mass int) (fuelRequired int) {
	fuelRequired = FuelRequirement1(mass)
	// clamp to zero
	if fuelRequired < 0 {
		return 0
	}
	fuelRequired += FuelRequirement2(fuelRequired)
	return fuelRequired
}

// TotalFuelRequirements calculates the sum of the required fuel for the given module masses.
func TotalFuelRequirements(calcFn func(int) int, moduleMasses ...int) (totalFuelRequired int) {
	for _, moduleMass := range moduleMasses {
		totalFuelRequired += calcFn(moduleMass)
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

	fmt.Printf("Total fuel required part 1: %v\n", TotalFuelRequirements(FuelRequirement1, moduleMasses...))
	fmt.Printf("Total fuel required part 2: %v\n", TotalFuelRequirements(FuelRequirement2, moduleMasses...))

}
