package main

import (
	"fmt"
	"os"
	"strconv"

	"myprojects/aoc/23/aoclib"
)

func calculatePart1(lines []string) int {
	res := 0
	for _, line := range lines {
		digits := [2]byte{firstDigit(line), lastDigit(line)}
		convertedVal, _ := strconv.ParseInt(string(digits[:]), 10, 0)
		res += int(convertedVal)
	}
	return res
}

func firstDigit(line string) byte {
	digit := byte(0)
	for _, char := range []byte(line) {
		if isNumber(char) {
			digit = char
			break
		}
	}
	return digit
}

func lastDigit(line string) byte {
	digit := byte(0)
	for i := len(line) - 1; i >= 0; i-- {
		if char := byte(line[i]); isNumber(char) {
			digit = char
			break
		}
	}
	return digit
}

func isNumber(char byte) bool {
	return char >= '1' && char <= '9'
}

const MIN_WORD_NUMBER_LENGTH = 3

var wordNumbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func calculatePart2(lines []string) int {
	res := 0
	for _, line := range lines {
		res += getCalibrationValue(line)
	}
	return res
}

func getCalibrationValue(line string) int {
	chars := []byte(line)
	digits := [2]byte{firstDigit2(chars), lastDigit2(chars)}
	convertedVal, _ := strconv.ParseInt(string(digits[:]), 10, 0)
	return int(convertedVal)
}

func firstDigit2(chars []byte) byte {
	for i := 0; i < len(chars); i++ {
		if digit := digitAtIndex(chars, i); digit != 0 {
			return digit
		}
	}
	return '0'
}

func lastDigit2(chars []byte) byte {
	for i := len(chars) - 1; i >= 0; i-- {
		if digit := digitAtIndex(chars, i); digit != 0 {
			return digit
		}
	}
	return '0'
}

func digitAtIndex(chars []byte, i int) byte {
	if char := chars[i]; isNumber(char) {
		return char
	} else if v := wordNumberToDigit(chars[i:]); v != 0 {
		return v
	} else {
		return byte(0)
	}
}

func wordNumberToDigit(line []byte) byte {
	if len(line) < MIN_WORD_NUMBER_LENGTH {
		return byte(0)
	}

	res := byte(0)
	for i, word := range wordNumbers {
		if len(line) < len(word) {
			continue
		}
		if word == string(line[:len(word)]) {
			res = byte(i+1) + '0'
			break
		}
	}
	return res
}

func main() {
	inputFilepath := os.Args[1]

	lines, err := aoclib.ReadFile(inputFilepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(calculatePart1(lines))
	fmt.Println(calculatePart2(lines))
}
