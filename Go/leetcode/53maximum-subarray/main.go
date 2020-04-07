package main

import "fmt"

func maxSubArray(nums []int) int {
	len :=len(nums)
	// 定义动态数组
	dp :=make([]int,len)
	// 最大结果是第一个数
	result := nums[0]
	// 第一个值
	dp[0]=nums[0]
	for i:=1;i<len;i++{
		dp[i] =Max(dp[i-1]+nums[i],nums[i])
		result = Max(dp[i],result)
	}
	return result
}

func Max(a,b int)int{
	if a>b {
		return a
	}
	return b
}

//53. 最大子序和 https://leetcode-cn.com/problems/maximum-subarray/
func main() {
	//fmt.Println(maxSubArray([]int{-2}))
	//fmt.Println(maxSubArray([]int{-2,1}))
	fmt.Println(maxSubArray([]int{-2,1,-3}))
	fmt.Println(maxSubArray([]int{2,1,-3,4,-1,2,1,-5,4}))
}
