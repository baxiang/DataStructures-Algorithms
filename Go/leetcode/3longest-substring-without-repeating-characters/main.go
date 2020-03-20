package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	freq := make([]int,128)
	var res = 0
	start,end := 0,-1
	for start<len(s){
		if end+1<len(s)&&freq[s[end+1]] == 0{
			end++
			freq[s[end]]++
		}else{
			freq[s[start]]--
			start++
		}
		res = max(res,end-start+1)
	}
	return res
}

func max(i,j int)int{
	if i>j{
		return i
	}else{
		return j
	}
}

func lengthOfLongestSubstring1(s string) int{
	ans := 0
	m := map[byte]int{}
	end,start :=0,0
	for  ;end<len(s);end++{
		 c := s[end]
		 l,ok := m[c]
		if ok {
			start = max(l,start)
		}
		 ans = max(ans,end-start+1)
		 m[c]=end+1
	}
	return ans
}



func lengthOfLongestSubstring2(s string) int{
	m :=make(map[byte]int)
	max :=0
    start :=0
	for i,_:=range s{
		v,ok := m[s[i]]
		if ok {
			j := start
			for ;j<=v;j++{
				delete(m,s[j])
			}
			start =j
		}
		m[s[i]] = i
		if max<i-start+1 {
			max = i-start+1
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring2("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring2("bbbb"))
	fmt.Println(lengthOfLongestSubstring2("pwwkew"))
	fmt.Println(lengthOfLongestSubstring2("abba"))
	fmt.Println(lengthOfLongestSubstring2("bpfbhmipx"))
    fmt.Println(lengthOfLongestSubstring2("eeydgwdykpv"))

}
