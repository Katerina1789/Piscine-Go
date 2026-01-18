package piscine

// BTreeApplyByLevel applies function f to each node's Data in level-order
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Create a queue to hold nodes
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// Pop the first node from the queue
		current := queue[0]
		queue = queue[1:]

		// Apply the function to the current node's Data
		f(current.Data)

		// Enqueue left and right children if they exist
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}
