package main

func main(){

}

func singleNumbers(nums []int) []int {
	c := 0
	for _,v := range nums{
		c^=v
	}
	a1 := 0
	a2 := 0
	mask := c&(-c)
	for _,v :=range nums{
		if v&mask ==0{
			a1 ^=v
		}else{
			a2^=v
		}
	}
	return []int{a1,a2}
}