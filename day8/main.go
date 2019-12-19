package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	inputs := string(f)
	largest := 0
	largestBalance := ""

	for len(inputs) > 0 {
		size := 25 * 6
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

	fmt.Printf("total %v\n", total)
}
