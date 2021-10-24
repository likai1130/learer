package main

import (
	"fmt"
	"sort"
)

/**
存在重复元素

给定一个整数数组，判断是否存在重复元素。

如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。

输入: [1,2,3,1]
输出: true


 */

 /**
 	暴力解法
 	时间复杂度O(n^2)
 	空间复杂度O(1)
  */
func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

/**
	哈希表
	时间复杂度O(n)
 	空间复杂度O(n)
 */
func containsDuplicate2(nums []int) bool {

	ints := map[int]int{}

	for i := 0; i < len(nums); i++ {
		if _,ok:= ints[nums[i]]; ok{
			return true
		}
		ints[nums[i]] = i

	}
	return false
}

/**
	排序
	时间复杂度：O(NlogN)，其中 NN 为数组的长度。需要对数组进行排序。

	空间复杂度：O(logN)，其中 NN 为数组的长度。注意我们在这里应当考虑递归调用栈的深度
 */
func containsDuplicate3(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}


func main() {
	ints := []int{1, 2, 3, 1}
	duplicate := containsDuplicate3(ints)
	fmt.Println(duplicate)
}