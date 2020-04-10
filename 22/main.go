package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"()"}
	}

	ret := []string{}
	for a := 0; a < n; a++ {
		int_ret := []string{}
		ret1 := generateParenthesis(n - a - 1)
		ret2 := generateParenthesis(a)
		//fmt.Println(ret1, ret2)
		for _, str1 := range ret1 {

			int_ret = append(int_ret, fmt.Sprintf("(%s)", str1))

		}
		//	fmt.Println(int_ret)
		for _, str2 := range ret2 {
			for _, str3 := range int_ret {
				fmt.Println(str2, str3)
				ret = append(ret, fmt.Sprintf("%s%s", str3, str2))
			}
		}

	}
	return ret
}
