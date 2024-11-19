package main

import (
	"fmt"
	"sync"
)

func printMessage(wg *sync.WaitGroup, message string) {
	fmt.Println(message)
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // añade la gorutine al waitgroup
	go printMessage(&wg, "Hi random")
	wg.Wait()
	fmt.Println("random")
}
