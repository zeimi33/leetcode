package main

func main() {

}
func singleNumber(nums []int) int {
	a := 0
	for _, v := range nums {
		a = a ^ v
	}
	return a
}
