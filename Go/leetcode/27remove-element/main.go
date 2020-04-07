package main

import "fmt"
//2n
func removeElement(nums []int, val int) int {
	var l []int
    for _,v :=range nums{
		if val!=v{
			l = append(l,v)
		}
	}
	for i,v:=range l{
		nums[i]=v
	}
	return len(l)
}
//双指针写法
func removeElement1(nums []int, val int) int {
	slow := 0
	quick :=0
	for ;quick<len(nums);quick++{
		if nums[quick]!=val {
			if quick>slow {
				nums[slow]=nums[quick]
			}
			slow++
		}
	}
	return slow
}
func removeElement2(nums []int, val int) int {
	var index = 0
	for _,v :=range nums{
		if val!=v{
		   nums[index]= v
		   index++
		}
	}
	return index
}
func main() {
	l1 := []int{3,2,2,3}
	r := removeElement2(l1,3)
	fmt.Println(l1,r)

	l2 :=[]int{0,1,2,2,3,0,4,2}
	r = removeElement2(l2,2)
	fmt.Println(l2,r)
	////l3 :=[]int{0,1,2,2,3,0,4,2}
	//r = removeElement2(l1,3)
	//fmt.Println(l1,r)
}
