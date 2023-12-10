package main

import (
	"fmt"
	"os"
	"strings"

	"myprojects/aoc/23/aoclib"
)

type Range struct {
	dstStart int
	srcStart int
	length   int
}

type SeedRange struct {
	start  int
	length int
}

type Almanac struct {
	usages [][]Range
}

func (a *Almanac) Location(seed int) int {
	cur := seed
	for _, usage := range a.usages {
		for _, r := range usage {
			if cur >= r.srcStart && cur < r.srcStart+r.length {
				offset := (cur - r.srcStart)
				cur = r.dstStart + offset
				break
			}
		}
	}
	return cur
}

func SolvePart1(lines []string) int {
	blocks := splitByEmptyLines(lines)
	seeds, almanac := parseSeeds(blocks[0][0]), parseUsages(blocks[1:])

	min := int(^uint(0) >> 1)
	for _, seed := range seeds {
		if l := almanac.Location(seed); l < min {
			min = l
		}
	}

	return min
}

func SolvePart2(lines []string) int {
	blocks := splitByEmptyLines(lines)
	seeds, almanac := parseSeedRanges(blocks[0][0]), parseUsages(blocks[1:])

	minCh := make(chan int)

	for i := range seeds {
		go findMinLocationSeed(almanac, seeds[i:i+1], minCh)
	}

	min := int(^uint(0) >> 1)
	for range seeds {
		if cur := <-minCh; cur < min {
			min = cur
		}
	}

	return min
}

func findMinLocationSeed(almanac Almanac, seeds []SeedRange, minLocation chan int) {
	min := int(^uint(0) >> 1)
	for _, r := range seeds {
		start := r.start
		end := r.start + r.length

		for seed := start; seed < end; seed++ {
			if l := almanac.Location(seed); l < min {
				min = l
			}
		}
	}

	minLocation <- min
}

func splitByEmptyLines(lines []string) [][]string {
	res := make([][]string, 0, 8)
	lastIdx := 0
	for i := 0; i < len(lines); i++ {
		if isEmptyLine(lines[i]) {
			res = append(res, lines[lastIdx:i])
			lastIdx = i + 1
		}
	}
	res = append(res, lines[lastIdx:])
	return res
}

func parseSeeds(line string) []int {
	idx := strings.Index(line, ":")
	split := strings.Fields(strings.Trim(line[idx+1:], " "))

	seeds := make([]int, 0, len(split))
	for _, v := range split {
		seeds = append(seeds, aoclib.ParseInt(v))
	}

	return seeds
}

func parseSeedRanges(line string) []SeedRange {
	idx := strings.Index(line, ":")
	split := strings.Fields(strings.Trim(line[idx+1:], " "))

	seeds := make([]SeedRange, 0, len(split))
	for i := 0; i < len(split); i += 2 {
		seeds = append(seeds, SeedRange{
			aoclib.ParseInt(split[i]),
			aoclib.ParseInt(split[i+1]),
		})
	}

	return seeds
}

func parseUsages(usages [][]string) Almanac {
	ranges := make([][]Range, 0, len(usages))
	for _, usage := range usages {
		ranges = append(ranges, parseRanges(usage[1:]))
	}
	return Almanac{ranges}
}

func parseRanges(m []string) []Range {
	res := make([]Range, 0, len(m))
	for _, line := range m {
		split := strings.Fields(line)
		res = append(res, Range{
			dstStart: aoclib.ParseInt(split[0]),
			srcStart: aoclib.ParseInt(split[1]),
			length:   aoclib.ParseInt(split[2]),
		})
	}
	return res
}

func isEmptyLine(line string) bool {
	return len(strings.Trim(line, " ")) == 0
}

func main() {
	inputFilepath := os.Args[1]

	lines, err := aoclib.ReadLines(inputFilepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(SolvePart1(lines))
	fmt.Println(SolvePart2(lines))
}
