package piscine

// BTreeApplyPreorder applies function f to each node's Data using preorder traversal
func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return // Base case: nothing to do
	}

	// Step 1: Apply function to current node's Data
	f(root.Data)

	// Step 2: Traverse left subtree
	BTreeApplyPreorder(root.Left, f)

	// Step 3: Traverse right subtree
	BTreeApplyPreorder(root.Right, f)
}
