package main

//import . "nc_tools"
func main(){
	a:=&ListNode{1,nil}
	b:=&ListNode{2,nil}
	c:=&ListNode{3,nil}
	d:=&ListNode{4,nil}
	e:=&ListNode{5,nil}
	f:=&ListNode{6,nil}
	g:=&ListNode{7,nil}
	h:=&ListNode{8,nil}
	i:=&ListNode{9,nil}
	j:=&ListNode{10,nil}
	a.Next =b
	b.Next=c
	c.Next =d
	d.Next =e
	e.Next =f
	f.Next = g
	g.Next = h
	h.Next = i
	i.Next =j
	reorderList(a)

}

type ListNode struct{
  Val int
  Next *ListNode
}


/**
 *
 * @param head ListNode类
 * @return void
 */
// 1 2 3 4
func reorderList( head *ListNode )  {
	if head == nil {
		return
	}
	step1 := head
	step2 := head
	step1Pre := head
	for step1 != nil && step2 != nil {
		step1Pre = step1
		step1 = step1.Next
		step2 = step2.Next
		if step2 == nil {
			break
		}
		step2 = step2.Next
	}
	step1Pre.Next = nil
	l1 := head
	l2 := reverse(step1)
	head =  merge(l1,l2)
	return
}

func reverse(head *ListNode)*ListNode{
	if head == nil {
		return nil
	}

	a1 := head.Next //2
	head.Next =nil
	// 1 - > 2 -> 3
	for a1 != nil {
		a2 := a1.Next //3
		a1.Next = head//2 -> 1
		head = a1
		a1 = a2
	}
	return head
}
// 1 2 3 1 2 3 4
// 1 4 2 3 3 2 1
func merge(l1 ,l2 *ListNode)*ListNode{
	l1iter := l1
	l2iter := l2
	// 1 2 3
	// 4 5 6
	for l2iter != nil {
		p1 := l1iter.Next//2
		p2 := l2iter.Next//5
		l1iter.Next = l2iter // 1 -> 4
		l2iter.Next = p1 // 1 -> 4 -> 2 -> 3
		l1iter = p1
		l2iter = p2
	}
	return l1
}