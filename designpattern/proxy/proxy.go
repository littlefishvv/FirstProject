package proxy
//有名组合了代理对象。但是和匿名组合有什么区别呢
type Proxy struct {
	real RealSubject
}
//可以看到，realSubject真正的包含在了代理对象里，也就是说这个是在运行前确定的，而装饰者模式机器根本不知道我们要装饰哪一个对象
//装饰者模式的被装饰者是一个统用的接口类型，我们用传参的形式传进去被装饰对象。
func (p Proxy) Do() string {
	var res string

	// 在调用真实对象之前的工作，检查缓存，判断权限，实例化真实对象等。。
	res += "pre:"

	// 调用真实对象
	res += p.real.Do()

	// 调用之后的操作，如缓存结果，对结果进行处理等。。
	res += ":after"

	return res
}
