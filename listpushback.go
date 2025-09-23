package piscine // Defines the package name

// NodeL represents a single element in the linked list
type NodeL struct {
	Data interface{} // Holds any type of data (string, int, etc.)
	Next *NodeL      // Pointer to the next node in the list (nil if it's the last)
}

// List represents the linked list itself, with pointers to the first and last nodes
type List struct {
	Head *NodeL // Points to the first node in the list
	Tail *NodeL // Points to the last node in the list
}

// ListPushBack adds a new node with the given data to the end of the list
func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data} // Create the new node

	if l.Head == nil {
		l.Head = newNode // First node in the list
		l.Tail = newNode
		return
	}

	l.Tail.Next = newNode // Link new node after Tail
	l.Tail = newNode      // Update Tail to the new node
}
