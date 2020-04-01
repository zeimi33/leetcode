package main 

import "fmt"

func main(){
	a := []int{1,2,3,4,1,2,3,4}
	b := hasGroupsSizeX(a)
	fmt.Println(b)
}

func hasGroupsSizeX(deck []int) bool {
	index := make(map[int]int)
	min := 10000000
	for i:=0 ; i< len(deck) ; i++{
		index[deck[i]]++
		min = minfind(min,index[deck[i]])

	}

	gongyue :=0
	for _,num := range index {
		gongyue = num
		break
	}
	for _,num := range index {
		gongyue	= gcd(gongyue,num)
		fmt.Println(gongyue)
		if gongyue ==1 {
			return false
		}
	}
	return true
}

 func minfind(a int , b int)int {
	if a > b {
		return b
	}
	return a 
 }
 func gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	} 
	return y
	
}