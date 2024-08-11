package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	var name string
	fmt.Printf("Enter your name:")
	fmt.Scan(&name)

	fmt.Printf("Hello %v, welcome to the Quiz game ", name)

	var age uint
	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)

	if age <= 10 {
		fmt.Println("Sorry you can't play this game. ")
	}

	fmt.Printf("Who is better, Messi or Ronaldo?: ")
	var answer string
	fmt.Scan(&answer)

	if answer == "Messi" {
		fmt.Println("Well done! you really know who is the GOAT")
	} else if answer == "Ronaldo" {
		fmt.Println("Sorry, you need to watch more football")
	} else {
		fmt.Println("Yo wrote it wrong, you lost your chance")
	}

	fmt.Printf("What show has a character named Cartman?: ")
	var answer2 string
	var answer3 string
	fmt.Scan(&answer2, &answer3)

	if answer2+""+answer3 == "South Park" {
		fmt.Println("Well done! i love that character")
	} else {
		fmt.Println("Sorry, the write answer was South Park. ")
	}

	fmt.Printf("What nation wons the World cup 2022? Argentina or France? : ")
	var answer4 string
	fmt.Scan(&answer4)

	if answer4 == "Argentina" {
		fmt.Println("Well done! ")
	} else if answer4 == "France" {
		fmt.Println("Sorry, they lost the final ")
	} else {
		fmt.Println("Yo wrote it wrong, you lost your chance")
	}
}
