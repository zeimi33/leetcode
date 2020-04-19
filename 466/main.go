package main

import "fmt"

func main() {
	fmt.Println(getMaxRepetitions("ecbafedcba", 4, "abcde", 1))
}

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	index2 := 0
	num1 := 0
	num2 := 0
	index1 := 0
	ret := 0
	var c1 rune
	m := make(map[int]int)
	for i := 0; i != -1; i++ {
		if i == 0 {
			m[i+1] = 0
		} else {
			m[i+1] = m[i]
		}
		for index1, c1 = range s1 {
			//	fmt.Println(index1, index2, rune(c1), rune(s2[index2]))
			if index2 == 0 && index1 == 0 && i != 0 {
				if n1 > num1 {
					//fmt.Println(n1, num1)
					ret = (n1/num1)*m[num1] + m[n1%num1]
				} else {
					ret = m[n1]
				}
				return ret / n2
			}
			if byte(c1) == s2[index2] {
				if index2 == len(s2)-1 {
					index2 = 0
					num2++
					m[num1+1] = num2
				} else {
					index2++
				}
			}
			//fmt.Println(m)
			//fmt.Println(index2, index1, i)

		}
		//fmt.Println(num1, n1)
		if num1+1 == n1 {
			return m[num1+1] / n2
		}
		num1++
	}

	return 0
}

//答案优化方法。
func getMaxRepetitions1(s1 string, n1 int, s2 string, n2 int) int {
	len1, len2 := len(s1), len(s2)
	index1, index2 := 0, 0 // 注意此处直接使用 Ra Rb 的下标，不取模

	if len1 == 0 || len2 == 0 || len1*n1 < len2*n2 {
		return 0
	}

	map1, map2 := make(map[int]int), make(map[int]int)
	ans := 0 // 注意，此处存储的是 Ra 中 Sb 的个数，而非 Ra 中 Rb 的个数

	for index1/len1 < n1 { // 遍历整个 Ra
		if index1%len1 == len1-1 { //在 Sa 末尾
			if val, ok := map1[index2%len2]; ok { // 出现了循环，进行快进
				cycleLen := index1/len1 - val/len1                 // 每个循环占多少个 Sa
				cycleNum := (n1 - 1 - index1/len1) / cycleLen      // 还有多少个循环
				cycleS2Num := index2/len2 - map2[index2%len2]/len2 // 每个循环含有多少个 Sb

				index1 += cycleNum * cycleLen * len1 // 将 Ra 快进到相应的位置
				ans += cycleNum * cycleS2Num         // 把快进部分的答案数量加上
			} else { // 第一次，注意存储的是未取模的
				map1[index2%len2] = index1
				map2[index2%len2] = index2
			}

		}

		if s1[index1%len1] == s2[index2%len2] {
			if index2%len2 == len2-1 {
				ans += 1
			}
			index2 += 1
		}
		index1 += 1
	}
	return ans / n2
}
