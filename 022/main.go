package main

import "fmt"

func main(){
	a := return_bit_of_two(uint32(100))
	fmt.Println(a)
}
//题目翻译: 给出一个整数，求它包含二进制1的位数。
// 例如，32位整数 11 的二进制表达形式
// 是 00000000000000000000000000001011 ，那么函数应该返回3。

func return_bit_of_two (num uint32)int{
	ret := 0
	for num > 0 {
		ret += (int)(num %2)
		num = num/2
	}
	return ret
}