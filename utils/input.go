package utils

import (
	"os"
	"strings"
)

// ReadInput reads the input file for a specific day and returns lines as a slice of strings.
func ReadInput(filePath string) ([]string, error) {
    data, err := os.ReadFile(filePath)
    if err != nil {
        return nil, err
    }
    return strings.Split(strings.TrimSpace(string(data)), "\n"), nil
}

