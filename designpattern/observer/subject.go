package observer

import "fmt"

/*观察者模式*/
//主题subject
type Subject struct{
	observers []Observer
	context string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}
//往主题里添加观察者
func (s *Subject) Attach(o Observer){
	s.observers=append(s.observers,o)
}
//当context改变时通知观察者
func (s *Subject) notify(){
	for _,o:=range s.observers{
		o.Update(s)
	}
}

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