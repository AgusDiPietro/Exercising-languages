package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hola, Go!")

	// Variables
	var tino string = "Edad 10"
	fmt.Println(tino)

	var mInt = 10
	fmt.Println(mInt)

	fmt.Println(tino, mInt)

	fmt.Println(reflect.TypeOf(tino))

	var mFloat = 5.5

	fmt.Println(mInt + int(mFloat))     // Conversión de float a int
	fmt.Println(mFloat + float64(mInt)) // Conversión de int a float

	// Crear una variable directamente
	variable := "Esto es una variable"
	fmt.Println(variable)

	// Constantes
	const mConst = "constante"
	fmt.Println(mConst)

	// Control de flujo
	if mInt == 5 && mInt == 10 { // Uso de operador de comparación
		fmt.Println("opcion 1")
	} else if mFloat == 6.5 || mInt == 11 { // Uso de operador de comparación
		fmt.Println("opcion 2")
	} else {
		fmt.Println("opcion 3")
	}

	// Arrays [posición]
	var mArray [4]int
	mArray[0] = 2
	mArray[3] = 5
	mArray[2] = 5
	fmt.Println(mArray)

	// Map
	myMap := make(map[string]int)
	myMap["Agus"] = 27
	myMap["Nay"] = 23
	myMap["pep"] = 22
	fmt.Println(myMap)

	// Otro mapa
	myMap2 := map[string]int{"Agus": 27, "Nay": 23, "tino": 10}
	fmt.Println(myMap2)

	// List
	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	myList.PushBack(5)
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Bucles
	for i := 0; i < len(mArray); i++ {
		fmt.Println(mArray[i])
	}

	for name, value := range myMap {
		fmt.Println(name, value)
	}

	fmt.Println(myFunction())

	// No hay clases, hay estructuras
	type MyStruct struct {
		name string
		age  int
		city string
	}

	myStruct := MyStruct{name: "Agus", age: 27, city: "Barcelona"}
	fmt.Println(myStruct)
}

// Función
func myFunction() string {
	return "mi función"
}
