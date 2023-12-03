package main

import "testing"

func TestPart1(t *testing.T) {
	inputLines := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	expected := 142

	if v := CalculatePart1(inputLines); v != expected {
		t.Fatalf("Expected: %d\nGot: %d", expected, v)
	}
}

func TestPart2(t *testing.T) {
	inputLines := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	expected := 281

	if v := CalculatePart2(inputLines); v != expected {
		t.Fatalf("Expected: %d\nGot: %d", expected, v)
	}
}
