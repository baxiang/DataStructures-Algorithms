package main

import "fmt"

type ListNode struct {
     Val int
     Next *ListNode
 }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
      //dump node
      d := &ListNode{Next:head}
      len :=0
      first := head
      for first!=nil{
      	first = first.Next
      	len++
	  }
      len =len-n// revert index
      first = d
      for len >0{
      	len--
      	first = first.Next
	  }
      first.Next = first.Next.Next
      return d.Next
}

func main() {
	head := &ListNode{Val:1}
	head.Next = &ListNode{Val:2}
	node := head.Next
	node.Next = &ListNode{Val:3}
	node = node.Next
	node.Next = &ListNode{Val:4}
	node = node.Next
	node.Next = &ListNode{Val:5}
	d := removeNthFromEnd(head,2)
	for  d!=nil {
		fmt.Println(d.Val)
		d = d.Next
	}
}
