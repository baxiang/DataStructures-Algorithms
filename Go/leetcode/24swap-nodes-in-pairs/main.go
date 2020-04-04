package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	  pre := &ListNode{}// 创建哨兵节点
	  pre.Next = head
	  temp := pre
      for temp.Next!=nil&&temp.Next.Next!=nil {
      	s :=temp.Next
      	e := temp.Next.Next
      	temp.Next = e
      	s.Next = e.Next
      	e.Next = s
      	temp = s
	  }
      return pre.Next
}
func swapPairs1(head *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	dummy.Next = head
	p := dummy
	for p.Next != nil && p.Next.Next != nil{
		first := p.Next
		second := first.Next
		//交换前两个节点
		first.Next = second.Next
		second.Next = first
		p.Next = second
		p = p.Next.Next
	}
	return dummy.Next
}

func tailNode(head *ListNode)*ListNode {
	curr := head
	var tail *ListNode
	for curr!=nil{
		tail = curr
		curr = curr.Next
	}
	fmt.Println(tail.Val)
	return tail
}
func swapPairs2(head *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := head
	tail := dummy
	for curr != nil && curr.Next != nil{
		t := &ListNode{
			Val:  curr.Val,
			Next: nil,
		}
		tail.Next = &ListNode{
			Val:  curr.Next.Val,
			Next: t,
		}
		tail = t
		curr = curr.Next.Next
	}
	if curr!= nil {
		tail.Next = &ListNode{
			Val:  curr.Val,
		}
	}
	return dummy.Next
}

func swapPairs3(head *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := head
	tail := dummy
	for curr != nil && curr.Next != nil{
		tmp := curr.Next.Next

		curr = tmp
	}
	tail.Next = curr
	return dummy.Next
}

func revertNode(root *ListNode) *ListNode{
	dummy := &ListNode{Val:  0, Next: nil}
	curr := root
	for curr!=nil{//遍历
		head := dummy.Next// 当前头节点
		newNote := &ListNode{
			Val:  curr.Val,
		}
		dummy.Next = newNote
		newNote.Next = head
		curr = curr.Next
	}
	return dummy.Next
}
//for cur != nil{
//next := cur.Next // 为了遍历，临时变量存储当前节点的下一个节点
//head :=  dummyNode.Next // 临时变量存储当前哑结点的指向的头结点
//dummyNode.Next = cur // 往头部插入数据,当前节点变成了头结点了
//cur.Next = head // 为了维持链条 需要把之前的头结点变成下一个节点
//cur = next // 遍历移动
//}
func revertNode1(root *ListNode) *ListNode{
	dummy := &ListNode{Val:  0, Next: nil}
	curr := root
	for curr!=nil{//遍历
		temp := curr.Next
		h := dummy.Next// 当前头节点
		dummy.Next = curr
		curr.Next = h
		curr = temp
	}
	return dummy.Next
}


func main() {
	a := &ListNode{Val:  1}
	b := &ListNode{Val:  2}
	c := &ListNode{Val:  3}
	d := &ListNode{Val:  4}
	e := &ListNode{Val:  5}
	f := &ListNode{Val:  6}
	a.Next=b
	b.Next =c
	c.Next = d
	d.Next = e
	e.Next = f
	h := swapPairs(a)
	for h!=nil{
		//fmt.Println(h.Val)
		h=h.Next
	}
}
