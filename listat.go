package piscine

type NodeL6 struct {
	Data interface{}
	Next *NodeL
}

type List6 struct {
	Head *NodeL
	Tail *NodeL
}

// ListAt returns the node at position pos in the list, or nil if out of bounds
func ListAt(l *NodeL, pos int) *NodeL {
	current := l // Start from the head
	index := 0   // Track current position

	for current != nil {
		if index == pos {
			return current // Found the node at the desired position
		}
		current = current.Next // Move to the next node
		index++
	}

	return nil // If position is out of bounds, return nil
}
