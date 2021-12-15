package main
import (
	"fmt"
)
type Programmer interface {
	writeHelloWorld() string;
}
type GoProgrammer struct{

}
//duck type 我们并没有显式的使用interface来使用接口
func (got *GoProgrammer) writeHelloWorld() string{
	return "hello world";
}
func testClient(){
	var p Programmer
	p=new(GoProgrammer)
	fmt.Println(p.writeHelloWorld)
}
func main() {
	testClient()
}