package main

import "learner/leetcode/primary/tree_04/common"

/**
对称二叉树

给定一个二叉树，检查它是否是镜像对称的。
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

  	1
   / \
  2   2
 / \ / \
3  4 4  3

*/

func isSymmetric(root *common.TreeNode) bool {
	if root == nil {
		return true
	}


	return isSymmetricHelper(root.Left,root.Right)
}

func isSymmetricHelper(leftNode *common.TreeNode, rightNode *common.TreeNode)  bool{
	if leftNode == nil && rightNode == nil {
		return true
	}

	if leftNode == nil || rightNode == nil || leftNode.Val != rightNode.Val {
		return false
	}

	return isSymmetricHelper(leftNode.Left,rightNode.Right) && isSymmetricHelper(leftNode.Right,rightNode.Left)
}
