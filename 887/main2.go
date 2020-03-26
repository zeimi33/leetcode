package main



func main(){
	a := int32
}



func superEggDrop(K int, N int) int {
	dp := make([][]int, K+1)
	for i := 0; i <= K; i++ {
		dp[i] = make([]int, N+1)
	}

	// 只有 0 个鸡蛋，只能确定 0 层
	// 只要 1 个鸡蛋，最多只能确定 i 层
	// k 个鸡蛋，移动 1 步只能确定 1 层楼
	dp[0][1] = 0
	for i := 1; i < K; i++ {
		dp[i][1] = 1
	}

	for n := 0; n < N; n++ {
		dp[1][n] = n
	}
	for step := 1; step <= N; step++ {
		for k := 1; k <= K; k++ {
			dp[k][step] = dp[k][step-1] + dp[k-1][step-1] + 1
			if dp[k][step] >= N {
				return step
			}
		}
	}
	return N
}