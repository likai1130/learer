package main

import (
	"fmt"
	"learner/leetcode/primary/tree_04/common"
)

/**
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，



作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnd69e/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */


func maxDepth(root *common.TreeNode) int {
	var lh,rh int
	if root == nil {
		return 0
	}else {
		lh = maxDepth(root.Left)
		rh = maxDepth(root.Right)
	}
	if lh > rh{
		return 1+lh
	}
	return  1 + rh

}
func main() {
	left1 := &common.TreeNode{15,nil,nil}
	right1 := &common.TreeNode{7,nil,nil}
	left := &common.TreeNode{9,nil,nil}
	right := &common.TreeNode{20,left1,right1}

	root := common.TreeNode{3,left,right}

	depth := maxDepth(&root)
	fmt.Println(depth)
}