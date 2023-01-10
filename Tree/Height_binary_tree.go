package Tree

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func Height(root *Node) int {
	if root == nil {
		return 0
	}
	lh := Height(root.Left)
	rh := Height(root.Right)
	max := 1 + Max(lh, rh)
	return max
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func initialize() Node {
	root := Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Left.Left = &Node{Value: 4}
	root.Right = &Node{Value: 3}
	root.Right.Left = &Node{Value: 5}
	root.Right.Right = &Node{Value: 6}
	return root
}
