package piscine

// BTreeSearchItem searches for a node with Data equal to elem
func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil // Base case: not found
	}

	if root.Data == elem {
		return root // Found the node
	}

	if elem < root.Data {
		// Search in the left subtree
		return BTreeSearchItem(root.Left, elem)
	} else {
		// Search in the right subtree
		return BTreeSearchItem(root.Right, elem)
	}
}
