package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	var name string
	fmt.Printf("Enter your name:")
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the Quiz game!", name)

	var age uint
	fmt.Printf("Enter your age:")
	fmt.Scan(&age)

	if age <= 10 {
		fmt.Println("Sorry you can't play this game.")
	}

	fmt.Printf("Who is better, Messi or Ronaldo?")
	var answer string
	fmt.Scan(&answer)

}
