package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"

	"myprojects/aoc/23/aoclib"
)

type Rectangle struct {
	top   int
	bot   int
	left  int
	right int
}

type Gear struct {
	row int
	col int
}

type SchematicNumber struct {
	row   int
	left  int
	right int
}

func (n SchematicNumber) IsEqualTo(rhs SchematicNumber) bool {
	return n.row == rhs.row && n.left == rhs.left
}

func CalculatePart1(lines []string) int {
	res := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		for j := 0; j < len(line); j++ {
			if !unicode.IsDigit(rune(line[j])) {
				continue
			}

			num := getSchematicNumber(lines, i, j)
			j = num.right

			if isEnginePart(lines, num) {
				res += parseNumber(line[num.left : num.right+1])
			}
		}
	}
	return res
}

func CalculatePart2(lines []string) int {
	res := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if !isGear(lines[i][j]) {
				continue
			}

			surrounding := findSurroundingNumbers(lines, Gear{i, j})
			if len(surrounding) == 2 {
				res += multiplySchematicNumbers(lines, surrounding)
			}
		}
	}
	return res
}

func multiplySchematicNumbers(lines []string, numbers []SchematicNumber) int {
	product := 1
	for _, n := range numbers {
		line := lines[n.row]
		product *= parseNumber(line[n.left : n.right+1])
	}
	return product
}

func findSurroundingNumbers(lines []string, gear Gear) []SchematicNumber {
	res := make([]SchematicNumber, 0)

	r := getSurroundingRectangle(lines, gear.row, gear.col, gear.col)
	for i := r.top; i <= r.bot; i++ {
		for j := r.left; j <= r.right; j++ {
			if ch := lines[i][j]; !unicode.IsDigit(rune(ch)) {
				continue
			}

			num := getSchematicNumber(lines, i, j)
			if !isNumberPresent(res, num) {
				res = append(res, num)
			}
		}
	}

	return res
}

func isNumberPresent(nums []SchematicNumber, target SchematicNumber) bool {
	for _, n := range nums {
		if n.IsEqualTo(target) {
			return true
		}
	}
	return false
}

func getSchematicNumber(lines []string, i int, j int) SchematicNumber {
	line := lines[i]
	left := -1
	right := -1

	k := j - 1
	for ; k >= 0; k-- {
		if !unicode.IsDigit(rune(line[k])) {
			break
		}
	}
	left = k + 1

	k = j + 1
	for ; k < len(line); k++ {
		if !unicode.IsDigit(rune(line[k])) {
			break
		}
	}
	right = k - 1

	return SchematicNumber{row: i, left: left, right: right}
}

func isEnginePart(lines []string, num SchematicNumber) bool {
	r := getSurroundingRectangle(lines, num.row, num.left, num.right)
	for i := r.top; i <= r.bot; i++ {
		for j := r.left; j <= r.right; j++ {
			if isSymbol(lines[i][j]) {
				return true
			}
		}
	}
	return false
}

func getSurroundingRectangle(
	lines []string,
	row int,
	left int,
	right int,
) Rectangle {
	return Rectangle{
		top:   aoclib.Max(row-1, 0),
		bot:   aoclib.Min(row+1, len(lines)-1),
		left:  aoclib.Max(left-1, 0),
		right: aoclib.Min(right+1, len(lines[row])-1),
	}
}

func parseNumber(num string) int {
	parsed, _ := strconv.ParseInt(num, 10, 0)
	return int(parsed)
}

func isSymbol(ch byte) bool {
	return ch != '.' && !unicode.IsDigit(rune(ch))
}

func isGear(ch byte) bool {
	return ch == '*'
}

func main() {
	inputFilepath := os.Args[1]

	lines, err := aoclib.ReadFile(inputFilepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(CalculatePart1(lines))
	fmt.Println(CalculatePart2(lines))
}
