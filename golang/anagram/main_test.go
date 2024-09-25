package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagramMap_Success(t *testing.T) {

	first := "The eyes"
	second := "they see"

	if !AnagramMapApproach(first, second) {
		t.Fatal("the two strings are not anagrams")
	}

	assert.True(t, AnagramMapApproach(first, second))

}

func TestAnagramMap_Error(t *testing.T) {
	first := "The eyes!"
	second := "they see"

	assert.False(t, AnagramMapApproach(first, second))
}


func TestAnagramSort_Success(t *testing.T) {
	first := "The eyes"
	second := "they see"

	assert.True(t, AnagramsortApproach(first, second))

}

func TestAnagramSort_Error(t *testing.T) {
	first := "The eyes!"
	second := "they see"

	assert.False(t,AnagramsortApproach(first, second))
}
