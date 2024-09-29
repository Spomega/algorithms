package reverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseInt(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	s = ReverseInt(s)
	assert.Equal(t, []int{5, 4, 3, 2, 1}, s)
}

func TestReverseStr(t *testing.T) {
	s := "Hello"
	assert.Equal(t, "olleH", ReverseStr(s))
}
