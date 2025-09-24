package piscine

// SortListInsert inserts a new node into a sorted linked list
func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref} // Create the new node

	// Case 1: Insert at the beginning or into an empty list
	if l == nil || data_ref < l.Data {
		newNode.Next = l
		return newNode
	}

	// Case 2: Traverse to find the correct insertion point
	current := l
	for current.Next != nil && current.Next.Data < data_ref {
		current = current.Next
	}

	// Insert the new node after current
	newNode.Next = current.Next
	current.Next = newNode

	return l // Return the head of the list
}
