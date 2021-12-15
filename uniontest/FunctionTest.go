package main

import (
	"fmt"
	"math/rand"
	"time"
)

//go语言的函数可以有多个返回值
//并且函数可以当作参数 或者返回值 或者map的参数 或者其他方法的参数 用inner func
func main() {
	//不想要哪个返回值可以用_来代替
   /* a,_:=returnMutiValues();
    fmt.Println(a)*/
	//testFn()
	deferTest()
}

func returnMutiValues() (int,int ){
	return rand.Intn(15),rand.Intn(20)
}
//参数是一个函数 返回值也是一个函数 得到的是参数这个函数的执行时间
func timeSpend(in func(op int) int) func (op int) int{
	return func(n int) int{
		start:=time.Now()
		ret:=in(n)
		fmt.Println("time spent:",time.Since(start).Seconds())
	    return ret
	}
}
func slowFunc(op  int) int{
	time.Sleep(time.Second*1)
	return op
}
func testFn(){
	//这个变量是一个新的函数
	tsF:=timeSpend(slowFunc)
	fmt.Println(tsF(10))
}
//可变长参数 其实会被转化为数组 通过数组遍历来完成
func sum(ops ...int) int{
	s:=0;
	for _,op:=range ops{
		s+=op
	}
	return s;
}
//延迟执行函数 defer 类似于finally
func deferTest(){
	defer func(){
        fmt.Println("释放资源")
	}()
	fmt.Println("延迟执行函数")
}