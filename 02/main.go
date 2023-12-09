package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"myprojects/aoc/23/aoclib"
)

var maxCubeCounts = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func CalculatePart1(lines []string) int {
	res := 0
	for i, line := range lines {
		gameId := i + 1
		if isPossibleGame(line) {
			res += gameId
		}
	}
	return res
}

func isPossibleGame(line string) bool {
	for color, count := range parseMaxColorCounts(line) {
		if count > maxCubeCounts[color] {
			return false
		}
	}
	return true
}

func CalculatePart2(lines []string) int {
	res := 0
	for _, line := range lines {
		power := 1
		for _, count := range parseMaxColorCounts(line) {
			power *= count
		}
		res += power
	}
	return res
}

func parseMaxColorCounts(line string) map[string]int {
	res := map[string]int{}

	setStartIndex := strings.IndexAny(line, ":") + 2
	sets := strings.Split(line[setStartIndex:], ";")

	for _, set := range sets {
		cubes := strings.Split(strings.Trim(set, " "), ",")

		for _, cube := range cubes {
			parsed := strings.Fields(strings.Trim(cube, " "))

			count, _ := strconv.ParseInt(parsed[0], 10, 0)
			color := parsed[1]

			if v, ok := res[color]; !ok || v < int(count) {
				res[color] = int(count)
			}
		}
	}

	return res
}

func main() {
	inputFilepath := os.Args[1]

	lines, err := aoclib.ReadLines(inputFilepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(CalculatePart1(lines))
	fmt.Println(CalculatePart2(lines))
}
