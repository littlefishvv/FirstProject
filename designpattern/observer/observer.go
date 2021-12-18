package observer

import "fmt"

//一个观察者接口，reader实现了观察者接口。并且持有主题对象以便接受到主题消息。
type Observer interface{
	Update(*Subject)
}
type Reader struct {
	name string
}
func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}
func (r *Reader) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}