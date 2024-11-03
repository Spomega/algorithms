package main

import "testing"

// Helper function to convert linked list to a slice for easy comparison in tests
func (list *Linkedlist) ToSlice() []int {
	var elements []int
	current := list.head

	for current != nil {
		elements = append(elements, current.data)
		current = current.next
	}

	return elements
}

func TestInsertAtBeginning(t *testing.T) {
	list := &Linkedlist{}

	list.InsertAtBeginning(10)
	list.InsertAtBeginning(20)
	list.InsertAtBeginning(30)

	expected := []int{30, 20, 10}
	if got := list.ToSlice(); !equalSlices(got, expected) {
		t.Errorf("InsertAtBeginning() = %v, want %v", got, expected)
	}
}

// Helper function to check if two slices are equal
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
