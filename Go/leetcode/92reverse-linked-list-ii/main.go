package main

import "fmt"

type ListNode struct {
	     Val int
	     Next *ListNode
}

func reverseBetween1(head *ListNode, m int, n int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	begin := dummyNode
	for i:=0;i<m-1;i++{
		begin = begin.Next
	}
	var pre *ListNode
	curr := begin.Next
	for i:=0;i<n-m+1;i++{
		tmp := curr.Next
		curr.Next = pre
		pre = curr
		curr = tmp
	}
	begin.Next.Next = curr// 链接后节点
	begin.Next = pre // 开始反转的最后节点
	return dummyNode.Next
}


func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummyNode := &ListNode{}
	tail := dummyNode
	currNode := head
	index := 1
	for currNode!=nil{
		if index<m {
			tail.Next = currNode
			tail = tail
		}
		currNode = currNode.Next
	}


	return dummyNode.Next
}

func makeLinkedList(list []int)*ListNode{
	dummyNode := &ListNode{} // 创建哑结点 指向头节点
	curr := dummyNode
	for i:=0;i<len(list);i++{
		curr.Next= &ListNode{Val:list[i]}
		curr = curr.Next
	}
	return dummyNode.Next
}
func main() {
     head := makeLinkedList([]int{1,2,3,4,5,7,8})
     head = reverseBetween(head,2,4)
     for head!=nil{
     	fmt.Println(head.Val)
     	head = head.Next
	 }
}
