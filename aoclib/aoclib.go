package aoclib

import (
	"os"
	"strconv"
	"strings"
)

func ReadLines(filepath string) ([]string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return splitIntoLines(string(content)), nil
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func ParseInt(num string) int {
	parsed, _ := strconv.ParseInt(num, 10, 0)
	return int(parsed)
}

func splitIntoLines(s string) []string {
	res := strings.Split(s, "\n")
	if last := len(res) - 1; len(res[last]) == 0 {
		res = res[:last]
	}
	return res
}
