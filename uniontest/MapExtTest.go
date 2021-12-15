package main
import "fmt"
func main() {
	//map的value可以是有返回值的函数。也就是通过这个函数来实现工厂模式
	/*m:=map[int]func(op int)int{}
	m[1]= func(op int)int {
		return op
	}
	m[2]= func(op int) int{
		return op*op
	}
	fmt.Print(m[1](2),m[2](2))*/

	//go的set是用map来实现的 map[type]bool
	//testMapForSet()
	var mapLit map[string]int
	//var mapCreated map[string]float32
	var mapAssigned map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	//对一个关联map的操作会影响原来的map
	mapAssigned["two"] = 3
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"]) // Map literal at "one" is: 1
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"]) // Map created at "key2" is: 3.14159
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"]) // Map assigned at "two" is: 3
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapAssigned["two"]) // Map literal at "ten" is: 0

}
func testMapForSet(){
	mySet:=map[int]bool{}
	mySet[1]=true
	if mySet[1] {
		fmt.Println("1 is existing")
	}else{
		fmt.Println("1 is not existing")
	}
	mySet[3]=true
	fmt.Println(len(mySet))
	delete(mySet,1)

}
