package lru

import "fmt"

type Node struct {
	Key int
	Value int
	Prev *Node
	Next *Node
}

type LRUCache struct {
	Capacity int
	cache map[int]*Node
	head *Node
	tail *Node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		Capacity: capacity,
		cache: make(map[int]*Node),
		head: &Node{},
		tail: &Node{},
	}

	l.head.Next = l.tail
	l.tail.Prev = l.head
	return l
} 

//辅助函数
func (this *LRUCache) addToHead(node *Node) {
	node.Prev = this.head
	node.Next = this.head.Next
	this.head.Next.Prev = node
	this.head.Next = node
}

func (this *LRUCache) removeNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *Node {
	node := this.tail.Prev
	this.removeNode(node)
	return node
}

//取
func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.Value
	}
	return -1
}

//写
func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.Value = value
		this.moveToHead(node)
	} else {
		newNode := &Node{Key: key, Value: value}
		this.cache[key] = newNode
		this.addToHead(newNode)

		if len(this.cache) > this.Capacity {
			removed := this.removeTail()
			delete(this.cache, removed.Key)
			fmt.Printf("容量已满，淘汰 Key: %d\n", removed.Key)
		}
	}
}

//调试用
func (this *LRUCache) printList() {
	curr := this.head.Next
	fmt.Print("当前缓存链表(Head -> Tail):")
	for curr != this.tail {
		fmt.Printf("[%d:%d] -> ", curr.Key, curr.Value)
		curr = curr.Next
	}
	fmt.Println("END")
}
