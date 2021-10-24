package main

import (
	"fmt"
)

/**
旋转数组


给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
进阶：

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2skh7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
func rotate(nums []int, k int)  {
	k %= len(nums)
	ints := nums[len(nums)-k:]
	temp := nums[:len(nums)-k]
	for i:=0; i<len(temp) ;i++  {
		ints = append(ints, temp[i])
	}
	copy(nums,ints)
}

func main() {

	ints := []int{1, 2}
	rotate(ints,3)
	fmt.Println(ints)

}