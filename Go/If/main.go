//if else

package main

import (
	"fmt"
	"os"
)

func main() {
	var rutaAbuscar string
	rutaAbuscar = "main.go"
	if _, err := os.Stat(rutaAbuscar); !os.IsNotExist(err) {
		fmt.Println("existe")
	} else {
		fmt.Println("no existe")
	}
}
