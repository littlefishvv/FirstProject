package main

import (
	"container/list"
	"fmt"
)

func main() {
	//list可以存储任何类型的值
	l:=list.New()
	//第二种初始化方式：var l2 list.List
	l.PushBack(67)
	l.PushFront("first")
	l.PushFront("second")
	fmt.Println(l.Front().Value)
	l.Remove(l.Back())
	fmt.Println(l.Front().Value)
	fmt.Println(l.Back().Value)
	element:=l.PushBack("back")
	//insertAfter相当于在element加pushBack
	//insertBefore相当于在elemet加pushFront
	//比如 3 2 1 4 就是frOnt1 front2 front3 back4
	//再insertAfter就是 32145 如果是insertBefore就是32154
	l.InsertAfter("back2",element)
	//l.InsertBefore("back2",element)
	fmt.Println(l.Front().Value)
	fmt.Println(l.Back().Value)

	for i:=l.Front();i!=nil;i=i.Next(){
		fmt.Print(" ",i.Value)
	}
}
