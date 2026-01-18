package piscine

type NodeL5 struct {
	Data interface{} // Holds the value
	Next *NodeL      // Points to the next node
}

type List5 struct {
	Head *NodeL // First node in the list
	Tail *NodeL // Last node in the list
}

// ListClear deletes all nodes from the list
func ListClear(l *List5) {
	l.Head = nil // Remove reference to the first node
	l.Tail = nil // Remove reference to the last node
}
