package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Open the file with the locations
	file, err := os.Open("locations.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Read the file line by line
	var leftList, rightList []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line into two numbers
		line := strings.Fields(scanner.Text())
		if len(line) != 2 {
			fmt.Printf("Error: Invalid line format: %s\n", scanner.Text())
			continue
		}

		// Parse the numbers
		leftNum, err1 := strconv.Atoi(line[0]) //from text to number
		rightNum, err2 := strconv.Atoi(line[1])
		if err1 != nil || err2 != nil {
			fmt.Printf("Error: Invalid number in line: %s\n", scanner.Text())
			continue
		}

		// Append to respective lists
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	if len(leftList) == 0 || len(rightList) == 0 {
		fmt.Println("Error: One or both lists are empty. Unable to calculate distance.")
		return
	}

	// Sort both lists
	sort.Ints(leftList)
	sort.Ints(rightList)

	// Calculate the total distance
	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		totalDistance += abs(leftList[i] - rightList[i])
	}

	// Print the result
	fmt.Printf("Total Distance: %d\n", totalDistance)
}

// Helper function to calculate absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
