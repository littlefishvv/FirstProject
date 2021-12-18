package singleton

import "sync"
/*
// if once.Do(f) is called multiple times, only the first call will invoke f,
// even if f has a different value in each invocation. A new instance of
// Once is required for each function to execute.
*/
type Singleton struct{}
//小写的变量
var singleton *Singleton
var once sync.Once
//提供一个共有的获取对象方法
func GetInstance() *Singleton{
	once.Do(func(){
		singleton=&Singleton{}
	})
	return singleton
}