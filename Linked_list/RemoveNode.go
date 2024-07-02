package Linked_list

// You are given the head of a linked list.
//
// Remove every node which has a node with a greater value anywhere to the right side of it.
//
// Return the head of the modified linked list.
func RemoveNodes(head *ListNode) *ListNode {
	temp := &ListNode{}
	temp.Val = head.Next.Val
	for head.Next != nil {
		if head.Val < head.Next.Val {
			if temp.Val == head.Next.Val {
				temp.Val = head.Next.Val
			} else {
				temp.Next = head.Next

			}
		}
		head = head.Next
	}
	return temp
}
