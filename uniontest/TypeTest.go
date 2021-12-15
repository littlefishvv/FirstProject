package main
import (
	"fmt"
	"math"
)

func main(){
	typeTest()
 	var s string
	fmt.Println("*"+s+"*")
	fmt.Println(len(s))
}
func typeTest(){
	var a int32 = 1
	var b int64
	//不支持隐式类型转换
	a=int32(b)
	fmt.Print(a,b)
	t1:=math.MaxInt64
	t2:=math.MaxFloat32
	fmt.Print(t1," ",t2)
	fmt.Println()
}