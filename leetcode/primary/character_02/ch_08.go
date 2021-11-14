package main

import (
	"fmt"
	"strings"
)

/**
	编写一个函数来查找字符串数组中的最长公共前缀。

	如果不存在公共前缀，返回空字符串 ""。
 */

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	size := len(strs[0])
	baseStr := strs[0]
	for _, str := range strs{
		if len(str) < size {
			size = len(str)
			baseStr = str
		}
	}
	for len(baseStr) > 0 {
		for _, str := range strs{
			if !strings.Contains(str[:size],baseStr) {
				baseStr = baseStr[:size-1]
				break
			}
		}

		if size == len(baseStr) {
			return baseStr
		}
		size = len(baseStr)
	}

	return ""
}

func main() {
	pre := longestCommonPrefix2([]string{"flower","flow","flight"})
	fmt.Println(pre)
}

/**

横向扫描
 */
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	count := len(strs)
	prefix := strs[0]

	//挨个与prefix比较
	for i:=1; i < count ; i++  {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix

}

func lcp(str1, str2 string) string {
	//选出最短长度
	length := min(len(str1), len(str2))
	index := 0
	//每个字符作比较
	for index < length && str1[index] == str2[index] {
		index++
	}
	//截取得到前缀
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
