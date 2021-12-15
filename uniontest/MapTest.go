package main

import "fmt"

func main() {
	//map[keyType]valueType
	// m:=map[string]int{"one":1,"two":2}
	m1:=map[string]int{}
	m1["one"]=1
	m2:=map[int]int{}
	m2[4]=16
	fmt.Print(m2)
	if v,ok:=m2[4];ok{
		fmt.Println("key 3 not exist",v)
	}else{
		fmt.Println("")
	}
	for key,value:=range m2{
		fmt.Println(key,value)
	}
}
