package main

import (
	"errors"
	"fmt"
)

// node represents a single node in the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents the linked list structure
type Linkedlist struct {
	head *Node
}

// InsertAtBeginning adds a new node at the beginning on the list
func (list *Linkedlist) InsertAtBeginning(data int) {
	fmt.Printf("%v\n", list.head)
	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}

// InsertAtEnd adds a new node at the end of the list
func (list *Linkedlist) InsertAtEnd(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// InsertAtPosition adds a new node at the given position (1-based index)
func (list *Linkedlist) InsertAtPosition(data int, position int) error {
	if position < 1 {
		return errors.New("position must be greater than or equal to 1")
	}

	newNode := &Node{data: data}

	if position == 1 {
		newNode.next = list.head
		list.head = newNode
		return nil
	}

	current := list.head
	for i := 1; i < position-1; i++ {
		if current == nil {
			return errors.New("position out of bounds")
		}
		current = current.next
	}

	if current == nil {
		return errors.New("position out of bounds")
	}

	newNode.next = current.next
	current.next = newNode
	return nil
}

// Delete deletes the first node with the specified data
func (list *Linkedlist) Delete(data int) error {
	if list.head == nil {
		return errors.New("list is empty")
	}

	if list.head.data == data {
		list.head = list.head.next
		return nil
	}

	current := list.head
	for current.next != nil && current.next.data != data {
		current = current.next
	}

	if current.next == nil {
		return errors.New("elements not found")
	}

	current.next = current.next.next
	return nil
}

// Display prints all elements of the linked list
func (list *Linkedlist) Display() {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := &Linkedlist{}

	// Insert at the beginning
	list.InsertAtBeginning(1)
	list.InsertAtBeginning(2)
	list.InsertAtBeginning(3)
	fmt.Println("After inserting at the beginning:")
	list.Display()

	// Insert at the end
	list.InsertAtEnd(4)
	list.InsertAtEnd(5)
	fmt.Println("After inserting at the end:")
	list.Display()

	// // Insert at a specific position
	// err := list.InsertAtPosition(6, 3)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("After inserting at position 3:")
	// 	list.Display()
	// }

	// // Delete a node
	// err = list.Delete(2)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("After deleting 2:")
	// 	list.Display()
	// }
}
