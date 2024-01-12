package Tree

import "fmt"

//Given the root node of a binary tree and an integer targetSum, we want to print all paths where the sum of the values
//along each path equals targetSum. The path does not need to start at the root node or end with a leaf node.
//However, it must go downwards. That is, we traverse the path from parent nodes to child nodes.
// Also, the path can be a single node whose data value is targetSum. For example, we have 4 paths that satisfy the
//targetSum of 7 in the following binary tree:

func PrintPaths(root *Node, targetSum int) {
	path := []int{}
	printPathsRec(root, path, targetSum)
}

func printPathsRec(node *Node, path []int, targetSum int) {
	if node == nil {
		return
	}

	path = append(path, node.Value)
	pathSum := 0
	for i := len(path) - 1; i >= 0; i-- {
		pathSum += path[i]
		if pathSum == targetSum {
			printPath(path, i)
		}
	}

	printPathsRec(node.Left, path, targetSum)
	printPathsRec(node.Right, path, targetSum)
}

func printPath(path []int, startIdx int) {
	for i := startIdx; i < len(path); i++ {
		fmt.Print(path[i], " ")
	}
	fmt.Println()
}
