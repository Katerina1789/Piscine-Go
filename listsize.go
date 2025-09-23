package piscine

type NodeL struct {
	Data interface{} // Holds the value
	Next *NodeL      // Points to the next node
}

type List struct {
	Head *NodeL // First node in the list
	Tail *NodeL // Last node in the list
}

// ListSize returns the number of nodes in the linked list
func ListSize(l *List) int {
	count := 0        // Start with zero
	current := l.Head // Begin at the head of the list

	for current != nil { // Traverse until the end
		count++                // Count this node
		current = current.Next // Move to the next node
	}

	return count // Return the total count
}
