package main

func main(){

}


func maximalSquare(matrix [][]byte) int {
	ret := 0
	if len(matrix)==0{
		return 0
	}
	for i:=0 ;i<len(matrix);i++{
		for j:=0;j<len(matrix[0]);j++{
			if matrix[i][j]==0{
				continue
			}
			ret = max(ret,dfs(&matrix,i,j))
		}
	}
	return ret
}

func maximalSquare(matrix [][]byte) int {
    dp := make([][]int, len(matrix))
    maxSide := 0
    for i := 0; i < len(matrix); i++ {
        dp[i] = make([]int, len(matrix[i]))
        for j := 0; j < len(matrix[i]); j++ {
            dp[i][j] = int(matrix[i][j] - '0')
            if dp[i][j] == 1 {
                maxSide = 1
            }
        }
    }

    for i := 1; i < len(matrix); i++ {
        for j := 1; j < len(matrix[i]); j++ {
            if dp[i][j] == 1 {
                dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
                if dp[i][j] > maxSide {
                    maxSide = dp[i][j]
                }
            }
        }
    }
    return maxSide * maxSide
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
