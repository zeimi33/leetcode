package main

func main(){

}


  type ListNode struct {
      Val int
      Next *ListNode
  }
 
func reverseKGroup(head *ListNode, k int) *ListNode {
	count := 0
	fakehead := head
	index := head
	m := []*ListNode{}
	for i:=0 ;i<k;i++{
		m = append(m,index)
		if index == nil {
			return head
		}
		index = index.Next
		count++
	} 
for i := count-1;i>0;i--{
	m[i].Next = m[i-1]
}
	m[0].Next = reverseKGroup(index,k)
	return m[k-1]
}