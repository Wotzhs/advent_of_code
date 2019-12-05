package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseInt(strInt string) int {
	parsedInt, err := strconv.Atoi(strInt)
	if err != nil {
		panic(err)
	}
	return parsedInt
}

func main() {
	part1()
}

func part1() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), ",")

	opCode := 1

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

	fmt.Printf("part 1 solution: %v\n", input[0])
}
