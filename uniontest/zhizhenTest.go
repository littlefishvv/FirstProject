package main

import "fmt"

func main() {
/*	var t *int
	*t = 100//invalid memory address or nil pointer dereference*/
	var a int
	fmt.Println(&a)
	var p *int
	p = &a
	*p = 20
	fmt.Println(a)
}
