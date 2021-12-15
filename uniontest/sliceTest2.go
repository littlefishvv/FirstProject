package main
func tetsSlice(a *[]int){
	(*a)[0]=4
}
func main() {
	//指针类型需要先初始化
	var a []int
    a=make([]int,8)
	a[0]=2
	a[3]=4
	tetsSlice(&a)
	println(a[0])
}
