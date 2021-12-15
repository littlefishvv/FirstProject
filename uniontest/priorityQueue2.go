package main

import (
	"container/heap"
)

//自定义类型 []int切片类型
type iheap []int
//重写heap接口的方法
func (h iheap) Len() int{
	return len(h)
}
//建立小根堆
func (h iheap) Less(i,j int) bool {
	return h[i]<h[j]
}
//为什么swap不需要指针类型
func (h *iheap) Swap(i,j int){
	(*h)[i],(*h)[j]=(*h)[j],(*h)[i]
}
//这个不是堆里的方法 返回堆头
func (h iheap) Peek() int{
	return h[0]
}
//为什么要用指针类型
// 'Push' method has a pointer receiver 因为接收者是不确定的，所以必须要强制说明一下指针类型
/*
这个push也是一个钩子方法 这个钩子方法把数加到尾部 相当于堆底，所以加入之后要进行一个up操作。把刚加入到堆底的数上调到合适的位置。
func Push(h Interface, x interface{}) {
	h.Push(x)
	//把刚加入到堆底的数上调到合适的位置。
	up(h, h.Len()-1)
}
*/
func (h *iheap) Push(x interface{}){
	*h=append(*h,x.(int))
}
//为什么要去掉最后一个而不是第一个？因为这个pop是要被调用的相当于一个钩子方法。
/*
func Pop(h Interface) interface{} {
	n := h.Len() - 1
	//把堆顶调整到最后位置。
	h.Swap(0, n)
	//再给交换上来的堆顶找一个合适的位置，继续维护这个堆。经过了down会产生一个新的堆顶
	down(h, 0, n)
	//然后把最后那个位置给抛出去。
	return h.Pop()
}
*/
func (h *iheap) Pop() interface{}{
	n:=(*h).Len()
	x:=(*h)[n-1]
	*h=(*h)[:n-1]//把最后一个去掉
	return x
}

func findKthLargest(nums []int,k int) int{
	if k<=0 || len(nums)==0{
		return 0
	}
	h:=iheap{}

	heap.Init(&h)
	//构建k个容量的小根堆
	for i:=0;i<k;i++{
		heap.Push(&h,nums[i])
	}
	for i:=k;i<len(nums);i++{
		top:=h.Peek()
		if nums[i]>top{
			heap.Pop(&h)
			heap.Push(&h,nums[i])
		}
	}
	return h.Peek()
}
func main() {
	nums:=[]int{3,2,1,5,6,4}
	println(findKthLargest(nums,2))
}