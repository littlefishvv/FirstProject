package main

import "fmt"
//单向通道
//chan<- 放入
//<-chan 取出
func counter(out chan<-int){
	for i:=0;i<100;i++{
		//往out通道里放入i
		out<-i
	}
	close(out)
}
func squarer(out chan<-int,in <-chan int){
	for i:=range in{
		out<-i*i
	}
	close(out)
}
func printer(in <-chan int){
	for i:=range in{
		fmt.Println(i)
	}
}
func main() {
	ch1:=make(chan int)
	ch2:=make(chan int)
	//往ch1里放入1-100
	go counter(ch1)
	//从ch1中取出然后平方到ch2
	go squarer(ch2,ch1)
	//输出ch2的内容
	printer(ch2)
}