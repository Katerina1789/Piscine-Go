package piscine

// BTreeIsBinary checks if the tree follows BST properties
func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, nil, nil)
}

// Helper function with min/max bounds
func isBST(node *TreeNode, min *string, max *string) bool {
	if node == nil {
		return true // Empty subtree is valid
	}

	// Check current node against bounds
	if (min != nil && node.Data <= *min) || (max != nil && node.Data > *max) {
		return false
	}

	// Recursively check left and right subtrees
	return isBST(node.Left, min, &node.Data) &&
		isBST(node.Right, &node.Data, max)
}
