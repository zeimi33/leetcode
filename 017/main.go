package main

import "fmt"

func main(){
	a := "DID"
	fmt.Println(numPermsDISequence(a))
}
func numPermsDISequence(S string) int {
	mod := 10000000007
	dp := make([][]int,len(S)+1)
	dp[0] = []int{1,1}
	for i:=1;i<=len(S);i++{
		dp[i] = make([]int,len(S)+1)
		for j:=0 ;j <= i ;j++{
			if S[i-1] =='D'{
				for n :=j;n<=i-1;n++{
					dp[i][j] = (dp[i][j]+dp[i-1][n])%mod
				}
			}else{
				for n:= 0 ;n<=j-1;n++{
					dp[i][j] = (dp[i][j]+dp[i-1][n])%mod
				}
			}
		}
	}
	ret := 0
	for i:= 0 ;i<= len(S);i++{
		ret = (ret +dp[len(S)][i])%mod
	}
	return ret
}
