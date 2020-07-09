package _9_queue

import "fmt"

type ListNode struct {
	val  interface{}
	next *ListNode
}

type LinkedListQueue struct {
	head   *ListNode
	tail   *ListNode
	length int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{nil, nil, 0}
}

func (this *LinkedListQueue) EnQueue(v interface{}) {
	node := &ListNode{v, nil}
	if this.tail == nil {
		this.tail = node
		this.head = node
	} else {
		this.tail.next = node
		this.tail = node
	}
}

func (this *LinkedListQueue) DeQueue() interface{} {
	if nil == this.head {
		return nil
	}
	tmp := this.head
	this.head = this.head.next
	return tmp
}

func (this *LinkedListQueue) String() string {
	if this.head == nil {
		return "empty queue"
	}
	result := "head<-"
	for cur := this.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.val)
	}
	result += "<-tail"
	return result
}
