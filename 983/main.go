package main

func main(){
	fmt.Println()
}

func mincostTickets(days []int, costs []int) int {
	dp := make([]int,len(days))
	for i,v := range  days{
		for j:=i ;j<len(days);j++{
			if days[i] - v ==0{
				if dp[j]==0{
					dp[j] = costs[0]
				}
				continue
			}
			if days[i]-v >1 && days[i]-v <7 {
				min := 
			}
		}
	}
}