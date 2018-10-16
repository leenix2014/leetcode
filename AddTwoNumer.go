package main

import (
	"fmt"
	"math"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	sum := addTwoNumbers(toList(19999999), toList(11))
	fmt.Println(toInt(sum))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var first *ListNode
	var prev *ListNode
	carry := 0
	for {
		sum := getVale(l1) + getVale(l2) + carry
		node := &ListNode{Val: sum % 10}
		if first == nil {
			first = node
			prev = first
		} else {
			prev.Next = node
			prev = prev.Next
		}
		carry = sum/10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if l1 == nil && l2 == nil {
			if carry > 0 {
				node := &ListNode{Val: carry}
				prev.Next = node
			}
			break
		}
	}
	return first
}

func getVale(l *ListNode) int {
	if l == nil {
		return 0
	}
	return l.Val
}

func toInt(l *ListNode) int {
	num := 0
	n := 0
	for {
		num += l.Val * int(math.Pow10(n))
		n++
		if l.Next == nil {
			break
		}
		l = l.Next
	}
	return num
}

func toList(n int) *ListNode {
	var first *ListNode
	var prev *ListNode
	for {
		mod := n % 10
		n = n / 10
		node := &ListNode{Val: mod}
		if first == nil {
			first = node
			prev = first
		} else {
			prev.Next = node
			prev = prev.Next
		}
		if n <= 0 {
			break
		}
	}
	return first
}
