package piscine

// NodeL represents a single node in the linked list
type NodL struct {
	Data interface{} // Holds the value (can be any type)
	Next *NodeL      // Points to the next node in the list
}

// List represents the linked list with pointers to the first and last nodes
type Lit struct {
	Head *NodeL // First node in the list
	Tail *NodeL // Last node in the list
}

// ListPushFront inserts a new node at the beginning of the list
func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data} // Create a new node with the given data

	newNode.Next = l.Head // Link the new node to the current head
	l.Head = newNode      // Update the head to be the new node

	if l.Tail == nil {
		// If the list was empty, set Tail to the new node too
		l.Tail = newNode
	}
}
