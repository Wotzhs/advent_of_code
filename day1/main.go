package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(data), "\n")

	var fuelRequirement int
	for _, mass := range input {
		intMass, err := strconv.Atoi(mass)
		if err != nil {
			log.Fatal(err)
		}
		fuelRequirement = fuelRequirement + (intMass/3 - 2)
	}
	fmt.Printf("part 1 solution: %d\n", fuelRequirement)
}

func part2() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(data), "\n")

	var fuelRequirement int
	for _, mass := range input {
		intMass, err := strconv.Atoi(mass)
		if err != nil {
			log.Fatal(err)
		}
		fuelRequirement = fuelRequirement + calculateRequiredFuel(intMass)
	}
	fmt.Printf("part 2 solution: %d\n", fuelRequirement)
}

func calculateRequiredFuel(mass int) int {
	requiredFuel := mass/3 - 2

	if requiredFuel > 0 {
		requiredFuel = requiredFuel + calculateRequiredFuel(requiredFuel)
	}

	if requiredFuel < 0 {
		return 0
	}

	return requiredFuel
}
