package main

import "fmt"

//也就是说a()函数返回的就是一个函数，是一个闭包，如果不是显示的执行b 那么b就不会被执行。
func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}
func main() {
	c := a()
	c()
	c()
	c()
	a() //不会输出i
}