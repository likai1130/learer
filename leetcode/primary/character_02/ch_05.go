package main

import (
	"fmt"
	"strings"
)

/**
验证回文串
	给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

 

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
解释："amanaplanacanalpanama" 是回文串
示例 2:

输入: "race a car"
输出: false
解释："raceacar" 不是回文串
 

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xne8id/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var sgood string
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			sgood += string(s[i])
		}
	}
	for i := 0; i < len(sgood)/2; i++ {
		if sgood[i] != sgood[len(sgood) -1 -i]{
			return false
		}
	}
	return true
}

/**
	双指针
 */
func isPalindrome2(s string) bool {
	var sgood string
	s = strings.ToLower(s)
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			sgood += string(s[i])
		}
	}

	left, right := 0, len(sgood) - 1
	for left < right {
		if left < right {
			if sgood[left] != sgood[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}


func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func main() {
	fmt.Println(isPalindrome2("A man, a plan, a canal: Panama"))
}