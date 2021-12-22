package main
type Node struct{
	key int
	value int
	prev *Node
	next *Node
}
func initNode (key,value int) *Node{
	return &Node{
		key:key,
		value:value,
	}
}
type LRUCache struct {
	mapee map[int]*Node
	size int
	head *Node
	tail *Node
	capacity int
}

func Constructor(capacity int) LRUCache {
	cache:=LRUCache{
		mapee:map[int]*Node{},
		size:0,
		head:initNode(0,0),
		tail:initNode(0,0),
		capacity: capacity,
	}
	cache.head.next=cache.tail
	cache.tail.prev=cache.head
	return cache
}
func (this *LRUCache) Get(key int) int {
	mapee:=this.mapee
	if _, ok := mapee[key]; !ok {
		return -1
	}else{
		node:=mapee[key]
		this.moveToTail(node)
		return node.value
	}
}


func (this *LRUCache) Put(key int, value int)  {
	mapee:=this.mapee
	//用这种方式去取值更严谨   // 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	if _, ok := this.mapee[key]; ok {
		node := this.mapee[key]
		node.value = value
		this.moveToTail(node)
	}else{
		node:=&Node{
			key:key,
			value:value,
		}
		this.mapee[key]=node
		this.addToTail(node)
		this.size++
		if this.size>this.capacity{
			node=this.removeHead()
			delete(mapee,node.key)
			this.size--
		}
	}
}
func  (this *LRUCache) addToTail(node *Node){
	node.next=this.tail
	node.prev=this.tail.prev
	this.tail.prev.next=node
	this.tail.prev=node
}
func  (this *LRUCache)moveToTail(node *Node){
	this.removeNode(node);
	this.addToTail(node)
}
func   (this *LRUCache)removeNode(node *Node){
	node.prev.next = node.next
	node.next.prev = node.prev
}
func  (this *LRUCache)removeHead() *Node{
	node:=this.head.next
	this.removeNode(node)
	return node
}
func main() {
	obj := Constructor(1);
	obj.Put(5,6)
	obj.Put(6,7)
	obj.Get(5)
}