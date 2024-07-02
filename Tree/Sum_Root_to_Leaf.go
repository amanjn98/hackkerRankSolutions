package Tree

import (
	"math"
)

// You are given the root of a binary tree containing digits from 0 to 9 only.
//
// Each root-to-leaf path in the tree represents a number.
//
// For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.
// Return the total sum of all root-to-leaf numbers. Test cases are generated so that the answer will fit in a 32-bit integer.
//
// A leaf node is a node with no children.
// Input: root = [1,2,3]
// Output: 25
// Explanation:
// The root-to-leaf path 1->2 represents the number 12.
// The root-to-leaf path 1->3 represents the number 13.
// Therefore, sum = 12 + 13 = 25.
func sumNumbers(root *Node) int {
	path := []int{}
	if root.Left == nil && root.Right == nil {
		return root.Value
	}
	return printRoottoLeaf(root, path)
}

func printRoottoLeaf(node *Node, path []int) int {
	if node == nil {
		pathSum := 0
		count := 0
		for i := len(path) - 1; i >= 0; i-- {
			pathSum += path[i] * int(math.Pow(10, float64(count)))
			count++
		}
		return pathSum
	}

	path = append(path, node.Value)
	if node.Left == nil && node.Right == nil {
		return printRoottoLeaf(node.Left, path)
	}
	if node.Right == nil {
		return printRoottoLeaf(node.Left, path)
	}
	if node.Left == nil {
		return printRoottoLeaf(node.Right, path)
	}
	sum1 := printRoottoLeaf(node.Left, path)
	sum2 := printRoottoLeaf(node.Right, path)
	return sum1 + sum2
}

// Optimised Solution

//func sumNumbers(root *TreeNode) int {
//	return printRoottoLeaf(root, 0)
//}
//
//func printRoottoLeaf(node *TreeNode, num int) int {
//	if node == nil {
//		return 0
//	}
//	if node.Left==nil && node.Right==nil{
//		return num*10+node.Val
//	}
//	return printRoottoLeaf(node.Left, num*10+node.Val)+ printRoottoLeaf(node.Right, num*10+node.Val)
//}
