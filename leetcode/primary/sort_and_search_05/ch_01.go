package main

import (
	"fmt"
	"sort"
)

/**
合并两个有序数组

	给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。


作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnumcr/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */

func merge(nums1 []int, m int, nums2 []int, n int)  {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
	fmt.Println(nums1)


}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2,5,6}
	merge(nums1,3,nums2,3)
}


