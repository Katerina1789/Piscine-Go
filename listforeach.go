package piscine

type NodeL struct {
	Data interface{} // Holds the value
	Next *NodeL      // Points to the next node
}

type List struct {
	Head *NodeL // First node in the list
	Tail *NodeL // Last node in the list
}

// ListForEach applies the given function f to each node in the list
func ListForEach(l *List, f func(*NodeL)) {
	current := l.Head // Start from the head

	for current != nil {
		f(current)             // Apply the function to the current node
		current = current.Next // Move to the next node
	}
}

// Add2_node adds 2 to int data or appends "2" to string data
func Add2_node(node *NodeL) {
	switch v := node.Data.(type) {
	case int:
		node.Data = v + 2
	case string:
		node.Data = v + "2"
	}
}

// Subtract3_node subtracts 3 from int data or appends "-3" to string data
func Subtract3_node(node *NodeL) {
	switch v := node.Data.(type) {
	case int:
		node.Data = v - 3
	case string:
		node.Data = v + "-3"
	}
}
