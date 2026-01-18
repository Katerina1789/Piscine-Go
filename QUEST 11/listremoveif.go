package piscine

type NodL struct {
	Data interface{} // Holds any type of value
	Next *NodeL      // Points to the next node
}

type Lit struct {
	Head *NodeL // First node in the list
	Tail *NodeL // Last node in the list
}

// ListRemoveIf removes all nodes whose Data equals data_ref
func ListRemoveIf(l *List, data_ref interface{}) {
	// Remove matching nodes from the beginning of the list
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next // Move head forward
	}

	// If the list is now empty, reset Tail too
	if l.Head == nil {
		l.Tail = nil
		return
	}

	// Start from the first non-matching node
	current := l.Head

	// Traverse the list and remove matching nodes
	for current.Next != nil {
		if current.Next.Data == data_ref {
			// Skip the matching node
			current.Next = current.Next.Next
		} else {
			// Move to the next node
			current = current.Next
		}
	}

	// Update Tail: walk from Head to find the last node
	l.Tail = l.Head
	for l.Tail.Next != nil {
		l.Tail = l.Tail.Next
	}
}
