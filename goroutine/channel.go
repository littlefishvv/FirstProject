package main

import (
	"fmt"
)

func main() {
	//runtime.GOMAXPROCS(2)
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		/*
		    1.对一个关闭的通道再发送值就会导致panic。
		    2.对一个关闭的通道进行接收会一直获取值直到通道为空。
		    3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
		    4.关闭一个已经关闭的通道会导致panic。
		*/
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}

