package main

import "fmt"

func main() {
	c := strongPasswordChecker("aaaaabbbb1234567890ABA")
	fmt.Println(c)
}

func strongPasswordChecker(s string) int {
	if len(s) == 0 {
		return 6
	}
	lianxuzero := []int{}
	lianxutwo := []int{}
	lianxuone := []int{}
	numArray := len(s)
	del := 0
	add := 0
	step := 0
	var before rune
	beforenum := 0
	small := false
	big := false
	nnum := false
	if numArray > 20 {
		del = numArray - 20
	}
	if numArray < 6 {
		add = 6 - numArray
	}
	for _, c := range s {
		if !nnum && c >= '0' && c <= '9' {
			nnum = true
		}
		if !small && c >= 'a' && c <= 'z' {
			small = true
		}
		if !big && c >= 'A' && c <= 'Z' {
			big = true
		}
		if before != c {
			before = c
			if beforenum >= 3 {
				if beforenum%3 == 1 {
					lianxuone = append(lianxuone, beforenum)
				}
				if beforenum%3 == 2 {
					lianxutwo = append(lianxutwo, beforenum)
				}
				if beforenum%3 == 0 {
					lianxuzero = append(lianxuzero, beforenum)
				}
			}
			beforenum = 1
		} else {
			beforenum++
		}
	}
	if beforenum >= 3 {
		if beforenum%3 == 1 {
			lianxuone = append(lianxuone, beforenum)
		}
		if beforenum%3 == 2 {
			lianxutwo = append(lianxutwo, beforenum)
		}
		if beforenum%3 == 0 {
			lianxuzero = append(lianxuzero, beforenum)
		}
	}
	if add > 0 {
		if !big {
			step++
		}
		if !small {
			step++
		}
		if !nnum {
			step++
		}
		if add >= step {
			return add
		}

		return step
	}
	if add == 0 && del == 0 {
		fmt.Println("aaa")
		num := 0
		lianxu := []int{}
		lianxu = append(lianxu, lianxuone...)
		lianxu = append(lianxu, lianxutwo...)
		lianxu = append(lianxu, lianxuzero...)
		fmt.Println(lianxu)
		for _, i := range lianxu {
			num += i / 3
		}
		if !big {
			step++
		}
		if !small {
			step++
		}
		if !nnum {
			step++
		}
		if step > num {
			return step
		}
		return num
	}
	for del != 0 && (len(lianxuone) > 0 || len(lianxutwo) > 0 || len(lianxuzero) > 0) {
		fmt.Println(lianxuzero, lianxuone, lianxutwo, del)
		for i, n := range lianxuzero {
			if del == 0 {
				break
			}
			beforenum = n - 1
			if beforenum > 3 {
				lianxutwo = append(lianxutwo, beforenum)
			}
			del--
			step++
			lianxuzero = append(lianxuzero[:i], lianxuzero[i+1:]...)
			break
		}
		if len(lianxuzero) > 0 {
			continue
		}
		for i, n := range lianxuone {
			fmt.Println(n, del)
			if del == 0 {
				break
			}
			beforenum = n - 1
			if beforenum >= 3 {
				lianxuzero = append(lianxuzero, beforenum)
			}
			del--
			step++
			lianxuone = append(lianxuone[:i], lianxuone[i+1:]...)
			break
		}
		if len(lianxuone) > 0 || len(lianxuzero) > 0 {
			continue
		}
		for i, n := range lianxutwo {
			fmt.Println(i, n, "aaaa")
			if del == 0 {
				break
			}
			beforenum = n - 1
			if beforenum > 3 {
				lianxuone = append(lianxuone, beforenum)
			}
			del--
			step++
			lianxutwo = append(lianxutwo[:i], lianxutwo[i+1:]...)
			break
		}
		continue
	}
	fmt.Println(lianxuzero, lianxutwo, lianxuone, del, big, small, nnum)
	n := 0
	for _, i := range lianxuone {
		n += (i / 3)

	}
	for _, i := range lianxutwo {
		n += (i / 3)

	}
	for _, i := range lianxuzero {
		n += (i / 3)
	}
	zz := 0
	if !big {
		zz++
	}
	if !small {
		zz++
	}
	if !nnum {
		zz++
	}
	if zz > n {
		return zz + step + del
	}
	return n + step + del
}
