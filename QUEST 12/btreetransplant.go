package piscine

// BTreeTransplant replaces subtree rooted at node with subtree rooted at rplc
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == nil {
		return root // Nothing to replace
	}

	// Case 1: node is the root
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		// Case 2: node is a left child
		node.Parent.Left = rplc
	} else {
		// Case 3: node is a right child
		node.Parent.Right = rplc
	}

	// Update rplc's parent if it's not nil
	if rplc != nil {
		rplc.Parent = node.Parent
	}

	return root
}
