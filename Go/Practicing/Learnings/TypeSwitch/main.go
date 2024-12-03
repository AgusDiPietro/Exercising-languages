package main

import (
	"fmt"
)

func doSomething(i interface{}) {

	switch v := i.(type) {
	case int:
		fmt.Println("Int", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown")

	}
}
func main() {
	doSomething(10)
	doSomething("sfdssd")
	doSomething(3223)
}
