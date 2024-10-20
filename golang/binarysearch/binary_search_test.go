package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		array    []int
		expected string
	}{
		{"Found at position 4", 5, []int{1, 2, 3, 4, 5, 6}, "5 found at position 4"},
		{"Value not found", 7, []int{1, 2, 3, 4, 5, 6}, "Item 7 does not exist"},
		{"Empty array", 1, []int{}, "Item 1 does not exist"},
		{"Single element, found", 1, []int{1}, "1 found at position 0"},
		{"Single element, not found", 2, []int{1}, "Item 2 does not exist"},
		{"First element", 1, []int{1, 2, 3, 4, 5, 6}, "1 found at position 0"},
		{"Last element", 6, []int{1, 2, 3, 4, 5, 6}, "6 found at position 5"},
	}

	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t,tt.expected,search(tt.value,tt.array))
		})
	}
}
