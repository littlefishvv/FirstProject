package main

import "fmt"
//defer 是先进后出
//这个很自然,后面的语句会依赖前面的资源，因此如果先前面的资源先释放了，后面的语句就没法执行了。
//defer后面的语句在执行的时候，函数调用的参数会被保存起来，但是不执行。也就是复制了一份。
//但是并没有说struct这里的this指针如何处理，go语言并没有把这个明确写出来的this指针当作参数来看待。


/*    1. 关键字 defer 用于注册延迟调用。
2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
3. 多个defer语句，按先进后出的方式执行。
4. defer语句中的变量，在defer声明时就决定了。*/
func main() {
	var whatever [5]struct{}

	for i := range whatever {
		defer fmt.Println(i)
	}
}
