package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
整数反转
	给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。
 

示例 1：

输入：x = 123
输出：321
示例 2：

输入：x = -123
输出：-321
示例 3：

输入：x = 120
输出：21
示例 4：

输入：x = 0
输出：0

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnx13t/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
/**
执行用时：
0 ms
, 在所有 Go 提交中击败了
100.00%
的用户
内存消耗：
2.2 MB
, 在所有 Go 提交中击败了
11.38%
的用户
 */
func reverse(x int) int {
	strs := []byte{}
	if x  < 0 {
		x = int(math.Abs(float64(x)))
		strs = append(strs, 45)
	}
	itoa := strconv.Itoa(x)
	b := []byte(itoa)
	for i := len(b)-1; i >= 0; i-- {
		strs = append(strs, b[i])
	}
	if strs[0] == 48 {
		strs = strs[1:]
	}
	atoi, _ := strconv.Atoi(string(strs))
	if !(-1<<31 <= atoi && atoi <= 1<<31 - 1){
		return 0
	}
	return atoi
}
/**
	取模
时间复杂度：O(\log |x|)O(log∣x∣)。翻转的次数即 xx 十进制的位数。

空间复杂度：O(1)O(1)。
 */
func reverse1(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}

func main() {
	atoi := reverse(1646324351)
	fmt.Println(atoi)


}