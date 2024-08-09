package main

import (
	"fmt"
)

func main() {

	var hora int = 10
	switch hora {
	case 1, 2, 3, 4:
		fmt.Println("Es temprnao todavia")
	case 5, 6, 7:
		fmt.Println("Esta atardeciendo")
	case 8, 9:
		fmt.Println("Esta oscureciendo")
	case 10, 11:
		fmt.Println("Es de noche")
	default:
		fmt.Println("Madrugada")
	}
	// variable interfaz
	var x interface{}
	x = 14.5
	switch x.(type) {
	case nil:
		fmt.Println("Es una variable tipo nil")
	case int:
		fmt.Println("Es uan variable tipo int")
	case float64:
		fmt.Println("Es uan variable tipo float64")
	case string:
		fmt.Println("Es uan variable tipo string")
	default:
		fmt.Println("No es de niguno de los otros tipos")
	}

}
