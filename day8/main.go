package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	width  int = 25
	height int = 6
)

func main() {
	part1()
	part2()
}

func part1() {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	inputs := string(f)
	largest := 0
	largestBalance := ""

	for len(inputs) > 0 {
		size := width * height
		if len(inputs) < size {
			size = len(inputs)
		}

		currentGroup := inputs[:size]
		currentGroup = strings.Replace(currentGroup, "0", "", -1)

		if len(currentGroup) > largest {
			largest = len(currentGroup)
			largestBalance = currentGroup
		}

		inputs = inputs[size:]
	}

	counters := map[string]int{}

	for _, val := range strings.Split(largestBalance, "") {
		counters[val] = counters[val] + 1
	}

	total := 1
	for _, count := range counters {
		total *= count
	}

	fmt.Printf("part 1 total: %v\n", total)
}

func part2() {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	inputs := string(f)

	image := []string{}

	for len(inputs) > 0 {
		size := width * height
		if len(inputs) < size {
			size = len(inputs)
		}

		currentGroup := inputs[:size]
		image = append(image, currentGroup)

		inputs = inputs[size:]
	}

	visual := [height][width]string{}
	for row, line := range visual {
		for column, _ := range line {
			for _, layer := range image {
				if string(layer[column+row*width]) != "2" {
					visual[row][column] = string(layer[column+row*width])
					break
				}
			}
		}
	}

	fmt.Println("part 2 answer:")
	fmt.Println("")

	for _, line := range visual {
		for _, pixel := range line {
			if pixel == "0" {
				fmt.Printf("%v", " ")
			}
			if pixel == "1" {
				fmt.Printf("%v", "x")
			}
		}
		fmt.Println("")
	}

	fmt.Println("")
}
