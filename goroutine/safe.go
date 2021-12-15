package main

import (
	"fmt"
)


func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg1.Done()
}
func main() {
	wg1.Add(2)
	go add()
	go add()
	wg1.Wait()
	fmt.Println(x)
}
