package piscine

type NodeL struct {
	Data interface{} // Holds the value
	Next *NodeL      // Points to the next node
}

type List struct {
	Head *NodeL // First node in the list
	Tail *NodeL // Last node in the list
}

// ListReverse reverses the order of nodes in the linked list
func ListReverse(l *List) {
	var prev *NodeL   // Will point to the previous node
	current := l.Head // Start from the head
	l.Tail = l.Head   // After reversal, the old head becomes the new tail

	for current != nil {
		next := current.Next // Save the next node
		current.Next = prev  // Reverse the link
		prev = current       // Move prev forward
		current = next       // Move current forward
	}

	l.Head = prev // After loop, prev is the new head
}
