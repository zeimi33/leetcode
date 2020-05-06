func removeElement(nums []int, val int) int {
	if len(nums)==0{
		return 0
	}
	s:=0
	for i:=0;i<len(nums)-s;i++{
		if nums[i]==val{
			s++
			for j:=i;j<len(nums)-1;j++{
				nums[j]=nums[j+1]
			}
			i--
		}
	}
	return len(nums)-s
}
