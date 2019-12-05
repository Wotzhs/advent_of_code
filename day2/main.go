package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	fmt.Printf("part 1 solution: %v\n", getSolution(12, 2))
}

func part2() {
	input1count := 0
	input2count := 0
	for {
		sum := getSolution(input1count, input2count)
		if sum == 19690720 {
			fmt.Printf("found combination input1: %v input2: %v solution: %v\n", input1count, input2count, 100*input1count+input2count)
			break
		}

		if input1count < 99 {
			input1count = input1count + 1

		} else if input1count >= 99 {
			input1count = 0
			input2count = input2count + 1
		}
	}
}

func parseInt(strInt string) int {
	parsedInt, err := strconv.Atoi(strInt)
	if err != nil {
		panic(err)
	}
	return parsedInt
}

func getSolution(input1, input2 int) int {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), ",")

	opCode := 1

	input[1] = fmt.Sprintf("%d", input1)
	input[2] = fmt.Sprintf("%d", input2)

operation:
	for {
		operator := input[opCode*4-4]
		if operator == "99" {
			break operation
		}
		input1 := input[opCode*4-2]
		input2 := input[opCode*4-3]
		resPos := input[opCode*4-1]

		pos, err := strconv.Atoi(resPos)
		if err != nil {
			panic(err)
		}

		input1Int, err := strconv.Atoi(input1)
		if err != nil {
			panic(err)
		}

		input2Int, err := strconv.Atoi(input2)
		if err != nil {
			panic(err)
		}

		switch operator {
		case "1":
			input[pos] = fmt.Sprintf("%d", parseInt(input[input1Int])+parseInt(input[input2Int]))
		case "2":
			input[pos] = fmt.Sprintf("%d", parseInt(input[input1Int])*parseInt(input[input2Int]))
		case "99":
			break operation
		}
		opCode = opCode + 1
	}

	return parseInt(input[0])
}
