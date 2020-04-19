package main

func main() {

}

func canJump(nums []int) bool {
	l := len(nums)
	k := 0
	for i := 0; i < l; i++ {
		if i > k {
			return false
		}
		k = max(k, i+nums[i])
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
