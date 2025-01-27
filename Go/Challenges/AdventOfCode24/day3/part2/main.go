package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Read the input file using os.ReadFile
	data, err := os.ReadFile("../part1/input.txt")
	if err != nil {
		fmt.Printf("Error reading the file: %v\n", err)
		return
	}

	// Convert the file content to a string
	corruptedMemory := string(data)

	// Regular expression to find mul(X,Y), do(), and don't() instructions
	re := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`)

	// Find all matches in the corrupted memory
	matches := re.FindAllStringSubmatch(corruptedMemory, -1)

	// Variable to track if mul instructions are enabled
	enabled := true

	// Variable to accumulate the total sum
	totalSum := 0

	// Process each match
	for _, match := range matches {
		// Check the type of instruction
		switch match[1] {
		case "do()":
			// Enable mul instructions
			enabled = true
		case "don't()":
			// Disable mul instructions
			enabled = false
		default:
			// It's a mul(X,Y) instruction
			if enabled {
				// Extract X and Y
				x, err := strconv.Atoi(match[2])
				if err != nil {
					fmt.Printf("Error converting X: %v\n", err)
					return
				}

				y, err := strconv.Atoi(match[3])
				if err != nil {
					fmt.Printf("Error converting Y: %v\n", err)
					return
				}

				// Calculate the multiplication and add to the total
				totalSum += x * y
			}
		}
	}

	// Print the final result
	fmt.Printf("The total sum of the enabled multiplications is: %d\n", totalSum)
}
