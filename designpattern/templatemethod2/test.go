package templatemethod2
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

func  (*Cat) run(){
	println("cat自己的实现")
}
type Dog struct{
	 *Animal
}

func  (*Dog) run(){
	println("dog自己的实现")
}
/*//这个表示是覆盖了原来的方法
func (*Cat) eat(){
	println("cat 重写过的方法。")
}
func (*Dog) eat(){
	println("dog 重写过的方法。")
}*/
func test(){
	var animal AnimalAction
	animal=&Cat{&Animal{name: "cat"}}
	animal.eat()
	animal=&Dog{&Animal{name: "dog"}}
	animal.eat()
}