package Recursion

/**
* Definition for a Node.
 */
type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	ipArray := []int{}
	PostOrderTraversal(root, &ipArray)
	return ipArray
}

func PostOrderTraversal(root *Node, ipArray *[]int) {
	if root == nil {
		return
	}
	for i := 0; i < len(root.Children); i++ {
		PostOrderTraversal(root.Children[i], ipArray)
	}
	*ipArray = append(*ipArray, root.Val)
	return
}
