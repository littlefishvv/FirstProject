package main

import (
	"fmt"
	"math"
	"strings"
)

func test(fn func() int) int{
	return fn()
}
type FormatFunc func(s string,x,y int) string
func format(fn FormatFunc,s string,x,y int) string{
	return fn(s,x,y)
}
func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+1:]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
func main(){
	s1:=test(func() int {return 100})
	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)
	println(s1,s2)
	//匿名函数由一个不带函数名的函数声明和函数体组成。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
	getSqrt:=func(a float64) float64{
		return math.Sqrt(a)
	}
	fmt.Println(getSqrt(4))
}
