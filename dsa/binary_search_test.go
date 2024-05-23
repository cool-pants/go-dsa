package dsa

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntArrBinarySearchFound(t *testing.T) {
	var intArr []int = []int{1, 3, 4, 5, 7, 10, 11}
	pos, found := BinarySearch(intArr, 7)
	pos2, found2 := BinarySearch(intArr, 6)
	pos3, found3 := BinarySearch(intArr, 8)
	assert.Equal(t, 4, pos)
	assert.Equal(t, true, found)
	assert.Equal(t, 4, pos2)
	assert.Equal(t, false, found2)
	assert.Equal(t, 5, pos3)
	assert.Equal(t, false, found3)
}

func TestRuneArrBinarySearchFound(t *testing.T) {
	var runeArr []rune = []rune{'a', 'c', 'd', 'x', 'z'}
	pos, found := BinarySearch(runeArr, 'd')
	pos2, found2 := BinarySearch(runeArr, 'b')

	assert.Equal(t, 2, pos)
	assert.Equal(t, true, found)
	assert.Equal(t, 1, pos2)
	assert.Equal(t, false, found2)
}

type CustomComparableType struct {
	key   string
	value int64
}

func (c CustomComparableType) Equals(l Comparable) bool {
	return c.value == l.(CustomComparableType).value && c.key == l.(CustomComparableType).key
}

func (c CustomComparableType) Less(l Comparable) bool {
	return c.value < l.(CustomComparableType).value
}

type SortableCustomType []CustomComparableType

func (s SortableCustomType) Less(i, j int) bool {
	return s[i].Less(s[j])
}

func (s SortableCustomType) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortableCustomType) Len() int {
	return len(s)
}

func TestCustomComparableArrBinarySearchFound(t *testing.T) {
	var customArr SortableCustomType = SortableCustomType{
		CustomComparableType{
			key:   "a",
			value: 1,
		},
		CustomComparableType{
			key:   "b",
			value: 2,
		},
		CustomComparableType{
			key:   "c",
			value: 5,
		},
		CustomComparableType{
			key:   "d",
			value: 11,
		},
	}
	pos, found := BinarySearchCustom(customArr, CustomComparableType{key: "b", value: 2})
	pos2, found2 := BinarySearchCustom(customArr, CustomComparableType{key: "b", value: 3})

	assert.Equal(t, 1, pos)
	assert.Equal(t, true, found)
	assert.Equal(t, 2, pos2)
	assert.Equal(t, false, found2)
}

func TestCustomComparableArrBinarySearchNotFound(t *testing.T) {
	var customArr SortableCustomType = SortableCustomType{
		CustomComparableType{
			key:   "a",
			value: 1,
		},
		CustomComparableType{
			key:   "b",
			value: 2,
		},
		CustomComparableType{
			key:   "c",
			value: 5,
		},
		CustomComparableType{
			key:   "d",
			value: 11,
		},
	}
	pos, found := BinarySearchCustom(customArr, CustomComparableType{key: "b", value: 3})
	if pos != 2 || found != false {
		t.Fatalf("Binary Search failed for CustomComparableType Arr\nPosition:\n Ex: %d Got: %d\nFound:\n Ex: %v Got: %v", 2, pos, false, found)
	}
}
