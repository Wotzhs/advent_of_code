package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	firstIntersections, line1IntersectingMoves, line2IntersectingMoves := part1()
	part2New(firstIntersections, line1IntersectingMoves, line2IntersectingMoves)
	fmt.Printf("elapsed %.2fms\n", float64(time.Now().Sub(start).Nanoseconds())/float64(time.Millisecond))
}

func part1() ([][]float64, []int, []int) {
	line1Moves, line2Moves := getInputs()
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
	var line1IntersectingMoves []int
	var line2IntersectingMoves []int

	for i, move := range line1Moves {
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

					line1IntersectingMoves = append(line1IntersectingMoves, i)
					line2IntersectingMoves = append(line2IntersectingMoves, j)
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

					line1IntersectingMoves = append(line1IntersectingMoves, i)
					line2IntersectingMoves = append(line2IntersectingMoves, j)
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

	fmt.Printf("part1: smallest manhattan distance %v\n", smallest)
	return intersections, line1IntersectingMoves, line2IntersectingMoves
}

func part2New(intersections [][]float64, line1IntersectingMoves, line2IntersectingMoves []int) {
	line1Moves, line2Moves := getInputs()
	const x int = 0
	const y int = 1
	steps := []float64{}
	for i, coordinates := range intersections {

		line1Head := []float64{0, 0}
		line1Steps := 0.0
		for _, move := range line1Moves[:line1IntersectingMoves[i]] {
			axis, val := getDirection(move)
			line1Head[axis] = line1Head[axis] + val
			line1Steps = line1Steps + math.Abs(val)
		}

		stepsToIntersection := math.Abs(coordinates[x] - line1Head[x] + coordinates[y] - line1Head[y])
		line1Steps = line1Steps + stepsToIntersection

		line2Head := []float64{0, 0}
		line2Steps := 0.0
		for _, move := range line2Moves[:line2IntersectingMoves[i]] {
			axis, val := getDirection(move)
			line2Head[axis] = line2Head[axis] + val
			line2Steps = line2Steps + math.Abs(val)
		}

		stepsToIntersection = math.Abs(coordinates[x] - line2Head[x] + coordinates[y] - line2Head[y])
		line2Steps = line2Steps + stepsToIntersection

		steps = append(steps, line1Steps+line2Steps)
	}

	sort.Slice(steps, func(i, j int) bool {
		return steps[i] < steps[j]
	})
	fmt.Printf("part2: fewest combined steps %v\n", steps[0])
}

func getInputs() ([]string, []string) {
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

	return line1Moves, line2Moves
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
