package main
func main(){

}

  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

func maxDepth(root *TreeNode) int {


	a := 0
	func_return_depth(root,1,&a)
	return a
}

func func_return_depth(t* TreeNode,level int,max *int){
	if t.Left == nil && t.Right == nil {
		if level > *max {
			*max = level
		}
	}
	if t.Left != nil {
		func_return_depth(t.Left,level+1,max)
	}
	if t.Right != nil {
		func_return_depth(t.Right,level+1,max)
	}
}