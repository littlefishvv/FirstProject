package main

import (
	"fmt"
	"sync"

)

var lock sync.Mutex


func add1() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		X = X + 1
		lock.Unlock() // 解锁
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add1()
	go add1()
	wg.Wait()
	fmt.Println(x)


}
