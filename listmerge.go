package piscine

type NodeL struct {
	Data interface{} // Holds any type of value
	Next *NodeL      // Points to the next node
}

type List struct {
	Head *NodeL // First node in the list
	Tail *NodeL // Last node in the list
}

// ListMerge appends all nodes from l2 to the end of l1 without creating new nodes
func ListMerge(l1 *List, l2 *List) {
	// If l2 is empty, there's nothing to merge
	if l2.Head == nil {
		return
	}

	// If l1 is empty, just point l1 to l2's head and tail
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	}

	// Otherwise, link l1's tail to l2's head
	l1.Tail.Next = l2.Head
	l1.Tail = l2.Tail
}
