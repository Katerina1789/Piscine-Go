package piscine

// BTreeApplyInorder applies function f to each node's Data in ascending order
func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return // Base case: nothing to do
	}

	// Step 1: Traverse left subtree
	BTreeApplyInorder(root.Left, f)

	// Step 2: Apply function to current node's Data
	f(root.Data)

	// Step 3: Traverse right subtree
	BTreeApplyInorder(root.Right, f)
}
