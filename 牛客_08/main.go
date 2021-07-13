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
 * @return int整型二维数组
 */
func pathSum( root *TreeNode ,  sum int ) [][]int {
	// write code here
	ret := [][]int{}
	if root == nil {
		return ret
	}
	var f func(root *TreeNode,total int,num []int)
	f = func (r *TreeNode,total int,num []int){
		num = append(num,r.Val)
		sub := total - r.Val
		if r.Right == nil && r.Left == nil {
			if sub == 0 {
				tmp := make([]int,len(num))
				copy(tmp,num)
				ret = append(ret,tmp)
			}
		}
		if r.Left != nil {
			f(r.Left,sub,num)
		}
		if r.Right != nil {
			f(r.Right,sub,num)
		}
	}
	f(root,sum,[]int{})
	return ret
}

