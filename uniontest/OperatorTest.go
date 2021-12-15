package main

import (
	"fmt"
)

func main() {
	TestCompareArray()
}

func TestCompareArray(){
	a:=[...]int{1,2,3,4}
	b:=[...]int{1,2,3,4}
	fmt.Print(a==b)
}
