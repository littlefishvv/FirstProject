package main

import "fmt"

func main() {
	/*//切片 可变长数组
	var s0 []int;
	s0=append(s0,1)
	s1:=[]int{1,2,3,4}
	fmt.Print(len(s1))
    //创造一个切片2 长度是3 容量是5的切片 这个3就是我们可以访问的元素的个数
	s2:=make([]int,3,5)
	s2=append(s2,1)*/
	//切片是以不大于原容量的2^n为增长的。

	year:=[]string{"jan","feb","mar","Apr",
		"may","jun","jul","aug","sep","oct","nov","dec"}
	q2 := year[3:6]
	//长度是3 容量是9
	fmt.Println(q2,len(q2),cap(q2))
	summer:=year[5:8]
	fmt.Println(summer,len(summer),cap(summer))
	summer[0]="uknown"
	//共享同一块内存，影响是公共的。
	fmt.Println(q2)
	a:=[]int{1,2,3}
	//	slice = append(slice, elem1, elem2)
	//	slice = append(slice, anotherSlice...)
	//返回的是一个新的slice
	c:=append(a,[]int{0,1}...)
	fmt.Println(a,c)
	// 设置元素数量为1000
	const elementCount = 1000
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)
	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	// 引用切片数据
	refData := srcData
	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)
	// 将数据复制到新的切片空间中
	copy(copyData, srcData)
	// 修改原始数据的第一个元素
	srcData[0] = 999
	// 打印引用切片的第一个元素 可以看到引用切片也一起改变了。
	fmt.Println(refData[0])
	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1])
	// 复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6])
	//可以看到copyDate的0 1 被 4 5替换掉了
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
	//删除元素
	fmt.Println("...................删除元素......................")
	//Go语言中删除切片元素的本质是，以被删除元素为分界点，将前后两个部分的内存重新连接起来。
	/*连续容器的元素删除无论在任何语言中，都要将删除点前后的元素移动到新的位置，随着元素的增加，
	这个过程将会变得极为耗时，因此，当业务需要大量、频繁地从一个切片中删除元素时，如果对性能要求较高的话，
	就需要考虑更换其他的容器了（如双链表等能快速从删除点删除元素）。
	 */
	//删除开头元素 a=a[1:]  a = append(a[:0], a[1:]...)
	//a = append(a[:i], a[i+1:]...) 删除中间1个元素 删除第i个元素
	//a = a[:len(a)-1] 从尾部删除
}