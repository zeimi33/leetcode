//返回 A 的最短的非空连续子数组的长度，该子数组的和至少为 K 。

// 如果没有和至少为 K 的非空子数组，返回 -1 。

//

// 示例 1：

// 输入：A = [1], K = 1
// 输出：1
// 示例 2：

// 输入：A = [1,2], K = 4
// 输出：-1
// 示例 3：

// 输入：A = [2,-1,2], K = 3
// 输出：3
//

// 提示：

// 1 <= A.length <= 50000
// -10 ^ 5 <= A[i] <= 10 ^ 5
// 1 <= K <= 10 ^ 9
package main

type stack struct {
	no []int
}

func (s *stack) pop() {
	s.no = s.no[:len(s.no)-1]
}
func (s *stack) push(num int) {
	s.no = append(s.no, num)
}
func (s *stack) top() int {
	return s.no[len(s.no)-1]
}
func (s *stack) empty() bool {
	return len(s.no) == 0
}

func main() {

}
func shortestSubarray(A []int, K int) int {
	i := len(A)
	res := -1
	last := 0
	next := []int{}
	for num := 0; num < i; num++ {
		if A[num] >= K {
			return 1
		}
		next = append(next, last+A[num])
		last = last + A[num]
	}
	s := stack{}
	s.no = next
	j := 0
	for j < i {

	}

	return res
}
