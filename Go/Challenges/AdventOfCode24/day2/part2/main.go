package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file containing the levels
	file, err := os.Open("../part1/levels.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	safeReports := 0

	// Read each line from the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Parse the line into a slice of integers
		line := scanner.Text()
		report := parseLineToIntSlice(line)

		// Check if the report is safe with the dampener
		if isSafeWithDampener(report) {
			safeReports++
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Print the total number of safe reports
	fmt.Printf("Number of safe reports: %d\n", safeReports)
}

// parseLineToIntSlice parses a line into a slice of integers
func parseLineToIntSlice(line string) []int {
	fields := strings.Fields(line)
	nums := make([]int, len(fields))
	for i, field := range fields {
		num, _ := strconv.Atoi(field) // Ignoring errors as the input is assumed valid
		nums[i] = num
	}
	return nums
}

// isSafeWithDampener checks if a report is safe, allowing one bad level to be ignored
func isSafeWithDampener(report []int) bool {
	// If the report is safe without modifications, return true
	if isSafe(report) {
		return true
	}

	// Try removing each level one by one and check if the modified report is safe
	for i := range report {
		modifiedReport := append(report[:i], report[i+1:]...) // Remove the i-th level
		if isSafe(modifiedReport) {
			return true
		}

		// Try removing one element before and after as well (i-1, i+1) for broader checking
		if i > 0 {
			modifiedReport = append(report[:i-1], report[i:]...) // Remove i-1
			if isSafe(modifiedReport) {
				return true
			}
		}
		if i < len(report)-1 {
			modifiedReport = append(report[:i], report[i+2:]...) // Remove i+1
			if isSafe(modifiedReport) {
				return true
			}
		}
	}

	// If no modification makes the report safe, return false
	return false
}

// isSafe checks if a report is safe based on the rules
func isSafe(report []int) bool {
	if len(report) < 2 {
		return true // Reports with fewer than 2 levels are always safe
	}

	// Determine if the report is increasing or decreasing
	increasing := report[1] > report[0]

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		// Check if the difference is valid
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		// Check if the direction is consistent
		if (increasing && diff < 0) || (!increasing && diff > 0) {
			return false
		}
	}
	return true
}
