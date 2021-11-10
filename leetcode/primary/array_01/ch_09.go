package main

import "fmt"

/**
移动零

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

1. 必须在原数组上操作，不能拷贝额外的数组。
2. 尽量减少操作次数。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2ba4i/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
func moveZeroes(nums []int)  {
	left, right, n := 0, 0, len(nums)

	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right],nums[left]
			left++
		}
		right ++
	}
	fmt.Println(nums)
}

func main() {
	//moveZeroes([]int{0,1,0,3,12})
	sums := []int{3, 2, 2, 3}
	element := removeElement(sums, 3)
	fmt.Println(element)
	for i := 0; i < element; i++ {
		fmt.Print(sums[i])
	}
}

/**
	拓展
	给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
 */
func removeElement(nums []int, val int) int {
	j:=0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val{
			nums[j] = nums[i]

			j ++
		}
		nums[i] = 0
	}
	return j +1

}

