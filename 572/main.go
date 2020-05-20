package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := TreeNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if t == nil {
		return true
	}
	if s == nil {
		return false
	}
	if s.Val == t.Val {
		if readTree(s) == readTree(t) {
			return true
		}
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func readTree(t *TreeNode) string {
	if t == nil {
		return "_"
	}
	l := readTree(t.Left)
	r := readTree(t.Right)
	return fmt.Sprintf("%s%s%s", l, strconv.Itoa(t.Val), r)
}

//可以把readTree优化成只执行一次 不过没有什么必要
