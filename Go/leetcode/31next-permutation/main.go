package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int)  {
	n := len(nums)
	var i = n-1
	for ;i > 0; i-- {
		if nums[i] > nums[i-1] {
			j := i
			for j < n && nums[j] > nums[i-1] {
				j++
			}
			nums[i-1], nums[j-1] = nums[j-1], nums[i-1]
			break
		}
	}
	sort.Ints(nums[i:])
}

func nextPermutation1(nums []int)  {
	i :=len(nums)-2
	len := len(nums)
	for ;i>=0;i--{
		if nums[i]<nums[i+1] {
			break
		}
	}
	if i!=-1 {
		for j :=len-1;j>i;j--{
			if nums[j]>nums[i] {
				tmp :=nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
				break
			}
		}
	}
	for j :=1;j<=(len-i-i)/2;j++{
		tmp :=nums[j+1]
		nums[j+1] = nums[len-j]
		nums[len-j] = tmp
	}

}

func nextPermutation2(nums []int) {
	// 倒着寻找一个降序索引
	index := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			// 实际是index += (i-1)+1
			index += i
			break
		}
	}
	fmt.Println(index)
	// 当前已经是最大序
	if index == -1 {
		sort.Ints(nums)
	} else {
		// 从nums[index+1:]中寻找一个大于它的最小值
		for i := len(nums) - 1; i > index; i-- {
			if nums[i] > nums[index] {
				// 交换 & 重排
				nums[i], nums[index] = nums[index], nums[i]
				sort.Ints(nums[index+1:])
			}
		}
	}
}





func main() {
	num := []int{1,2,3}
	nextPermutation(num)
	fmt.Println(num)
	num1 := []int{3,2,1}
	nextPermutation(num1)
	fmt.Println(num1)
	num2 := []int{1,1,5}
	nextPermutation(num2)
	fmt.Println(num2)

	num3 := []int{1,3,5,7,9}
	nextPermutation2(num3)
	fmt.Println(num3)

}
