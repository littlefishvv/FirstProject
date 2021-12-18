package decorator
/*
装饰模式使用对象组合的方式动态改变或增加对象行为。
Go语言借助于匿名组合和非入侵式接口可以很方便实现装饰模式。
使用匿名组合，在装饰器中不必显式定义转调原对象方法。
*/
//一个componet接口，concreteComponent和装饰者都是这个类型的，只是装饰者通过组合接口的方式来实现调用被装饰者的方法
type Component interface {
	Calc() int
}

type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	return 0
}
//接口可以被组合在结构体里。装饰者是匿名组合，代理模式是有名组合
type MulDecorator struct {
	Component
	num int
}

func WarpMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

func (d *MulDecorator) Calc() int {
	//取出被装饰者的基础方法d.Component.Calc()然后自己再加上自己的东西。
	return d.Component.Calc() * d.num
}

type AddDecorator struct {
	Component
	num int
}

func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}