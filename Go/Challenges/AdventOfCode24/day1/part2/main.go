package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file containing the location IDs
	file, err := os.Open("../part1/locations.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Read the file and parse the left and right lists
	var leftList, rightList []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text()) //every line diveded in 2 numb
		if len(line) != 2 {
			fmt.Printf("Invalid line format: %s\n", scanner.Text())
			continue
		}
		leftNum, err1 := strconv.Atoi(line[0])
		rightNum, err2 := strconv.Atoi(line[1])
		if err1 != nil || err2 != nil {
			fmt.Printf("Error converting line: %s\n", scanner.Text())
			continue
		}
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Calculate frequency of numbers in the right list
	rightFreq := make(map[int]int)
	for _, num := range rightList {
		rightFreq[num]++
	}

	// Calculate the similarity score
	similarityScore := 0
	for _, num := range leftList {
		similarityScore += num * rightFreq[num]
	}

	// Print the result
	fmt.Printf("Similarity Score: %d\n", similarityScore)
}
