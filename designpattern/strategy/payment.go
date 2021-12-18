package strategy

import "fmt"

//结构体payment包含了一些属性，也包好了支付策略，这个支付策略是一个接口类型。
type Payment struct{
	context *PaymentContext
	strategy PaymentStrategy
}
type PaymentContext struct {
	Name,CardID string
	Money int
}
type PaymentStrategy interface{
	Pay(ctx *PaymentContext)
}


func (p *Payment) Pay(){
	p.strategy.Pay(p.context)
}
type CashStrategy struct{}

func (*CashStrategy) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash", ctx.Money, ctx.Name)
}

type ZfbStrategy struct {}
func (*ZfbStrategy) Pay(ctx *PaymentContext){
	fmt.Printf("Pay $%d to %s by zfb", ctx.Money, ctx.Name)
}

func NewPayment(name, cardid string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardid,
			Money:  money,
		},
		strategy: strategy,
	}
}
//执行测试文件之前需要先将源文件进行编译.