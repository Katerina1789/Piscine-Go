package piscine

type NodeL9 struct {
	Data interface{} // Holds any type of value (int, string, etc.)
	Next *NodeL      // Points to the next node in the list
}

type List9 struct {
	Head *NodeL // Points to the first node in the list
	Tail *NodeL // Points to the last node in the list
}

// IsPositiveNode returns true if the node's Data is a positive number
func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		// Type assertion: convert Data to int and check if it's > 0
		return node.Data.(int) > 0
	default:
		// If Data is not a number, return false
		return false
	}
}

// IsAlNode returns true if the node's Data is NOT a number (e.g., a string)
func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		// If Data is a number, return false
		return false
	default:
		// If Data is not a number, return true
		return true
	}
}

// ListForEachIf applies function f to each node in the list l,
// but only if the condition function cond returns true for that node.
func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	current := l.Head // Start from the head of the list

	// Traverse the list node by node
	for current != nil {
		// Check if the current node satisfies the condition
		if cond(current) {
			// If it does, apply the function f to this node
			f(current)
		}
		// Move to the next node in the list
		current = current.Next
	}
}
