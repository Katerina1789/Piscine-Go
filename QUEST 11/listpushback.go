package piscine // Defines the package name

// NodeL represents a single element in the linked list
type NodeL1 struct {
	Data interface{} // Holds any type of data (string, int, etc.)
	Next *NodeL      // Pointer to the next node in the list (nil if it's the last)
}

// List represents the linked list itself, with pointers to the first and last nodes
type List1 struct {
	Head *NodeL // Points to the first node in the list
	Tail *NodeL // Points to the last node in the list
}

// ListPushBack adds a new node with the given data to the end of the list
func ListPushBack(l *List1, data interface{}) {
	// Create a new node with the provided data
	newNode := &NodeL{Data: data}

	// If the list is empty (no Head), set both Head and Tail to the new node
	if l.Head == nil {
		l.Head = newNode // First node becomes the Head
		l.Tail = newNode // And also the Tail since it's the only node
	} else {
		// If the list is not empty, link the current Tail to the new node
		l.Tail.Next = newNode // Set the current Tail's Next to the new node
		l.Tail = newNode      // Update Tail to point to the new last node
	}
}
