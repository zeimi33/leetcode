package main
func main(){

}

func findPeakElement(nums []int) int {//这道题没有边界条件会变得极为复杂,边界条件是 两边 都是-∞
	l := 0
	r := len(nums)-1
	for l<r{
		if nums[(l+r)/2] <nums[(l+r)/2+1]{
			l = (l+r)/2+1
		}else{
			r = (l+r)/2
		}
	}
	return l
}
