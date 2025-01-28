package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Step 1: Read the input file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading the file: %v\n", err)
		return
	}

	// Step 2: Convert the file content into a grid (matrix) of runes
	grid := string(data)
	lines := strings.Split(grid, "\n")

	// Create a matrix of runes to represent the grid
	var matrix [][]rune
	for _, line := range lines {
		matrix = append(matrix, []rune(line))
	}

	// Step 3: Define the directions to search for the word "XMAS"
	// Directions include horizontal, vertical, diagonal, and reverse directions
	directions := [][]int{
		{1, 0},   // Right
		{0, 1},   // Down
		{1, 1},   // Diagonal right-down
		{1, -1},  // Diagonal right-up
		{-1, 0},  // Left
		{0, -1},  // Up
		{-1, -1}, // Diagonal left-up
		{-1, 1},  // Diagonal left-down
	}

	// Step 4: Define the word to search for
	word := "XMAS"
	wordRunes := []rune(word)
	count := 0 // Variable to count the number of occurrences

	// Step 5: Iterate through the matrix to search for the word
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			// Check if the current position matches the first letter of the word
			if matrix[i][j] == wordRunes[0] {
				// Search in all 8 directions
				for _, dir := range directions {
					dx, dy := dir[0], dir[1] // Direction offsets
					x, y := i, j             // Starting position
					match := true            // Assume a match until proven otherwise

					// Check the remaining letters of the word
					for k := 1; k < len(wordRunes); k++ {
						x += dx
						y += dy

						// Check if the new position is out of bounds or doesn't match the word
						if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[x]) || matrix[x][y] != wordRunes[k] {
							match = false
							break
						}
					}

					// If a match is found, increment the count
					if match {
						count++
					}
				}
			}
		}
	}

	// Step 6: Print the total number of occurrences of the word "XMAS"
	fmt.Printf("The word 'XMAS' appears %d times.\n", count)
}
