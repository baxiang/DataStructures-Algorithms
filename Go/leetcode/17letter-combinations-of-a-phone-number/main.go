package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	if digits == "" || len(digits) == 0 {
		return nil
	}
	nums :=[]string{"","","abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"}
	list :=[]string{""}// 默认为空字符
	for i:=0;i<len(digits);i++{
		s :=nums[digits[i]-'0']
		size :=len(list)
		for j:=0;j<size;j++{
			first := list[0]
			list = list[1:]
			for k:=0;k<len(s);k++{
				list = append(list,first+string(s[k]))
			}
		}
	}
	return list
}
func main() {
	fmt.Println(letterCombinations("23"))
}
