package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	lnums := len(nums)
	if lnums == 0 {
		return nil
	}
	next := [][]int{{nums[0]}}
	for i := 1; i < lnums; i++ {
		fmt.Println(next, "next1")
		new := next
		next = [][]int{}
		for _, ns := range new {
			con := ns
			fmt.Println(new, "ew")
			for j := 0; j <= i; j++ {
				fmt.Println(con[:j], con[j:])
				c := append([]int{}, con[:j]...)
				c = append(c, nums[i])
				c = append(c, con[j:]...)
				fmt.Println(c)
				next = append(next, c)
			}
		}
		fmt.Println(next, "next")
	}
	return next
}
