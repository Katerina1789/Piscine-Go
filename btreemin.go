package piscine

// BTreeMin returns the node with the minimum value in the tree
func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil // Empty tree
	}

	current := root
	for current.Left != nil {
		current = current.Left // Keep going left
	}

	return current // Leftmost node is the min
}
