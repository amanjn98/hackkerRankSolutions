package Tree

// Given the root of a binary tree, return the sum of all left leaves. A leaf is a node with no children.
//A left leaf is a leaf that is the left child of another node.

//Input: root = [3,9,20,null,null,15,7]
//Output: 24
//Explanation: There are two left leaves in the binary tree, with values 9 and 15 respectively.

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	return PrintSum(root, sum)
}

func PrintSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}
	sum += PrintSum(root.Left, 0)
	sum += PrintSum(root.Right, 0)
	return sum
}
