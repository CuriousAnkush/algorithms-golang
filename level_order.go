package algorithms

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var queue []*TreeNode

	queue = append(queue, root)

	for queueIDx := 0; queueIDx < len(queue); queueIDx++ {
		var level []int
		queueSize := len(queue)
		for queueSize > 0 {
			node := queue[queueIDx]
			level = append(level, node.Val)
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
			queueIDx++
			queueSize--
		}
		result = append(result, level)
	}
	return result
}
