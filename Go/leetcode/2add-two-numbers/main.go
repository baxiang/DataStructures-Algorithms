package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := l1
	q := l2
	currentHead := head
	var sum,carry int
	for true  {
		if q != nil || p != nil {
			var x,y int
			if p!=nil {
				x=p.Val
				p = p.Next
			}
			if q!=nil {
				y =q.Val
				q = q.Next
			}

			sum = carry + x + y
			carry = sum / 10
		} else {
			break
		}
		currentHead.Val = sum % 10
		if p!=nil||q!=nil {
			currentHead.Next = &ListNode{}
			currentHead = currentHead.Next
		}

	}
	if carry != 0 {
		currentHead.Next = &ListNode{carry, nil}
	}
	return head
}

func main() {
	l1:= &ListNode{Val:2,Next:&ListNode{Val:4,Next:&ListNode{Val:3}}}
	l2:= &ListNode{Val:5,Next:&ListNode{Val:6,Next:&ListNode{Val:4}}}
	r := addTwoNumbers(l1,l2)
	curr := r
	var str strings.Builder
	for curr!=nil{
		str.WriteString(fmt.Sprintf("%d -> ",curr.Val))
		curr = curr.Next
	}
	fmt.Println(strings.TrimRight(str.String()," -> "))
}