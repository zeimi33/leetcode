package main

//在一个数组里面移除指定value，并且返回新的数组长度.不能新建另一个数组。
import "fmt"

func main() {
	num := 0
	var array []int
	tem := 0
	rem := 0
	fmt.Println("输入数组的个数:")
	fmt.Scanf("%d", &num)
	fmt.Println("输入数组:")
	for i := 0; i < num; i++ {
		fmt.Scanf("%d", &tem)
		array = append(array, tem)
	}
	fmt.Println("输入数组中移除的数字:")
	fmt.Scanf("%d", &rem)
	for i := 0; i < num; i++ {
		if rem == array[i] {
			array = append(array[0:i], array[i+1:]...)
			i--
			num--
		}
	}
	fmt.Println(array)
}
