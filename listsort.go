package piscine

type NodeI struct {
	Data int    // Holds the integer value
	Next *NodeI // Points to the next node
}

// ListSort sorts the linked list in ascending order using bubble sort
func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil // Empty list, nothing to sort
	}

	swapped := true // Flag to track if any swaps occurred

	// Keep looping until no swaps are needed
	for swapped {
		swapped = false
		current := l

		// Traverse the list and compare adjacent nodes
		for current.Next != nil {
			if current.Data > current.Next.Data {
				// Swap the data values
				current.Data, current.Next.Data = current.Next.Data, current.Data
				swapped = true // A swap occurred, so we need another pass
			}
			current = current.Next // Move to the next node
		}
	}

	return l // Return the sorted list
}
