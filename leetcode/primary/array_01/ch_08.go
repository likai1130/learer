package main

import "fmt"

/**
加一

给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
示例 3：

输入：digits = [0]
输出：[1]

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2cv1c/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

 */
func plusOne(digits []int) []int {
	len := len(digits)
	sums := make([]int,len+1 ,len+1)
	digits[len-1] = digits[len-1] + 1
	for i:=len-1; i >= 0 ; i-- {
		if digits[i] >= 10 {
			mo := digits[i] % 10
			sums[i+1] = mo
			sums[i] = digits[i] / 10

		}else {
			sums[i+1] +=  digits[i]
			if  sums[i+1] >= 10{
				sums[i] = sums[i+1] / 10
				mo := sums[i+1] % 10
				sums[i+1] = mo
			}
		}
	}

	if sums[0] == 0 {
		sums = sums[1:]
	}

	return sums
}

func main() {
	digits := []int{9,9}
	one := plusOne(digits)
	fmt.Println(one)
}