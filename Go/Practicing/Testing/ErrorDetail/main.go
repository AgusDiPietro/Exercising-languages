package main

import (
	"fmt"
)

type ErrorDetail struct {
	Arg  float64
	Prob string
}

func (e *ErrorDetail) Error() string {
	return fmt.Sprintf("%f - %s", e.Arg, e.Prob)
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, &ErrorDetail{x, "can't be negative"}
	}
	return x * x, nil
}

func main() {
	_, err := sqrt(-4)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
