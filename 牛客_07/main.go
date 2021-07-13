package main

/**
 *
 * @param prices int整型一维数组
 * @return int整型
 */
func maxProfit( prices []int ) int {
	// write code here
	cost := 0
	min := prices[0]
	for _, num := range prices{
		if num > min{
			if num - min > cost{
				cost = num - min
			}
		}else{
			min = num
		}
	}
	return cost
}