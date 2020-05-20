package main

import "fmt"

func main() {
	fmt.Println()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	ret = append(ret, []int{root.Val})
	ret1 := levelOrder(root.Left)
	ret2 := levelOrder(root.Right)
	if len(ret1) > len(ret2) {
		for i, v := range ret2 {
			ret1[i] = append(ret1[i], v...)
		}
		ret = append(ret, ret1...)
	} else {
		index := len(ret1)
		for i, v := range ret2 {
			if i == index {
				break
			}
			ret1[i] = append(ret1[i], v...)
		}
		ret1 = append(ret1, ret2[index:]...)
		ret = append(ret, ret1...)
	}
	return ret
}
