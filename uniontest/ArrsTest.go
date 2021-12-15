package main

import "fmt"

func testArray(){
	var a [3]int;
	a[0]=1;
	//b:=[3]int{1,2,3}
	//c:=[2][2]int{{1,2},{3,4}}
	//自动初始化到元素个数的长度
	//d:=[...]{1,2,3,4,5}


}
func main() {
	d:=[...]int{1,2,3,4,5}
	for i:=0;i<len(d);i++{
		fmt.Print(d[i])
	}
    //对应java中的for each
    //index是下标 e是值 通过range关键字输出值
	for idx,e:=range d{
		fmt.Println(idx,e)
	}
	//如果不想输出idx可以用_来占位
	for _,e:=range d{
		fmt.Println(e)
	}
	//数组截取 按下标前闭后开
	fmt.Println(
		d[1:2],d[1:3],d[1:],d[:3])



}