package observer

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
		//调用观察者自身的update方法。
		o.Update(s)
	}
}
//当改变上下文时，通知观察者队列。
func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}