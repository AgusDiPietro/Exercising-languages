package main

import (
	"fmt"
)

/* func main() {
	messages := make(chan string) //chan para el chanel y string para decirle el tipo de dato que va a contener
	go func() {
		messages <- "random message" // envia string al canal
	}()
	msg := <-messages
	fmt.Println(msg)
}*/

func main() {
	messages := make(chan string, 2) //el dos simboliza que podemos mandarle 2 valores al chanel
	messages <- "random1"
	messages <- "random2"
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
