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
	// calculate sum and set value to array
	ansArray := []int{}
	carry := 0
	for i := 0; l1 != nil || l2 != nil; i++ {
		l1Val := setVal(l1)
		l2Val := setVal(l2)
		sum := l1Val + l2Val + carry
		ansArray = append(ansArray, sum%10)
		carry = sum / 10
		l1 = setListNode(l1)
		l2 = setListNode(l2)
	}
	// set the carry to the end of the array if the last calculation is carrying out
	if carry == 1 {
		ansArray = append(ansArray, carry)
	}

	// convert to ListNode from array
	ans := ListNode{ansArray[0], nil}
	tmp := &ans
	for _, v := range ansArray[1:] {
		tmp.Next = &ListNode{v, nil}
		tmp = tmp.Next
	}

	return &ans
}

//For simple test
func toArray(l *ListNode) []int {
	result := []int{}

	for i := 0; l != nil; i++ {
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
