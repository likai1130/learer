package main

import "fmt"

/**

只出现一次的数字

给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：
 */

 /**
 	循环双
 暴力
  */
func singleNumber(nums []int) int {
	if len(nums) < 2{
		return nums[0]
	}

	for i :=  0; i<len(nums);  i++ {
		for j := 0; j <  len(nums);  j ++ {
			if i == j && i != len(nums) -1  {
				continue
			}
			if nums[i] == nums[j] && i != len(nums) -1  {
				break
			}
			if j == len(nums)-1 {
				return nums[i]
			}

		}
	}
	return 0
}


/**
	异或
时间复杂度：O(n)O(n)，其中 nn 是数组长度。只需要对数组遍历一次。

空间复杂度：O(1)O(1)。
 */
func singleNumber1(nums []int) int {

	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func main() {
	ints := []int{2,2,1,4}
	duplicate := singleNumber1(ints)
	fmt.Println(duplicate)
}