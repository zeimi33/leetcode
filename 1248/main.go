package main

import "fmt"

func main() {
	fmt.Println(numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2))
}
func numberOfSubarrays(nums []int, k int) int {
	index := 0
	i := 1
	m := []int{}
	ret := 0

	for _, n := range nums {
		if n%2 == 1 {
			m = append(m, index)
			index = 0
			i++
		} else {
			index++
		}
	}
	m = append(m, index)
	for loc, num := range m {
		if loc+k < len(m) {
			ret += (m[loc+k] + 1) * (num + 1)
		}
	}
	return ret
}

// func numberOfSubarrays(nums []int, k int) int {
// 	index := 0
// 	i := 1
// 	m := map[int]int{}
// 	ret := 0

// 	for _, n := range nums {
// 		if n%2 == 1 {
// 			m[i] = index
// 			index = 0
// 			i++
// 		} else {
// 			index++
// 		}
// 	}
// 	m[i] = index
// 	fmt.Println(m)
// 	for loc, num := range m {
// 		fmt.Println(loc, loc+k)
// 		if x1, ok := m[loc+k]; ok {
// 			ret += (x1 + 1) * (num + 1)
// 		}
// 	}
// 	return ret
// }

// func numberOfSubarrays(nums []int, k int) int {
// 	dp := make([]int, 0)
// 	cnt, ret := 0, 0
// 	for i := 0; i < len(nums); i++ {
// 		cnt++
// 		if nums[i]%2 == 1 {
// 			dp = append(dp, cnt)
// 			cnt = 0
// 		}
// 		if len(dp) >= k {
// 			ret += dp[len(dp) - k]
// 		}
// 	}
// 	return ret
// }
