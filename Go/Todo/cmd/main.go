package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type task struct {
	name        string
	description string
	completed   bool
}

func main() {
	tasks := make(map[int]task)
	lastID := 0

	for {
		fmt.Println("Selecciona la opcion deseada:")
		fmt.Println("1. Ver tareas")
		fmt.Println("2. Agregar tarea")
		fmt.Println("3. Marcar tarea como terminada")
		fmt.Println("4. Eliminar tarea")
		fmt.Println("5. Salir")
		fmt.Print("Opcion: ")

		// scanner para obtener la entrada del SO
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		// convertir lo introducido en int
		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		case 1:
		case 2:
		case 3:
		case 4:
		case 5:
			fmt.Println("Saliendo...")
			os.Exit(0)

		}

	}

}
