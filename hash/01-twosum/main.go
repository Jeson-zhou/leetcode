package main

import (
	"fmt"
)

/*
URL:https://leetcode.cn/problems/two-sum/

给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

示例 2：
输入：nums = [3,2,4], target = 6
输出：[1,2]

示例 3：
输入：nums = [3,3], target = 6
输出：[0,1]
提示：

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案
进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？

Related Topics
数组
哈希表
*/

func main() {
	nums := [][]int{{2, 7, 11, 15}, {3, 2, 4}, {2, 3, 3, 4, 6}, {2, 5, 5, 11}}
	targets := []int{9, 6, 6, 10}
	for i, j := 0, 0; i < len(nums) && j < len(targets); i, j = i+1, j+1 {
		num := nums[i]
		target := targets[j]
		res := twoSum(num, target)
		fmt.Println(res)
	}
}

/*
[0 1]
[1 2]
[0 3]
[1 2]
*/

// twoSum 以空间换时间。O(n)的时间复杂度
func twoSum(nums []int, target int) []int {
	var result []int
	// 判空
	if len(nums) == 0 {
		return result
	}
	tmpMap := map[int]int{}
	for index, value := range nums {
		if i, ok := tmpMap[target-value]; ok {
			result = append(result, i, index)
			return result
		}
		tmpMap[value] = index
	}
	return result
}

// 双重循环，暴力求解
//func twoSum(nums []int, target int) []int {
//	if len(nums) == 0 {
//		return []int{}
//	}
//	// 3, 2, 4     6
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if target == (nums[i] + nums[j]) {
//				return []int{i, j}
//			}
//		}
//	}
//
//	return []int{}
//}
