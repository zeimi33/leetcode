package main

func main() {

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var ans *TreeNode
	var search func(r *TreeNode) bool

	search = func(r *TreeNode) bool {
		if r == nil {
			return false
		}

		var self, left, right int

		if l := search(r.Left); l {
			left = 1
		}

		if r := search(r.Right); r {
			right = 1
		}

		if r == p || r == q {
			self = 1
		}

		if (left + self + right) >= 2 {
			ans = r
		}

		return (left + self + right) > 0
	}

	search(root)
	return ans
}
