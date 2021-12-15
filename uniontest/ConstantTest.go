package main

import "fmt"

const(
	Monday=iota+1
	TuesDay
	WednestDay
)
func TestConstant(){
	fmt.Print(Monday)
	fmt.Print(TuesDay)
	fmt.Print(WednestDay)
}
func main(){
	TestConstant()
}