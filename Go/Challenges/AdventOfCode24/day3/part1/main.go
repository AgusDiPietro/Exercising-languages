package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func main() {
	// Read the input file
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading the file: %v\n", err)
		return
	}

	// Convert the file content to a string
	corruptedMemory := string(data)

	// Regular expression to find valid mul(X,Y) instructions
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// Find all matches in the corrupted memory
	matches := re.FindAllStringSubmatch(corruptedMemory, -1)

	// Variable to accumulate the total sum
	totalSum := 0

	// Process each match
	for _, match := range matches {
		// match[1] is X, match[2] is Y
		x, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Printf("Error converting X: %v\n", err)
			return
		}

		y, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Printf("Error converting Y: %v\n", err)
			return
		}

		// Calculate the multiplication and add to the total
		totalSum += x * y
	}

	// Print the final result
	fmt.Printf("The total sum of the multiplications is: %d\n", totalSum)
}
