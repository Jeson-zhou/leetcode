package main

import (
	"fmt"
	"sort"
)

/*
URL:https://leetcode.cn/problems/group-anagrams/

49. 字母异位词分组
字母异位词：两个单词如果包含相同的字母，次序不同，则称为字母易位词(anagram)。例如，“silent”和“listen”是字母
易位词，而“apple”和“aplee”不是易位词。

给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

示例 1:
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

示例 2:
输入: strs = [""]
输出: [[""]]

示例 3:
输入: strs = ["a"]
输出: [["a"]]


提示：
1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母
*/

func main() {
	strs := [][]string{{"eat", "tea", "tan", "ate", "nat", "bat"},
		{""},
		{"a"},
		{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"},
	}

	for _, str := range strs {
		res := groupAnagrams(str)
		fmt.Println(res)
	}
}

// groupAnagrams 思路如下：
// 1、创建一个map[string][]string 的map结构。
// 2、对 strs 中每一个string 进行排序。如果是 异位词 的两个单词，再排序之后，肯定是一样的。我们就以排序后的string为key，而value就是未排序前的string
func groupAnagrams(strs []string) [][]string {
	// 判空
	if len(strs) == 0 {
		return [][]string{{""}}
	}
	tmpMap := map[string][]string{}

	// for 循环遍历每一个子str,然后对子str进行排序
	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		// 以排序后的string为key，排序前的string为value
		if _, ok := tmpMap[string(bytes)]; ok {
			tmpMap[string(bytes)] = append(tmpMap[string(bytes)], str)
		} else {
			tmpMap[string(bytes)] = []string{str}
		}
	}

	// 遍历map，将结果保存在res中
	var res [][]string
	for v, _ := range tmpMap {
		//fmt.Println(v)
		res = append(res, tmpMap[v])
	}

	return res
}
