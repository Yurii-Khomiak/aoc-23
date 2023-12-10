package aoclib

import (
	"math"
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

func PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func Contains[T comparable](s []T, v T) bool {
	for _, el := range s {
		if el == v {
			return true
		}
	}
	return false
}

func MakeChannels[T any](n int) []chan T {
	channels := make([]chan T, 0, n)
	for i := 0; i < cap(channels); i++ {
		channels = append(channels, make(chan T))
	}
	return channels
}

func splitIntoLines(s string) []string {
	res := strings.Split(s, "\n")
	if last := len(res) - 1; len(res[last]) == 0 {
		res = res[:last]
	}
	return res
}
