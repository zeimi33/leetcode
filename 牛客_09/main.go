package main
//import . "nc_tools"

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}


/**
 *
 * @param root TreeNode类
 * @param sum int整型
 * @return bool布尔型
 */
func hasPathSum( root *TreeNode ,  sum int ) bool {
	// write code here
	if root == nil {
		return false
	}
	if root.Right == nil && root.Left == nil {
		return sum == root.Val
	}
	return hasPathSum(root.Left,sum-root.Val)||hasPathSum(root.Right,sum-root.Val)
}