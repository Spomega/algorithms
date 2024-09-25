package main

import (
	"fmt"
	"sort"
	"strings"
	//"reflect"
)

func main() {
	first := "hello World"
	second := "lleho world"

	fmt.Printf("%v\n", AnagramMapApproach(first, second))
	fmt.Printf("%v", AnagramsortApproach(first, second))

}

// Uses Character Map approach
func AnagramMapApproach(first string, second string) bool {

	if len(first) != len(second) {
		return false
	}

	first = removeSpaces(first)
	second = removeSpaces(second)

	//reflect.DeepEqual(m,n)

	m := buildCharMap(first)
	n := buildCharMap(second)

	for k, mv := range m {
		if nv, ok := n[k]; !ok || mv != nv {
			return false
		}
	}
	return true
}

// uses split sort approach
func AnagramsortApproach(first string, second string) bool {
	// Remove any spaces and convert to lowercase
	first = removeSpaces(first)
	second = removeSpaces(second)

	// If the lengths don't match, they can't be anagrams
	if len(first) != len(second) {
		return false
	}

	// Convert strings to rune slices (arrays of characters)
	firstRune := []rune(first)
	secondRune := []rune(second)

	// Sort both rune slices
	sortSlices(firstRune)
	sortSlices(secondRune)

	return string(firstRune) == string(secondRune)

}

func removeSpaces(value string) string {
	return strings.ToLower(strings.ReplaceAll(value, " ", ""))
}

func sortSlices(slice []rune) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

func buildCharMap(value string) map[rune]int {
	m := make(map[rune]int)
	for _, v := range value {
		m[v]++
	}
	return m
}
