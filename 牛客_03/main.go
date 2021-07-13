package main
type ListNode struct{
	Val int
	Next *ListNode
}
func EntryNodeOfLoop(pHead *ListNode) *ListNode{
	iter := pHead
	m := make(map[*ListNode]interface{})
	for iter != nil {
		if _, ok := m[iter];ok{
			return iter
		}else {
			m[iter] = struct {}{}
		}
		iter = iter.Next
	}
	return nil
}
