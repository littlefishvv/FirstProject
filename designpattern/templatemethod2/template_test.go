package templatemethod2

import "fmt"

func ExampleTpm() {
	person := &Person{}
	//模板方法模式到最后的乐星还是person 调用的也是person的方法，而不是子类重新实现的方法。
	person.Specific = &Boy{}
	person.SetName("zhangshan")
	person.Out()
	fmt.Println("----------------------")
	person.Specific = &Girl{}
	person.SetName("lisi")
	person.Out()
	// Output:
	// Pay $123 to Ada by cash
}