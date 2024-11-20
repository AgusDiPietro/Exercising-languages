package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.RWMutex //lo usamos para bloquear acceso simultaneo a los recursos compartidos
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock() //bloqueamos y desbloqueamos la go rutine
	c.value++
}

func (c *Counter) Value() int {
	c.mu.RLock() //bloqueo solo para la lectura
	defer c.mu.RUnlock()
	return c.value
}

func main() {
	c := &Counter{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Actual value:", c.Value())
		}()

	}
	wg.Wait()
	fmt.Println("Final Counter:", c.Value())

}
