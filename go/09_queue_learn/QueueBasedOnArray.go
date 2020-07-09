package _9_queue

import "fmt"

type ArrayQueue struct {
	arr      []interface{}
	capacity int
	head     int
	tail     int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 1}
}

func (this *ArrayQueue) EnQueue(v interface{}) bool {

	if this.tail == this.capacity {
		return false
	}
	this.arr[this.tail] = v
	this.tail++
	return true

}

func (this *ArrayQueue) DeQueue() interface{} {

	if this.head == this.tail {
		return nil
	}
	tmp := this.arr[this.head]
	this.head++
	return tmp
}

func (this *ArrayQueue) String() string {
	if this.head == this.tail {
		return "empty queue"
	}
	result := "head"
	for i := this.head; i <= this.tail-1; i++ {
		result += fmt.Sprintf("<-%+v", this.arr[i])
	}
	result += "<-tail"
	return result
}
