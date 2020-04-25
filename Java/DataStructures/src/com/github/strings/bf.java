package com.github.strings;

public class bf {
    public static int search(String haystack, String needle){
        int sLen = haystack.length();// 主字符串
        int pLen = needle.length();// 模式串长度
        // 需要匹配的次数
        for (int i=0;i<=sLen-pLen;i++){
            int j ;
            // 遍历模式串
            for (j=0;j<pLen;j++){
                if (needle.charAt(j)!=haystack.charAt(i+j)){
                    break;
                }
            }
            // 如果j移动到模板末尾了 说明匹配成功了
            if (j==pLen) return i ;
        }
        return -1;
    }
    public static void main(String[] args) {
        System.out.println(search("ababbabaa","aaab"));
        System.out.println(search("hello","ll"));
        System.out.println(search("aaaaa","bba"));
    }
}


