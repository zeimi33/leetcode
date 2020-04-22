package main

func main(){

}


type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
 
func rightSideView(root *TreeNode) []int {
	
	if root == nil{
		return nil
	}
	ret := []int{root.Val}
	if root.Left == nil &&root.Right ==nil{
		return []int{root.Val}
	}
	left := rightSideView(root.Left)
	right := rightSideView(root.Right)
	if len(right) >=len(left){
		return append(ret,right...)
	}

	return append(ret,append(right,left[len(right)]...))
}