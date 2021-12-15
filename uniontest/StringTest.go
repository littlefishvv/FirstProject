package main

import (
	"fmt"
	"strings"
)

func main() {
	//在go语言中，string是只读的byte slice
	//string的byte数组可以存放任何数据
    var s string
    s="hello"
    //len是 byte数 一个汉字是3个byte 一个英文字母是一个byte
    fmt.Println(len(s))
    //utf8和unicode的关系 unicode是一种字符集也是一种编码，而utf8是这种编码的存储实现
    //rune数据类型

    c:=[]rune(s)
    fmt.Println(c)
    for _,t:=range s{
    	fmt.Print("%[1]c",t)
	}
	//常用方法
	//1 字符串分割
	str:="a,b,c"
	parts:=strings.Split(str,",")
	fmt.Println(parts)
	//字符串连接
	strs:=strings.Join(parts,",")
	fmt.Println(strs)
}
