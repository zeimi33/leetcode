package main

func main() {

}

func maxDistance(grid [][]int) int {
	hang := len(grid)
	raw := hang
	ret := make(chan int)
	guandao := make(chan int)
}

func findMax(ret chan int, guandao chan int, index int) {
	max := 0
	num := 0
	for i := 0; i < index; i++ {
		num = <-guandao
		if max < num {
			max = num
		}
	}
	index <- num
}

func checkLast(grid [][]int, x, y int, guandao chan int) {
	ret :=0
	num := len(grid)
	for i := 0; i < num; i++ {//x移动
		for k := 0; k < num; k++ {//y移动
			if()
		}
	}
}