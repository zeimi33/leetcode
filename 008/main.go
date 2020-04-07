package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "-13+8"
	fmt.Println(myAtoi(a))
}

func myAtoi(str string) int {
	iNT_MAX := 2147483647
	iNT_MIN := 2147483648
	index := 1
	singal := false
	c := []rune{}
	num := 0
	for _, a := range str {

		if a == ' ' {
			if len(c) == 0 && !singal {
				continue
			} else {
				break
			}
		}
		if a == '+' || a == '-' {
			if len(c) == 0 {
				if !singal {
					singal = true
					if a == '-' {
						index = -1
					}
					continue
				} else {
					return 0
				}
			} else {
				break
			}
		}
		fmt.Println(a)
		if a < '0' || a > '9' {
			break
		}

		c = append(c, a)
		num++
	}
	fmt.Println(c)
	if len(c) == 0 {
		return 0
	}

	ret, _ := strconv.Atoi(string(c))
	if index == -1 {
		if ret > iNT_MIN {
			return index * iNT_MIN
		}
	} else {
		if ret > iNT_MAX {
			return iNT_MAX
		}
	}
	return index * ret
}
