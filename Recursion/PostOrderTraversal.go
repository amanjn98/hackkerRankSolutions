package Recursion

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	ipArray := []int{}
	PrintPostOrderTraversal(root, &ipArray)
	return ipArray
}

func PrintPostOrderTraversal(root *TreeNode, ipArray *[]int) {
	if root == nil {
		return
	}
	PrintPostOrderTraversal(root.Left, ipArray)
	PrintPostOrderTraversal(root.Right, ipArray)
	*ipArray = append(*ipArray, root.Val)
	return
}
