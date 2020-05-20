package main

import "fmt"

func main(){
	fmt.Println(mySqrt(4))
}
func mySqrt(x int) int {
	l :=0
	r := x
	var ret int 
	for l <= r {
		 mid := l + (r - l) >> 1
		if mid * mid <=x{
			ret =mid
			l = mid+1
			continue
		}
		r = mid -1
	}
	return ret
}

// func mySqrt(x int) int {
//     l, r := 0, x
//     ans := -1
//     for l <= r {
//         mid := l + (r - l) >> 1
//         if mid * mid <= x {
//             ans = mid 
//             l = mid + 1
//         } else {
//             r = mid - 1
//         }
//     }
//     return ans 
// }
