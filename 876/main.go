package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := ListNode{1, nil}
	b := ListNode{2, nil}
	c := ListNode{3, nil}
	d := ListNode{4, nil}
	e := ListNode{5, nil}
	a.Next = &b
	b.Next = &c
	c.Next = &d
	d.Next = &e
	e.Next = nil
	f := middleNode(&a)
	fmt.Println(f.Val)
}
func middleNode(head *ListNode) *ListNode {
	a := head
	index := 0
	for a.Next != nil {
		a = a.Next
		index++
		fmt.Println(a.Val, index)
	}
	fmt.Println(index)
	if (index % 2) == 0 {
		index = index / 2
	} else {
		index = index/2 + 1
	}
	fmt.Println(index)
	a = head
	for index != 0 {
		index--
		a = a.Next
	}
	return a
}
给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。

示例:

输入: 
words = ["oath","pea","eat","rain"] and board =
[
  ['o','a','a','n'],
  ['e','t','a','e'],
  ['i','h','k','r'],
  ['i','f','l','v']
]

输出: ["eat","oath"]
