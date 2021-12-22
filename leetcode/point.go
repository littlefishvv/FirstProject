package main
type Pt struct{
	name string
}
func (p *Pt) toname(name string){
	//这里进行了隐式转换。p就是(*p)
	p.name=name
}
func main() {
	newpt:=&Pt{}
	newpt.toname("hello world")
	println(newpt.name)
}
