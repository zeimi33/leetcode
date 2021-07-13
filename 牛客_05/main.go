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
 * @return int整型
 */
func sumNumbers( root *TreeNode ) int {
	var ret []int
	var a func (r *TreeNode,num int)
	a = func (r *TreeNode,num int){
		if r == nil {
			return
		}
		if r.Right == nil && r.Left == nil {
			ret = append(ret,num*10+r.Val)
		}
		n := num *10 + r.Val
		a(r.Left,n)
		a(r.Right,n)
	}
	a(root,0)
	var num int
	for _,x := range ret{
		num = num + x
	}
	return num
}

