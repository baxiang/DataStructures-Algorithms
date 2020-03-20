package main

import "fmt"
//https://leetcode-cn.com/problems/two-sum/
func main() {
	n :=[]int{2,7,11,15}
	fmt.Println( twoSum(n,9))
}
func twoSum(nums []int, target int) []int {
	mapIndex := make(map[int]int)
	for i,v :=range nums{
		if v,ok :=mapIndex[target-v];ok {
			return []int{v,i}
		}
		mapIndex[v] = i
	}
	return nil
}
