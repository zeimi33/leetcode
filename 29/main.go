package main

import (
	"fmt"
)

func main() {
	fmt.Println(divide(-2147483648, -1))
}


func divide(dividend int, divisor int) int {
	index := 1
	if dividend<0 && divisor<0{
		dividend = -dividend
		divisor = -divisor
	}
	if (dividend >0 && divisor <0)||(dividend<0 &&divisor>0){
		index =-1
		if dividend <0 {
			dividend = 0-dividend
		}
		if divisor <0 {
			divisor = 0-divisor
		}
	}

	bei := 0
		for dividend >= divisor {
			weiyi := 0
			for dividend >=divisor{
				n := dividend >>weiyi
				fmt.Println(dividend,weiyi)
				if n<=divisor {
					break
				}
				weiyi++
			}
			fmt.Println(weiyi,"weiyi")
						if weiyi>=1{
				weiyi--
			}
			if weiyi ==0{
				if dividend >=divisor {
					bei++
				}
			}

			dividend -= divisor<<weiyi
			fmt.Println(dividend)
			bei += beishu(weiyi)
		}

		if index >0{
			if bei > 2147483647{
				return 2147483647
			}
			return bei
		}
		if bei > 2147483648{
			return -2147483647
		}
		return  0-bei
}

func beishu(a int)int{
	if a >0{
		return int(math.Exp2(float64(a)))
	}
	return 0
}
