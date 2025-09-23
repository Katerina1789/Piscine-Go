package piscine

type NodeL struct {
	Data interface{} // Holds any type of value (e.g., string, int)
	Next *NodeL      // Points to the next node in the list
}

type List struct {
	Head *NodeL // First node in the list
	Tail *NodeL // Last node in the list
}

// CompStr compares two interface{} values for equality
func CompStr(a, b interface{}) bool {
	return a == b
}

// ListFind searches for the first node whose Data matches ref using comp,
// and returns a pointer to that Data (i.e., *interface{})
func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	current := l.Head // Start from the head of the list

	for current != nil {
		// Use the comparison function to check if current Data matches ref
		if comp(current.Data, ref) {
			// Return the address of the Data field
			return &current.Data
		}
		// Move to the next node
		current = current.Next
	}

	// If no match is found, return nil
	return nil
}
