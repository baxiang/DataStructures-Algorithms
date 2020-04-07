package main

import "fmt"

func removeDuplicates(nums []int) int {
	 m := make(map[int]struct{})
	 index :=0
     for _,v:=range nums{
     	if _,ok:=m[v];!ok{
     		m[v] = struct{}{}
			nums[index]=v
     		index++
		 }
	 }
	 nums = nums[:index]
	 return len(m)
}
//快慢指针
func removeDuplicates1(nums []int) int {
	 slow := 0//慢指针
	 quick :=1//快指针变 块指针便利数组
	for ; quick<len(nums) ;quick++  {
		if nums[slow]!=nums[quick] {
			slow++// 只有当前值和快指针不相同
			nums[slow]=nums[quick]
		}
	}
	return slow +1
}


func main() {
	l := []int{1,1,2}
	r := removeDuplicates1(l)
	for i:=0;i<r;i++{
		fmt.Println(l[i])
	}
}
