package main

import "fmt"

func testFor()  {
	n:=0;
	for n<5{
		n++;
		fmt.Print(n)
	}
}
func testIF(){
	if a:=1;a==1{
		fmt.Println(a)
	}
	for i:=0;i<5;i++{
		switch i {
		//case里可以有多个命中条件
		case 0,2:
			fmt.Println("even")
		case 1,3:
			fmt.Println("odd")

		default:
			fmt.Println("default")
		}
	}
}
func main() {
	testFor()
	testIF()
}
