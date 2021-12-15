package main

import (
	"errors"
	"fmt"
)

func main() {
	if v,err:=fbnq(4);err!=nil {
		fmt.Println("出现了异常")
	}else{
		fmt.Println(v)
	}
}
func fbnq(n int) ([]int,error){
	if n<2||n>100 {
		return nil,errors.New("n should be in [2,100]")
	}
	fibList:=[]int{1,1}
	for i:=2;i<n;i++{
		fibList=append(fibList,fibList[i-2]+fibList[i-1])
	}
	return fibList,nil;
}
