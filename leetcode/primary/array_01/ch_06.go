package main

import "sort"

/**
两个数组的交集 II
给定两个数组，编写一个函数来计算它们的交集。


说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
我们可以不考虑输出结果的顺序。
进阶：

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2y0c2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */

func intersect(nums1 []int, nums2 []int) []int {
	//排序
	sort.Ints(nums1)
	sort.Ints(nums2)
	i,j := 0,0
	ret := []int{}

	for i< len(nums1) && j < len(nums2){
		if nums1[i] == nums2[j] {
			ret = append(ret, nums1[i])
			i ++
			j ++
		}else if nums1[i] < nums2[j] {
			i ++
		}else {
			j ++
		}

	}
	return ret
}

