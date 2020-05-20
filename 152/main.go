package main

func main(){

}

func maxProduct(nums []int) int {
	max := 1
	min := 1
	mutil := 1
	ret := nums[0]
	for _, v := range nums{
		if v  < 0 {
			tmp := max
			max = min
			min = tmp
		}
		min = fmin(min*v,v)
		max = fmax(max*v,v)
		ret = fmax(max,ret)
	}
	return ret
}


func fmax(a,b int)int{
	if a>b{
		return a 
	}
	return b
}

func fmin(a,b int)int{
	if a < b {
		return a
	}
	return b
}