package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeInsertData inserts a new node into the binary search tree
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	// If the tree is empty, create and return the new root
	if root == nil {
		return &TreeNode{Data: data}
	}

	// If the new data is less than current node, go left
	if data < root.Data {
		if root.Left == nil {
			// Insert here
			root.Left = &TreeNode{Data: data, Parent: root}
		} else {
			// Recurse into left subtree
			BTreeInsertData(root.Left, data)
		}
	} else {
		// If the new data is greater or equal, go right
		if root.Right == nil {
			// Insert here
			root.Right = &TreeNode{Data: data, Parent: root}
		} else {
			// Recurse into right subtree
			BTreeInsertData(root.Right, data)
		}
	}

	return root // Return the unchanged root
}
