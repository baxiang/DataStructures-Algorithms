package main

import (
	"fmt"
	"sort"
)

func nextPermutation3(nums []int)  {
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

func nextPermutation(nums []int) {
	// 第一步从右往左查找最大的升序队列
	i := len(nums)-2// 倒数第二个数
	for i >= 0&&nums[i]>=nums[i+1] {
		i--
	}

	if i==-1 {
		// 从右往左一直是升序
		sort.Ints(nums)
	}else {
		j := len(nums)-1
		// 从右往左找第一个比i索引值大的数字索引坐标
		for j>=0&&nums[j]<=nums[i]{
			j--
		}

		//交换他两位置
		nums[i], nums[j] = nums[j], nums[i]
		// 重新排序
		sort.Ints(nums[i+1:])

	}
	fmt.Println(nums)
}




func main() {
	//nextPermutation([]int{1,2,3})
	//nextPermutation([]int{3,2,1})
	nextPermutation([]int{1,5,1})
	//nextPermutation([]int{1,1,5})
	//nextPermutation([]int{1,3,5,7,9})
}
