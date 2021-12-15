package main
import (
	"fmt"
	"testing"
)

func  TestFirst(t *testing.T)  {
	t.Log("my first try")
}
func TestFibList(){
	var (
		b int =1
		a int=1
	)

	for i:=0;i<5;i++{
		fmt.Print(" ",b);
		tmp:=a
		a=b;
		b=tmp+a;
	}
}
func main(){
	TestFibList()
}