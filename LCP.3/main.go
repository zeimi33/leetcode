package main

func main() {
	robot("UURRR", nil, 12, 3)
}

func robot(command string, obstacles [][]int, x int, y int) bool {
	indexx := 0
	indexy := 0
	step := map[int][]int{}
	for _, c := range command {
		if c == 'R' {
			indexx++
			if step[indexx] == nil {
				yy := []int{}
				step[indexx] = yy
			}
			step[indexx] = append(step[indexx], indexy)
		}
		if c == 'U' {
			indexy++
			if step[indexx] == nil {
				yy := []int{}
				step[indexx] = yy
			}
			step[indexx] = append(step[indexx], indexy)
		}
	}
	zhouqi := x / indexx
	
	return false
}
