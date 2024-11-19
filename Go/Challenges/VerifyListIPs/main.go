package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Verifica si una IP es privada
func isPrivateIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsPrivate() {
		return true
	}
	return false
}

func main() {
	// Abrir el archivo de texto con la lista
	file, err := os.Open("List.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Escsner para leer el archivo linea por linea
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		ip := net.ParseIP(line)
		if ip == nil {
			fmt.Printf("%s: IP invalida \n", line)
			continue
		}

		// Si es una dirección IPv4 valida
		if ip.To4() != nil {
			if isPrivateIP(ip) {
				fmt.Printf("%s: IP privada\n", line)
			} else {
				fmt.Printf("%s: IP publica\n", line)
			}
		} else {
			fmt.Printf("%s: IP invalida o dato no es IP\n", line)
		}
	}

	// Manejar errores de escaneo
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}
}
