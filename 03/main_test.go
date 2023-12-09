package main

import "testing"

func TestPart1(t *testing.T) {
	inputLines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	expected := 4361

	if v := SolvePart1(inputLines); v != expected {
		t.Fatalf("Expected: %d\nGot: %d", expected, v)
	}
}

func TestPart2(t *testing.T) {
	inputLines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	expected := 467835

	if v := SolvePart2(inputLines); v != expected {
		t.Fatalf("Expected: %d\nGot: %d", expected, v)
	}
}
