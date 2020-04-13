package main
 type TreeNode struct {
	     Val int
	     Left *TreeNode
	    Right *TreeNode
}

func min(a,b int)int{
	if a>b {
		return b
	}
	return a
}
func minDepth(root *TreeNode) int {
     if root==nil{
     	return 0
	 }
	return min(minDepth(root.Right),minDepth(root.Left))
}

func main() {
	
}
