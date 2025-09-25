package piscine

// BTreeDeleteNode deletes a node from the BST and returns the new root
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	// Case 1: node has no left child
	if node.Left == nil {
		root = BTreeTransplant(root, node, node.Right)

		// Case 2: node has no right child
	} else if node.Right == nil {
		root = BTreeTransplant(root, node, node.Left)

		// Case 3: node has two children
	} else {
		// Find the in-order successor (minimum in right subtree)
		successor := BTreeMin(node.Right)

		if successor.Parent != node {
			// Replace successor with its right child
			root = BTreeTransplant(root, successor, successor.Right)
			successor.Right = node.Right
			if successor.Right != nil {
				successor.Right.Parent = successor
			}
		}

		// Replace node with successor
		root = BTreeTransplant(root, node, successor)
		successor.Left = node.Left
		if successor.Left != nil {
			successor.Left.Parent = successor
		}
	}

	return root
}
