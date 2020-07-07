package _6_linkedlist

import "fmt"

type ListNode struct {
	next  *ListNode
	value interface{}
}
type LinkedList struct {
	head   *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}
func (note *ListNode) GetNext() *ListNode {
	return note.next
}
func (note *ListNode) GetValue() interface{} {
	return note.value
}
func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

//在某个节点之后插入节点
func (list *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newNote := NewListNode(v)
	newNote.next = p.next
	p.next = newNote
	list.length++
	return true
}

//在某个节点之前插入节点
func (list *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || list.head == p {
		return false
	}
	cur := list.head.next
	tmp := list.head
	for p != cur {
		if nil == cur {
			return false
		}
		tmp = cur
		cur = cur.next
	}
	newNote := NewListNode(v)
	newNote.next = cur
	tmp.next = newNote
	list.length++
	return true
}

//在链表头部插入节点
func (list *LinkedList) InsertToHead(v interface{}) bool {
	return list.InsertAfter(list.head, v)
}

//在链表尾部插入节点
func (list *LinkedList) InsertToTail(v interface{}) bool {
	cur := list.head
	tmp := cur
	for nil != cur {
		tmp = cur
		cur = cur.next
	}
	newNode := NewListNode(v)
	tmp.next = newNode
	list.length++
	return true

}

//根据索引查找节点
func (list *LinkedList) FindByIndex(index uint) *ListNode {
	if index > list.length-1 {
		return nil
	}
	var i uint = 0
	cur := list.head
	for ; i <= index; i++ {
		cur = cur.next
	}

	return cur
}

//删除传入的节点
func (list *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}
	cur := list.head
	tmp := cur
	for cur != p  {
		tmp = cur
		cur = cur.next
	}
	if nil == cur{
		return false
	}
	tmp.next = cur.next
	list.length --
	return true
}

//打印链表
func (list *LinkedList) Print() {
	cur := list.head.next

	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

/*
删除倒数第N个节点
*/
func (list *LinkedList) DeleteBottomN(n int) {
	tmp := list.FindByIndex(list.length - uint(n))
	list.DeleteNode(tmp)
}
