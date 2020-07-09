package _8_stack

import (
	"fmt"
)

/*
基于数组实现的栈
*/

type ArrayStack struct {
	data []interface{}
	top  int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}
func (stack *ArrayStack) IsEmpty() bool {
	if stack.top < 0 {
		return true
	}
	return false
}

func (stack *ArrayStack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.data[stack.top]
}

func (stack *ArrayStack) Flush() {
	stack.top = -1
}

func (stack *ArrayStack) Push(v interface{}) {

	stack.top++

	if stack.top > len(stack.data)-1 {
		stack.data = append(stack.data, v)
	} else {
		stack.data[stack.top] = v
	}
}
func (stack *ArrayStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	stack.top--
	return stack.data[stack.top+1]
}

func (stack *ArrayStack) Print() {
	if stack.IsEmpty() {
		fmt.Println("empty statck")
	} else {
		for i := stack.top; i >= 0; i-- {
			fmt.Println(stack.data[i])
		}
	}
}
