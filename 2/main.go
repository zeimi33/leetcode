package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	mod := 0
	add := 0
	head := new(ListNode)
	ptr := head
	for l1 != nil || l2 != nil || add > 0 {
		num1 := 0
		num2 := 0
		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}
		mod = (num1 + num2 + add) / 10
		add = (num1 + num2 + add) % 10
		ptr.Next = new(ListNode)
		ptr.Next.Val = mod
		ptr = ptr.Next
	}
	return head.Next
}
