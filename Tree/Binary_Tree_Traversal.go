package Tree

//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Level Order Traversal of a binary tree using BFS
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	result := [][]int{}
	level := 0
	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			q := queue[0]
			queue = queue[1:]
			if len(result) <= level {
				result = append(result, []int{q.Val})
			} else {
				result[level] = append(result[level], q.Val)
			}
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		level++
	}
	return result
}
