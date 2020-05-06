package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	index := []*ListNode{}
	fakehead := &ListNode{}
	fakehead.Next = head
	index = append(index, fakehead)
	indexM := fakehead
	for indexM != nil {
		index = append(index, indexM)
		indexM = indexM.Next
	}
	index = append(index, nil)
	i := len(index)

	index[i-n-3].Next = index[i-n-1]
	return fakehead.Next
}

//双指针法
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first := head
	seconde := head
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		seconde = seconde.Next
	}
	seconde.Next = seconde.Next.Next
	return head
}
