package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals)==0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		 return intervals[i][0]<intervals[j][0]
	})
	queue :=[][]int{intervals[0]}
	for i:=1;i<len(intervals);i++{
		top :=queue[len(queue)-1]
		curr :=intervals[i]
		if curr[0]>top[1] {//当前值首地址大于栈顶的尾部元素
			queue = append(queue,curr)
			continue
		}
		//{a,b}{c,d} b>c
		if curr[1]>top[1] {
			top[1]=curr[1]// 合并区间
		}
	}
	return queue
}

func main() {
	fmt.Println(merge([][]int{{1,3},{2,4},{8,10},{15,18}}))
	fmt.Println(merge([][]int{{1,4},{4,5}}))
}
