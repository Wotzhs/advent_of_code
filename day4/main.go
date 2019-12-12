package main

import (
	"fmt"
	"strconv"
)

func main() {
	part1()
}

func part1() {
	var match int

	for i := 171309; i < 643603; i++ {
		str := strconv.Itoa(i)
	check:
		for j := range str {
			if j+1 > len(str)-1 {
				break
			}

			if parseInt(string(str[j])) > parseInt(string(str[j+1])) {
				break
			}

			if parseInt(string(str[j])) == parseInt(string(str[j+1])) {
				for k := range str[j+1:] {
					if k+1 > len(str[j+1:])-1 {
						break
					}
					if parseInt(string(str[j+1:][k])) > parseInt(string(str[j+1:][k+1])) {
						break check
					}
				}

				match++
				break
			}
		}
	}

	fmt.Printf("part1 solution: %v\n", match)
}

func parseInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
