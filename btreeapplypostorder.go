package piscine

// BTreeApplyPostorder applies function f to each node's Data using postorder traversal
func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return // Base case: nothing to do
	}

	// Step 1: Traverse left subtree
	BTreeApplyPostorder(root.Left, f)

	// Step 2: Traverse right subtree
	BTreeApplyPostorder(root.Right, f)

	// Step 3: Apply function to current node's Data
	f(root.Data)
}
