package main

import "fmt"

/**
字符串中的第一个唯一字符
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
示例：
s = "leetcode"
返回 0
s = "loveleetcode"
返回 2
 */

/**
时间复杂度：O(n)O(n)，其中 nn 是字符串 ss 的长度。我们需要进行两次遍历。

空间复杂度：O(|\Sigma|)O(∣Σ∣)，其中 \SigmaΣ 是字符集，在本题中 ss 只包含小写字母，因此 |\Sigma| \leq 26∣Σ∣≤26。我们需要 O(|\Sigma|)O(∣Σ∣) 的空间存储哈希映射。

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string/solution/zi-fu-chuan-zhong-de-di-yi-ge-wei-yi-zi-x9rok/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
func firstUniqChar(s string) int {

	m := map[uint8]int{}
	bytes := []byte(s)

	for i := 0; i < len(bytes); i++ {
		m[bytes[i]]++
	}
	for i := 0; i < len(bytes); i++ {
		if m[bytes[i]] == 1 {
			return i
		}

	}
	return -1
}

func main() {
	char := firstUniqChar("aadadaad")
	fmt.Println(char)
}