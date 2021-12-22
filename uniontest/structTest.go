package main
type pers struct{
	age int
	name string
}
func changeName(name string,p pers){
	p.name=name
}
func main() {
	man:=pers{age:10,name: "li",}
	println(man.name)
	//默认是值传递，无法改变原值
	changeName("chen",man)
	println(man.name)
	//但是在同一方法体内肯定是改变的
	man.name="chen"
	println(man.name)
}