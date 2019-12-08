package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	part1()
}

func part1() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")
	input1, input2 := getLine(input[0]), getLine(input[1])
	line1Moves, line2Moves := input1, input2
	if len(input1) < len(input2) {
		line1Moves, line2Moves = input2, input1
	}

	line2Path := [][]float64{}

	head := []float64{0, 0}
	for _, move := range line2Moves {
		axis, val := getDirection(move)
		head[axis] = head[axis] + val
		movement := []float64{head[0], head[1]}
		line2Path = append(line2Path, movement)
	}

	head = []float64{0, 0}
	tail := []float64{0, 0}
	const x int = 0
	const y int = 1
	intersections := [][]float64{}
	for _, move := range line1Moves {
		axis, val := getDirection(move)
		head[axis] = head[axis] + val

		changedYAxis := head[y] != tail[y]
		changedXAxis := head[x] != tail[x]

		if changedYAxis {
			minY, maxY := math.Min(head[y], tail[y]), math.Max(head[y], tail[y])
			for j, coordinates := range line2Path {
				if j < 1 {
					continue
				}
				minX, maxX := math.Min(line2Path[j-1][x], coordinates[x]), math.Max(line2Path[j-1][x], coordinates[x])
				if (minY <= line2Path[j-1][y] && maxY >= line2Path[j][y]) &&
					minX <= head[x] && maxX >= head[x] {
					intersections = append(intersections, []float64{head[x], line2Path[j][y]})
				}
			}
		}

		if changedXAxis {
			minX, maxX := math.Min(head[x], tail[x]), math.Max(head[x], tail[x])
			for j, coordinates := range line2Path {
				if j < 1 {
					continue
				}
				minY, maxY := math.Min(line2Path[j-1][y], coordinates[y]), math.Max(line2Path[j-1][y], coordinates[y])
				if (minX <= line2Path[j-1][x] && maxX >= line2Path[j][x]) &&
					minY <= head[y] && maxY >= head[y] {
					intersections = append(intersections, []float64{line2Path[j][x], head[y]})
				}
			}
		}

		tail[x], tail[y] = head[x], head[y]
	}

	var smallest float64
	for i, coords := range intersections {
		dist := math.Abs(coords[x]) + math.Abs(coords[y])
		if i == 0 {
			smallest = dist
		}
		if dist < smallest {
			smallest = dist
		}
	}

	fmt.Printf("smallest Manhattan distance %v\n", smallest)
}

func getDirection(input string) (idx int, val float64) {
	switch string(input[0]) {
	case "R":
		idx, val = 0, parseFloat(input[1:])
	case "L":
		idx, val = 0, -parseFloat(input[1:])
	case "U":
		idx, val = 1, parseFloat(input[1:])
	case "D":
		idx, val = 1, -parseFloat(input[1:])
	}
	return
}

func parseFloat(str string) float64 {
	const bitsize10 int = 10
	val, err := strconv.ParseFloat(str, bitsize10)
	if err != nil {
		panic(err)
	}
	return val
}

func getLine(input string) []string {
	return strings.Split(input, ",")
}
