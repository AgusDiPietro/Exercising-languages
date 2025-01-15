package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("levels.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var safeCount int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Read a line and convert it into a slice of int numb
		line := scanner.Text()
		levels := parseLineToIntSlice(line)

		// Check if the report is safe
		if isSafe(levels) {
			safeCount++
		}
	}

	// Handle potential scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Print the number of safe reports
	fmt.Printf("Number of safe reports: %d\n", safeCount)
}

// Convert a line of text into a list of numb
func parseLineToIntSlice(line string) []int {
	parts := strings.Fields(line) // Split the line into fields
	levels := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part) // Convert each field to an int
		if err != nil {
			fmt.Printf("Error parsing number: %v\n", err)
			continue
		}
		levels[i] = num
	}
	return levels
}

// Determine if a report is safe
func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return false // A report with fewer than 2 levels is not safe
	}

	// Determine if the levels are increasing or decreasing
	isIncreasing := levels[1] > levels[0]
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1] // Calculate the difference between adjacent levels

		// Check that the difference is between -3 and 3, and not zero
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		// Ensure the direction (increasing or decreasing) remains consistent
		if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
			return false
		}
	}

	return true
}
