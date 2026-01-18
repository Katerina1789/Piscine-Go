package piscine

// BTreeMax returns the node with the maximum value in the tree
func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil // Empty tree
	}

	current := root
	for current.Right != nil {
		current = current.Right // Keep going right
	}

	return current // Rightmost node is the max
}
