package piscine

// BTreeLevelCount returns the number of levels (height) of the tree
func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0 // No levels in an empty tree
	}

	// Recursively get the height of left and right subtrees
	leftHeight := BTreeLevelCount(root.Left)
	rightHeight := BTreeLevelCount(root.Right)

	// Add 1 for the current level and return the max of both sides
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}
