package main

type AnimalAction interface{
	run()
	eat()
}
type Animal struct{
	//如果想调用子类的方法，就需要再集合一个action 比如在模板方法模式里就需要调用子类的方法。
	//action
	AnimalAction
	name string
	age int
}
func  (*Animal) eat(){
	println("默认实现")
}
func  (*Animal) run(){
	println("默认实现")
}
type Cat struct{
	*Animal
}

func  (c *Cat) run(){
	println(c.name+"cat自己的实现")
}
type Dog struct{
	*Animal
}

func  (d *Dog) run(){
	println(d.name+"dog自己的实现")
}
/*//这个表示是覆盖了原来的方法
func (*Cat) eat(){
	println("cat 重写过的方法。")
}
func (*Dog) eat(){
	println("dog 重写过的方法。")
}*/
func testDuitai(){
	//一个完美诠释多态的例子。
	//因为cat和dog匿名组合了animal，而animal实现了animalAction的接口
	//代表了全部都是animalAction类型的 并且因为匿名组合了animal 所以animal实现的接口方法，就是dog和cat
	//的默认方法。dog和cat后面实现的都是重写的方法。注意必须是匿名组合。
	var animal AnimalAction
	animal=&Cat{&Animal{name: "cat1"}}
	animal.eat()
	animal.run()
	animal=&Dog{&Animal{name: "dog1"}}

	animal.eat()
	animal.run()
}
func main() {
	testDuitai()
}