package Tree

// Question
//Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths where the sum of the node values
//in the path equals targetSum. Each path should be returned as a list of the node values, not node references.
//A root-to-leaf path is a path starting from the root and ending at any leaf node. A leaf is a node with no children.
//Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//Output: [[5,4,11,2],[5,8,4,5]]
//Explanation: There are two paths whose sum equals targetSum:
//5 + 4 + 11 + 2 = 22
//5 + 8 + 4 + 5 = 22

// pathSum is a function that takes a root node of a binary tree and a target sum as parameters and returns a 2D slice
// containing all root-to-leaf paths that sum to the target.
// It initializes an empty slice called path to store the current path and a 2D slice called result to store all valid paths.
// If the root node is nil, indicating an empty tree, the function returns an empty 2D slice.
// Otherwise, it calls the helper function printPaths to recursively find and append the valid paths to the result.
// Finally, it returns the result slice.
func pathSum(root *TreeNode, targetSum int) [][]int {
	path := []int{}
	result := [][]int{}
	if root == nil {
		return [][]int{}
	}
	printPaths(root, path, targetSum, &result)
	return result
}

func printPaths(node *TreeNode, path []int, targetSum int, result *[][]int) {
	if node == nil {
		return
	}

	path = append(path, node.Val)
	pathSum := 0
	for i := 0; i < len(path); i++ {
		pathSum += path[i]
	}
	if pathSum == targetSum && node.Left == nil && node.Right == nil {
		*result = append(*result, append([]int{}, path...))

	}
	printPaths(node.Left, path, targetSum, result)
	printPaths(node.Right, path, targetSum, result)
}
