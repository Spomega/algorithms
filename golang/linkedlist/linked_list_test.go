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

func TestInsertAtEnd(t *testing.T) {
	list := &LinkedList{}
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)

	expected := []int{10, 20, 30}
	if got := list.ToSlice(); !equalSlices(got, expected) {
		t.Errorf("InsertAtEnd() = %v, want %v", got, expected)
	}
}

func TestInsertAtPosition(t *testing.T) {
	list := &LinkedList{}
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)

	// Insert in the middle
	if err := list.InsertAtPosition(15, 2); err != nil {
		t.Fatalf("InsertAtPosition() unexpected error: %v", err)
	}

	// Insert at the beginning
	if err := list.InsertAtPosition(5, 1); err != nil {
		t.Fatalf("InsertAtPosition() unexpected error: %v", err)
	}

	// Insert at the end
	if err := list.InsertAtPosition(35, 6); err != nil {
		t.Fatalf("InsertAtPosition() unexpected error: %v", err)
	}

	expected := []int{5, 10, 15, 20, 30, 35}
	if got := list.ToSlice(); !equalSlices(got, expected) {
		t.Errorf("InsertAtPosition() = %v, want %v", got, expected)
	}
}

func TestDelete(t *testing.T) {
	list := &LinkedList{}
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	list.InsertAtEnd(40)

	// Delete a middle element
	if err := list.Delete(20); err != nil {
		t.Fatalf("Delete() unexpected error: %v", err)
	}

	expected := []int{10, 30, 40}
	if got := list.ToSlice(); !equalSlices(got, expected) {
		t.Errorf("Delete() = %v, want %v", got, expected)
	}

	// Delete head element
	if err := list.Delete(10); err != nil {
		t.Fatalf("Delete() unexpected error: %v", err)
	}

	expected = []int{30, 40}
	if got := list.ToSlice(); !equalSlices(got, expected) {
		t.Errorf("Delete() = %v, want %v", got, expected)
	}

	// Delete tail element
	if err := list.Delete(40); err != nil {
		t.Fatalf("Delete() unexpected error: %v", err)
	}

	expected = []int{30}
	if got := list.ToSlice(); !equalSlices(got, expected) {
		t.Errorf("Delete() = %v, want %v", got, expected)
	}

	// Delete non-existent element
	if err := list.Delete(50); err == nil {
		t.Fatalf("Delete() expected error for non-existent element")
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
