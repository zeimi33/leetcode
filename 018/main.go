package main

// func searchInsert(nums []int, target int) int {
//     for i:=0;i<len(nums);i++{
//         if target <= nums[i]{
//             return i
//         }
//     }
//     return len(nums)
// }
func searchInsert(nums []int, target int) int {
	if len(nums)==1{
		if nums[0] >=target{
			return 0
		}else{
			return 1
		}
	}
	if nums[len(nums)/2]>=target{
		return searchInsert(nums[:len(nums)/2],target)
	}else{
		a := len(nums)/2
		return a+searchInsert(nums[len(nums)/2:],target)
	}
}//就爱用递归，哈哈
