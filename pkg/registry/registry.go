package registry

import (
  "github.com/abalamilla/aoc-2024/pkg/core"
	"github.com/abalamilla/aoc-2024/solutions/day1"
	// Import other days
)

// GetSolution returns the solution for a given day.
func GetSolution(day int) core.Solution {
	switch day {
	case 1:
		return day1.Day1Solution{}
	default:
		return nil
	}
}
