package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
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
	fmt.Println(fuelRequirement)
}
