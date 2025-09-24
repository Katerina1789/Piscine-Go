package piscine

func TreeInsert(root *TreeNode, data int) *TreeNode {
	if root == nil {
		// If the tree is empty, create a new node
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		// Insert into the left subtree
		root.Left = TreeInsert(root.Left, data)
	} else {
		// Insert into the right subtree
		root.Right = TreeInsert(root.Right, data)
	}

	return root // Return the unchanged root
}

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}
