package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("aacdefcaa"))
}

//  s ...............
//s(re)
//.
//.
//.

func longestPalindrome(s string) string {
	ls := len(s)
	if ls <= 1 {
		return s
	}

	mStr := make([]byte, ls*2+1)
	i, j := 0, 0
	for i < ls {
		mStr[j] = 0
		j++
		mStr[j] = s[i]
		j++
		i++
	}
	mStr[j] = 0
	ls = j + 1

	radius := make([]int, ls)
	r, c, max, mid := -1, -1, -1, 0
	for i = 0; i < ls; i++ {
		if r > i {
			if radius[2*c-i] > r-i+1 {
				radius[i] = r - i + 1
			} else {
				radius[i] = radius[2*c-i]
			}
		} else {
			radius[i] = 1
		}
		for i+radius[i] < ls && i-radius[i] > -1 {
			if mStr[i-radius[i]] == mStr[i+radius[i]] {
				radius[i]++
			} else {
				break
			}
		}
		if i+radius[i] > r {
			r = i + radius[i] - 1
			c = i
		}
		if max < radius[i] {
			max = radius[i]
			mid = i
		}
	}
	max--
	mid = mid/2 - max/2
	return s[mid : mid+max]
}

// func longestPalindrome(s string) string {
// 	if len(s) == 0 {
// 		return s
// 	}
// 	re := reverse(s)
// 	max := 1
// 	indexx := 1
// 	c := len(s)
// 	pb := [][]int{}
// 	for i := 0; i < c; i++ {
// 		pb = append(pb, make([]int, c))
// 	}
// 	for x := 0; x < c; x++ {
// 		if s[x] == re[0] {
// 			pb[0][x] = 1
// 		}
// 	}
// 	for y := 0; y < c; y++ {
// 		if s[0] == re[y] {
// 			pb[y][0] = 1
// 		}
// 	}
// 	fmt.Println(pb)
// 	for y := 1; y < c; y++ {
// 		for x := 1; x < c; x++ {
// 			if re[y] == s[x] {
// 				pb[y][x] = pb[y-1][x-1] + 1
// 				if pb[y][x] > max {
// 					max = pb[y][x]
// 					indexx = x + 1
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println(indexx, max)
// 	return s[indexx-max : indexx]
// }

// func reverse(str string) string {
// 	var result string
// 	strLen := len(str)
// 	for i := 0; i < strLen; i++ {
// 		result = result + fmt.Sprintf("%c", str[strLen-i-1])
// 	}
// 	return result
// }
