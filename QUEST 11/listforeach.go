package piscine

type NodeL8 struct {
	Data interface{} // Holds the value
	Next *NodeL      // Points to the next node
}

type List8 struct {
	Head *NodeL8 // First node in the list
	Tail *NodeL8 // Last node in the list
}

// ListForEach applies the given function f to each node in the list
func ListForEach(l *List, f func(*NodeL)) {
	curren := l.Head // Start from the head

	for curren != nil {
		f(curren)            // Apply the function to the current node
		curren = curren.Next // Move to the next node
	}
}

// Add2_node adds 2 to int data or appends "2" to string data
func Add2_node(node *NodeL8) {
	switch v := node.Data.(type) {
	case int:
		node.Data = v + 2
	case string:
		node.Data = v + "2"
	}
}

// Subtract3_node subtracts 3 from int data or appends "-3" to string data
func Subtract3_node(node *NodeL8) {
	switch v := node.Data.(type) {
	case int:
		node.Data = v - 3
	case string:
		node.Data = v + "-3"
	}
}
