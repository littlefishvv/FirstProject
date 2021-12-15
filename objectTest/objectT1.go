package main

import   (
	"fmt"
	"unsafe"
)

//定义结构体
type Employee struct{
	Id string
	Name string
	Age int
}

func main() {
	e:=Employee{"0","bob",20}
	//e1:=Employee{Name:"bob",Age:20}
	e2:=new(Employee)//返回指针（地址）
	e2.Id="2"
	fmt.Println(e.getName())
}
//行为定义 func  方法名(参数名 参数类型) 返回值类型
func (e Employee) getName() string{
	fmt.Println(" ",unsafe.Pointer((&e.Name)))
	return e.Name
}
