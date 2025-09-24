package piscine

// SortedListMerge merges two sorted linked lists into one sorted list
func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	// If either list is empty, return the other
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}

	var head *NodeI    // Head of the merged list
	var current *NodeI // Pointer to build the merged list

	// Initialize head with the smaller first node
	if n1.Data < n2.Data {
		head = n1
		n1 = n1.Next
	} else {
		head = n2
		n2 = n2.Next
	}
	current = head

	// Merge nodes in sorted order
	for n1 != nil && n2 != nil {
		if n1.Data < n2.Data {
			current.Next = n1
			n1 = n1.Next
		} else {
			current.Next = n2
			n2 = n2.Next
		}
		current = current.Next
	}

	// Attach remaining nodes from either list
	if n1 != nil {
		current.Next = n1
	} else {
		current.Next = n2
	}

	return head
}
