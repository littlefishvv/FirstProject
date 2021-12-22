package main
func tetsSlice(a *[]int){
	(*a)[3]=3
}
func main() {
	//指针类型需要先初始化
	var a []int
    a=make([]int,8)
	a[0]=2
	a[3]=4
	a[5]=5
	temp1:=a[:4]
	for _,v:=range temp1{
		print(v)
	}
	//tetsSlice(&a)
	a[3]=3
	for _,v:=range temp1{
		print(v)
	}
	println(a[3])
}
