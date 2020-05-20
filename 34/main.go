package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{2,2},3))
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0{
		return []int{-1,-1}
	}
	index := -1
	l := 0
	r := len(nums)-1
	for l <= r {
		mid := (l + r)/2
		if nums[mid] == target{
			index = mid
			break
		}
		if nums[mid] > target{
			r = mid -1
		}else{
			l = mid +1
		}

	}
	if index == -1 {
		return []int{-1,-1}
   }
	if len(nums) ==1 {
		return []int{0,0}
	}

	return []int{findLeft(nums,target,index),findRight(nums,target,index)}
}
func findLeft(nums []int, target int, index int) int{
	l := 0 
	r := index 
	for l < r {
		mid := (l + r)/2
		fmt.Println("findl",l,r)
		if nums[mid] == target{
			r = mid
			continue
		}
		if nums[mid] < target{
			l = mid+1
		}
	}
	return l
}

func findRight(nums []int, target int, index int) int{
	l := index
	r := len(nums)-1
	for l < r {
		fmt.Println(l,r)
		mid := (l + r)/2
		if nums[mid] == target{
			if l == mid {
				if nums[l+1] != target {
					return l
				}
				l++
			}else{
				l = mid
			}
		}
		if nums[mid] > target{
			r = mid-1
		}
	}
	return l
}