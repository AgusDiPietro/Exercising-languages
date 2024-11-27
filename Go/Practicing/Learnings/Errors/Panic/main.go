package main

import (
	"fmt"
)

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}

func riskyFunction() {
	defer handlePanic()
	panic("unexpected situation")
}

func main() {
	fmt.Println("before the risky function")
	riskyFunction()
	fmt.Println("after risky function")
}
