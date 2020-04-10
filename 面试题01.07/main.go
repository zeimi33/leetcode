package main

func main() {

}

func rotate(matrix [][]int) {
	n := len(rotate)
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matmatrix[n-i-1][j] = matmatrix[n-i-1][j], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j],matrix[j][i] =matrix[j][i] ,matrix[i][j]
		}
	}
}
