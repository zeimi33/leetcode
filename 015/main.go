package main

import (
	"fmt"
	"math"
)
//"38455498359"
//"999999999999999999"
func main(){
	a := superpalindromesInRange("L = \"38455498359\"","R = \"999999999999999999\"")
	fmt.Println(a)
}

func superpalindromesInRange(L string, R string) int {
		l:=0
		fmt.Sscanf(L,"L = \"%d\"",&l)//nolint :handle err
		r:=0
		fmt.Sscanf(R,"R = \"%d\"",&r)
		l_min := int(math.Sqrt(float64(l)))
		r_min := int(math.Sqrt(float64(r)))
		ret :=0
		for ;l_min<=r_min;l_min++{
			if solution(l_min)&& solution(l_min*l_min){
				fmt.Println(l_min)
				ret++
			}
		}
		return ret

}
func solution(num int)bool{
	if 	num <0 {
		return false
	}
	tmp := num
	num_test :=0
	for num>0{
		num_test = num_test*10 + num %10
		num = num/10
	}
	return  tmp == num_test
}