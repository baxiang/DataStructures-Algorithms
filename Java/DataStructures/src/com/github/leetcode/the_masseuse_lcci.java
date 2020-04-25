package com.github.leetcode;

public class the_masseuse_lcci {

    public static int massage(int[] nums) {
        if(nums == null || nums.length == 0)
            return 0;
        else if(nums.length == 1){
            return nums[0];
        }
        // res表示前i个预约的最优解，即最优的预约时长
        int[] res = new int[nums.length];
        res[0] = nums[0];
        res[1] = Math.max(nums[0],nums[1]);
        for (int i = 2; i < nums.length; i++) {
            // 只能间隔安排时间，1要么前一个工作，当前休息；要么前一个休息，当前工作。选两个方案最大的值。
            res[i] = Math.max(res[i-1], res[i-2]+nums[i]);
        }
        return res[nums.length-1];
    }

    public static void main(String[] args) {
      System.out.println( massage(new int[]{1,2,3,1}));
        System.out.println( massage(new int[]{2,7,9,3,1}));
        System.out.println(massage(new int[]{2,1,4,5,3,1,1,3}));
        System.out.println(massage(new int[]{2,1,1,2}));

    }
}
