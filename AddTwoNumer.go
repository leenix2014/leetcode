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
	sum := addTwoNumbers(toList(342), toList(465))
	fmt.Println(toInt(sum))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := toInt(l1) + toInt(l2)
	return toList(sum)
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
