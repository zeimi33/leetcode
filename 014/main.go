package main
func main(){

}
type stack  struct {
	no []int
}
func (s * stack)pop(){
	s.no = s.no[:len(s.no)-1]
}
func (s *stack)push(num int){
	s.no = append(s.no,num)
}
func (s *stack)top()int{
	return s.no[len(s.no)-1]
}

func largestRectangleArea(heights []int) int {
	stack :=stack{}
	stack.push(-1)
	max :=0
	for i:=0;i<len(heights);i++{
		for stack.top() !=-1 &&heights[stack.top()]>=heights[i]{//新加入的小于栈顶 进行计算
				top := stack.top()//栈顶
				stack.pop()
				area:=heights[top]*(i-1-stack.top())//i为当前的位
				if area >max {
					max = area
				}
		}
		stack.push(i)
	}
	for stack.top() !=-1{
		top := stack.top()
		stack.pop()
		area :=heights[top] *(len(heights)-1-stack.top())
		if area >max {
			max =area
		}
	}
	return max
}
