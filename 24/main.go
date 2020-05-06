package main

func main() {
	a := ListNode{1, nil}
	b := ListNode{2, nil}
	c := ListNode{3, nil}
	d := ListNode{4, nil}
	a.Next = &b
	b.Next = &c
	c.Next = &d
	swapPairs(&a)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head != nil && head.Next == nil {
		return head
	}

	b := head.Next.Next
	a := head
	head = head.Next
	head.Next = a

	//fmt.Println(head.Val, head.Next.Val)
	//	fmt.Println(b.Val, b.Next.Val)
	head.Next.Next = swapPairs(b)
	return head
}
