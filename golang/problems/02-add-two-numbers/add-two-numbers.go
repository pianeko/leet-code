package main

import "fmt"

//Type definition
type ListNode struct {
	Val  int
	Next *ListNode
}

//My solution
func setVal(l *ListNode) int {
	if l == nil {
		return 0
	}
	return l.Val
}
func setListNode(l *ListNode) *ListNode {
	if l == nil {
		return l
	}
	return l.Next
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//calculate sum and set value to ListNode
	ans := ListNode{0, nil}
	//temporary variable where the pointer is stored
	next := &ans
	carry := 0
	for !(l1 == nil && l2 == nil && carry == 0) {
		//pointer to ListNode stored in Next
		tmp := &ListNode{0, nil}
		sum := setVal(l1) + setVal(l2) + carry
		//ignore the carry and retrieve the value
		tmp.Val = sum % 10
		//retrieve the carry
		carry = sum / 10
		//set pointer to Next
		next.Next = tmp
		//set next pointer to temporary variable
		next = next.Next
		//set next pointer to each arg
		l1 = setListNode(l1)
		l2 = setListNode(l2)
	}

	return ans.Next
}

//For simple test
func toArray(l *ListNode) []int {
	result := []int{}

	for l != nil {
		result = append(result, l.Val)
		l = l.Next
	}

	return result
}
func main() {
	l1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	ans := addTwoNumbers(&l1, &l2)
	fmt.Printf("%v\n", toArray(ans))
}
