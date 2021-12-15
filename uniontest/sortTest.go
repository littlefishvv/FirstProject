package main

import (
	"fmt"
	"sort"
)

/*结构体排序*/
type Person struct {
	Name string // 姓名
	Age  int    // 年纪
}
/*
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
重写这三个方法
*/
// 按照 Person.Age 从大到小排序
//PersonSlice是一个自定义类型，本质上相当于person类型的切片。
type PersonSlice []Person

func (a PersonSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a PersonSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
//这个其实相当于实现了comparable接口
func (a PersonSlice) Less(i, j int) bool { // 重写 Less() 方法， 从小到大排序
	return a[i].Age < a[j].Age
}

func main() {
	people := PersonSlice{
		{"zhang san", 12},
		{"li si", 30},
		{"wang wu", 52},
		{"zhao liu", 26},
	}

	fmt.Println(people)
	//sort.Sort(PersonSlice(people)) 把person转换成personSlice类型。
	sort.Sort(people) // 按照 Age 的升序排序 传入一个重写了三个方法的类型。
	fmt.Println(people)
    //
	sort.Sort(sort.Reverse(people)) // 按照 Age 的降序排序
	fmt.Println(people)
}
