package Linked_list

import "math"

//Double a Number Represented as a Linked List
//You are given the head of a non-empty linked list representing a non-negative integer without leading zeroes.
//
//Return the head of the linked list after doubling it.

func doubleIt(head *ListNode) *ListNode {
	num := []int{}
	curr := head
	for curr != nil {
		num = append(num, head.Val)
		curr = curr.Next
	}
	res := 0
	for i := 0; i < len(num); i++ {
		res += num[i] * int(math.Pow(float64(10), float64(len(num)-1-i)))
	}
	res = res * 2
	number := []int{}
	rem := 0
	for res != 0 {
		rem = res % 10
		res = res / 10
		number = append(number, rem)
	}

}
