package piscine

type NodeL4 struct {
	Data interface{} // Holds the value
	Next *NodeL      // Points to the next node
}

type List4 struct {
	Head *NodeL // First node in the list
	Tail *NodeL // Last node in the list
}

// ListLast returns the data of the last node in the list
func ListLast(l *List) interface{} {
	if l.Tail == nil {
		return nil // If the list is empty, return nil
	} else {
		return l.Tail.Data // Otherwise, return the data of the last node
	}
}
