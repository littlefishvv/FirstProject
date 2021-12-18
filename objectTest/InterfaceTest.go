package main


type Common struct{
	name string
	age int
}
type Programmer interface {
	writeHelloWorld() string;
}
func (got *Common) printF() {
	println(got.name,"诞生于",got.age)
}
type GoProgrammer struct{
	info *Common
}
type JavaProgrammer struct{
    info *Common
}
//duck type 我们并没有显式的使用interface来使用接口
func (got *GoProgrammer) writeHelloWorld() string  {
	(*got).info.printF()
	return "fmt.println(hello world)"
}
func (got *JavaProgrammer) writeHelloWorld() string {
	(*got).info.printF()
	return "sout(hello world)"
}
func testClient(){
	var p Programmer
	p=&GoProgrammer{&Common{name:"Go",age:2002}}
	println(p.writeHelloWorld())
	p=&JavaProgrammer{&Common{name:"java",age:1992}}
	println(p.writeHelloWorld())
}
func main() {
	testClient()
}