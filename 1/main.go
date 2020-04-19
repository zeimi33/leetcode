package main

func main() {

}

func twoSum(nums []int, target int) []int {
	a := []int{}
	for i := 0; i < len(nums); i++ {
		for c := i + 1; c < len(nums); c++ {
			if nums[i]+nums[c] == target {
				return []int{i, j}
			}
		}
	}

	return a
}
