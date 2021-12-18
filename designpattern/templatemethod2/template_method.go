package templatemethod2

import "fmt"

type Boy struct {
	Person
}

//重写了父类person的方法
func (b *Boy) BeforeOut() {
	fmt.Println("get up...")
}

type Girl struct {
	Person
}
//重写了父类person的方法。
func (g *Girl) BeforeOut() {
	fmt.Println("get up and dress up...")
}

