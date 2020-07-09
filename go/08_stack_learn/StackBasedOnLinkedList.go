package _8_stack

import "fmt"

/*
基于链表实现的栈
*/
type Node struct {
	next *Node
	val  interface{}
}
type LinkedListStack struct {
	head *Node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		head: nil,
	}
}
func (stack *LinkedListStack) Flush() {
	stack.head = nil
}
func (stack *LinkedListStack) IsEmpty() bool {
	return stack.head == nil
}

func (stack *LinkedListStack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.head.val
}
func (stack *LinkedListStack) Push(v interface{}) {
	newNode := &Node{
		next: stack.head,
		val:  v,
	}
	stack.head = newNode
}
func (stack *LinkedListStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	tmp := stack.head
	stack.head = stack.head.next
	return tmp.val
}
func (stack *LinkedListStack) Print() {
	cur := stack.head
	for nil != cur {
		fmt.Println(cur)
		cur = cur.next
	}
}
