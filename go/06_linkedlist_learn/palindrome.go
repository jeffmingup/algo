package _6_linkedlist

/*
思路1：开一个栈存放链表前半段
时间复杂度：O(N)
空间复杂度：O(N)
*/
func isPalindrome1(l *LinkedList) bool {
	len := l.length
	if len == 0 {
		return false
	}
	if 1 == len {
		return true
	}
	cur := l.head
	s := make([]string, 0, len/2)
	for i := uint(1); i < len; i++ {
		cur = cur.next
		//忽略奇数中间值
		if len%2 != 0 && i == len/2+1 {
			continue
		}

		if i <= len/2 {
			s = append(s, cur.GetValue().(string))
		} else {
			if s[len-i] != cur.GetValue().(string) {
				return false
			}
		}

	}

	return true
}

/*
思路2
找到链表中间节点，将前半部分转置，再从中间向左右遍历对比
时间复杂度：O(N)
*/
func isPalindrome2(l *LinkedList) bool {
	len := l.length
	if 0 == len {
		return false
	}
	if 1 == len {
		return true
	}

	step := len / 2
	var pre *ListNode = nil
	cur := l.head.next

	for i := uint(1); i <= step; i++ {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}

	mid := cur

	if len%2 != 0 {
		mid = mid.next
	}
	var left, right *ListNode = pre, mid
	isPalindrome := true
	for left != nil && right != nil {
		if left.GetValue() != right.GetValue() {
			isPalindrome = false
			break
		}
		left = left.next
		right = right.next
	}

	//复原链表
	//1 《-2 《-3 《-4  5  6 -》7 -》8

	for i := uint(1); i <= step; i++ {
		tmp := pre.next
		pre.next = cur
		cur = pre
		pre = tmp
	}
	l.head.next = pre

	return isPalindrome
}
