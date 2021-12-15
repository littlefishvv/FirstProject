package main

import "fmt"

//方法和函数的区别在于 方法需要指定实体，而函数不需要
type User struct{
	Name string
	Email string
}
func (u User) Notify(name string) string{

	u=User{name,"go.web"}
	fmt.Printf("Name:%v,value:%v",u.Name,u.Email)
	return u.Name
}
func main() {
	u:=User{}
	u.Notify("gsy")
}