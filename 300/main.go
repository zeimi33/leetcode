package main

import "fmt"

func main() {
	a := []int{4, 10, 4, 3, 8, 9}
	c := lengthOfLIS(a)
	fmt.Println(c)

}
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	a := []int{}
	a = append(a, nums[0])
	for _, g := range nums {
		if g > a[len(a)-1] {
			a = append(a, g)
		} else {
			l, r := 0, len(a)-1
			for r > l {

				mid := (l + r) / 2
				if a[mid] >= g {
					r = mid
				} else {
					l = mid + 1
				}
			}
			a[r] = g
		}
		fmt.Println(a)
	}

	return len(a)
}
