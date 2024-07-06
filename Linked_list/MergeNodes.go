package Linked_list

//You are given the head of a linked list, which contains a series of integers separated by 0's. The beginning and end of the linked list will have Node.val == 0.
//
//For every two consecutive 0's, merge all the nodes lying in between them into a single node whose value is the sum of all the merged nodes. The modified list should not contain any 0's.
//
//Return the head of the modified linked list.

//Example 1:
//
//
//Input: head = [0,3,1,0,4,5,2,0]
//Output: [4,11]
//Explanation:
//The above figure represents the given linked list. The modified list contains
//- The sum of the nodes marked in green: 3 + 1 = 4.
//- The sum of the nodes marked in red: 4 + 5 + 2 = 11.
//Example 2:
//
//
//Input: head = [0,1,0,3,0,2,2,0]
//Output: [1,3,4]
//Explanation:
//The above figure represents the given linked list. The modified list contains
//- The sum of the nodes marked in green: 1 = 1.
//- The sum of the nodes marked in red: 3 = 3.
//- The sum of the nodes marked in yellow: 2 + 2 = 4.

// curr linked list is used to traverse through the linkedlist
// prev is used to keep track of the modified linkedlist and update it accordingly
func MergeNodes(list1 *ListNode) *ListNode {
	curr := list1.Next
	sum := 0
	prev := list1
	for curr != nil {
		if curr.Val != 0 {
			sum += curr.Val
		} else {
			curr.Val = sum
			prev.Next = curr
			prev = prev.Next
			sum = 0
		}
		curr = curr.Next
	}
	return list1.Next
}

func Insert(head **ListNode, value int) {
	newNode := &ListNode{Val: value}
	if *head == nil {
		*head = newNode
	} else {
		current := *head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}
