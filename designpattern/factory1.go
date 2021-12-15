package main
/*简单工厂模式*/
import "fmt"

//API是一个接口，有未实现的方法
type API interface {
	Say(name string) string
}
//go如果实现了接口的方法，就相当于实现了这个接口了。
//一个hiAPI结构体
type hiAPI struct{}
func (*hiAPI) Say(name string) string{
	return fmt.Sprintf("Hi,%s",name)
}
type helloAPI struct{}
func (*helloAPI) Say(name string) string{
	return fmt.Sprintf("Hello,%s",name)
}

//go没有构造函数一说，一般用NewXX代替构造函数，这其实就是简单工厂
func NewAPI(t int) API{
	if t==1{
		return &hiAPI{}
	}else if t==2{
		return &helloAPI{}
	}
	return nil
}

func main() {
	api:=NewAPI(2)
	println(api.Say("xiaoming"))
}


