package templatemethod2

import "fmt"
//因为父类需要调用子类的方法，所以子类需要匿名组合父类的同时，父类需要持有子类类型的引用。
type IPerson interface {
	SetName(name string)
	BeforeOut()
	Out()
}
//有名，这是组合。但是实现了整个接口 故可以看作是IPerson类型。
type Person struct {
	//因为要执行子类的方法，所以person类要持有子类类型的引用，而子类就是IPerson类型的。为什么是iperson类型？
	//因为子类匿名组合了父类，而父类实现了所有iPerson的接口是IPerson类型，所以子类也是Iperson类型。
	//这就相当于多态。
	Specific IPerson
	name     string
}


//可以看到person给这三个方法均提供了默认实现。
func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) BeforeOut() {
	if p.Specific == nil {
		return
	}

	p.Specific.BeforeOut()
}

func (p *Person) Out() {
	p.BeforeOut()
	fmt.Println(p.name + " go out...")
}

