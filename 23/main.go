package main

func main(){


}


  type ListNode struct {
      Val int
      Next *ListNode
  }
 
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0{
        return nil
    }
	if len(lists) ==1{
			return lists[0]
	}
	if len(lists) ==2{
		return transf(lists[0],lists[1])
	}
	return transf(mergeKLists(lists[:len(lists)/2]),mergeKLists(lists[len(lists)/2:]))
}

func transf(list1 *ListNode,list2 *ListNode)*ListNode{
	var head  *ListNode
	index :=  &ListNode{}
	head = index
	for {
		if list1 == nil{
			index.Next = list2
            break
		}
		if list2 == nil {
			index.Next = list1
            break
		}
		if list1.Val < list2.Val{
			index.Next = list1
			index = index.Next
			list1 = list1.Next
		}else{
			index.Next = list2
			index = index.Next
			list2 = list2.Next
		}
	}
	return head.Next
}	