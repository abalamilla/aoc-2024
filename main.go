package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/abalamilla/aoc-2024/pkg/registry"
	"github.com/abalamilla/aoc-2024/utils"
)

func main() {
	// Parse the day parameter
	dayPtr := flag.Int("day", 0, "Specify the day of the Advent of Code challenge")
	flag.Parse()

	day := *dayPtr
	if day <= 0 {
		log.Fatalf("Invalid or missing day. Use the '-day' flag to specify a valid day.")
	}

	inputPath := fmt.Sprintf("input/day%d.txt", day)
	lines, err := utils.ReadInput(inputPath)
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	solution := registry.GetSolution(day)
	if solution == nil {
		log.Fatalf("No solution implemented for day %d", day)
	}

	fmt.Printf("Day %d, Part 1: %s\n", day, solution.SolvePart1(lines))
	fmt.Printf("Day %d, Part 2: %s\n", day, solution.SolvePart2(lines))
}
