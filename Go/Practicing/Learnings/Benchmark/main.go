package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("it can't divide by 0")
	}
	return a / b, nil
}

func main() {
	result, err := divide(4, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result", result)
}