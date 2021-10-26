package main

import (
	"fmt"
	"sort"
)

/**
有效字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn96us/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */

func isAnagram(s string, t string) bool {
	m := map[uint8]int{}
	n := map[uint8]int{}
	bytes := []byte(s)
	tbytes := []byte(t)

	if len(bytes) != len(tbytes){
		return false
	}

	for i := 0; i < len(bytes); i++ {
		m[bytes[i]]++
	}
	for j := 0; j < len(tbytes); j++ {
		n[tbytes[j]]++
	}

	for k,v:= range m  {
		if v != n[k] {
			return false
		}
	}

	return true
}

/**
	排序
 */

func isAnagram2(s string, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}
func main() {
	fmt.Println(isAnagram3("anagram","nagaraml"))
}

func isAnagram3(s string, t string) bool {
	// 本质，是字母数量一致
	// map，
	// 可以有3个优化点：
	// 1. 可以用数组下标作为key，因字符可以作为 int，也就可以作为下标
	// 2. 而对比数组，也可以用++，--来优化
	// 3. 小写字符，所以26的数组长度就OK
	cnt := make([]int, 26)
	for _, v := range s {
		cnt[v-'a']++
	}
	for _, v := range t {
		cnt[v-'a']--
	}
	for _, v := range cnt {
		if v != 0 {
			return false
		}
	}
	return true
}
