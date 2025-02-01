package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error reading the file: %v\n", err)
		return
	}
	defer file.Close()

	var rules []string
	var updates [][]int
	readingRules := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readingRules = false
			continue
		}
		if readingRules {
			rules = append(rules, line)
		} else {
			nums := parseNumbers(line)
			updates = append(updates, nums)
		}
	}

	ruleMap := buildRuleMap(rules)
	sum := 0
	for _, update := range updates {
		if isValidUpdate(update, ruleMap) {
			mid := update[len(update)/2]
			sum += mid
		}
	}

	fmt.Println("Sum of middle pages:", sum)
}

func parseNumbers(line string) []int {
	split := strings.Split(line, ",")
	var nums []int
	for _, s := range split {
		n, _ := strconv.Atoi(s)
		nums = append(nums, n)
	}
	return nums
}

func buildRuleMap(rules []string) map[int]map[int]bool {
	ruleMap := make(map[int]map[int]bool)
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		if ruleMap[x] == nil {
			ruleMap[x] = make(map[int]bool)
		}
		ruleMap[x][y] = true
	}
	return ruleMap
}

func isValidUpdate(update []int, ruleMap map[int]map[int]bool) bool {
	position := make(map[int]int)
	for i, num := range update {
		position[num] = i
	}
	for x, ys := range ruleMap {
		for y := range ys {
			px, existsX := position[x]
			py, existsY := position[y]
			if existsX && existsY && px > py {
				return false
			}
		}
	}
	return true
}
