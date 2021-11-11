package main

import (
	"fmt"
)

/**
	两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

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

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2jrse/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
 //方法一
func twoSum(nums []int, target int) []int {
	ret := []int{}
	for i := 0; i < len(nums) - 1; i++  {
		for j := 1; j <= len(nums) - 1; j++  {
			if i != j && nums[i] + nums[j] == target{
				ret = append(ret,i,j)
				return ret
			}
		}
	}

	return ret
}

//方法二
func twoSum2(nums []int, target int) []int {
	indexMap := map[int]int{}
	indexMap[nums[0]] = 0
	for i := 1; i < len(nums) - 1; i++ {
		cha := target - nums[i]
		if index, ok  := indexMap[cha]; ok{
			return []int{i,index}
		}
		indexMap[nums[i]] = i
	}
	return []int{}
}


func main() {
	sum := twoSum2([]int{2,5,5,11}, 10)
	fmt.Println(sum)
}