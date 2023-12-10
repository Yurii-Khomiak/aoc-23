package main

import (
	"fmt"
	"os"
	"strings"

	"myprojects/aoc/23/aoclib"
)

func SolvePart1(games []string) chan int {
	ch := make(chan int)
	go func() {
		res := 0
		for _, game := range games {
			winning, ours := parseGame(game)
			n := numberOfMatches(winning, ours)
			res += aoclib.PowInt(2, n-1)
		}
		ch <- res
	}()
	return ch
}

func SolvePart2(games []string) chan int {
	ch := make(chan int)
	go func() {
		res := 0
		copies := make([]int, len(games))
		for i, game := range games {
			cardsCount := copies[i] + 1

			winning, ours := parseGame(game)
			n := numberOfMatches(winning, ours)

			end := aoclib.Min(i+n+1, len(games))
			for j := i + 1; j < end; j++ {
				copies[j] += cardsCount
			}

			res += cardsCount
		}
		ch <- res
	}()
	return ch
}

func parseGame(game string) ([]string, []string) {
	idx := strings.Index(game, ":")
	split := strings.Split(strings.Trim(game[idx+1:], " "), "|")

	winning := strings.Fields(strings.Trim(split[0], " "))
	ours := strings.Fields(strings.Trim(split[1], " "))

	return winning, ours
}

func numberOfMatches(winning, ours []string) int {
	n := 0
	for _, el := range ours {
		if aoclib.Contains(winning, el) {
			n++
		}
	}
	return n
}

func main() {
	inputFilepath := os.Args[1]

	lines, err := aoclib.ReadLines(inputFilepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	part1, part2 := SolvePart1(lines), SolvePart2(lines)

	fmt.Println(<-part1, <-part2)
}
