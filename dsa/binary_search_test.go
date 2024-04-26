package dsa

import "testing"


func TestIntArrBinarySearchFound(t *testing.T) {
    var intArr []int = []int{1, 3, 4, 5, 6, 10, 11}
    pos, found := BinarySearch(intArr, 6)
    if pos != 4 || found != true {
        t.Fatalf("Binary Search failed for Int Arr\nPosition:\n Ex: %d Got: %d\nFound:\n Ex: %v Got: %v", 4, pos, true, found)
    }
}

func TestIntArrBinarySearchNotFoundA(t *testing.T) {
    var intArr []int = []int{1, 3, 4, 5, 7, 10, 11}
    pos, found := BinarySearch(intArr, 8)
    if pos != 5 || found != false {
        t.Fatalf("Binary Search failed for Int Arr\nPosition:\n Ex: %d Got: %d\nFound:\n Ex: %v Got: %v", 5, pos, false, found)
    }
}

func TestIntArrBinarySearchNotFoundB(t *testing.T) {
    var intArr []int = []int{1, 2, 5, 11}
    pos, found := BinarySearch(intArr, 3)
    if pos != 2 || found != false {
        t.Fatalf("Binary Search failed for Int Arr\nPosition:\n Ex: %d Got: %d\nFound:\n Ex: %v Got: %v", 2, pos, false, found)
    }
}

func TestRuneArrBinarySearchFound(t *testing.T) {
    var runeArr []rune = []rune{'a', 'c', 'd', 'x', 'z'}
    pos, found := BinarySearch(runeArr, 'd')
    if pos != 2 || found != true {
        t.Fatalf("Binary Search failed for Rune Arr\nPosition:\n Ex: %d Got: %d\nFound:\n Ex: %v Got: %v", 2, pos, true, found)
    }
}

func TestRuneArrBinarySearchNotFound(t *testing.T) {
    var runeArr []rune = []rune{'a', 'c', 'd', 'x', 'z'}
    pos, found := BinarySearch(runeArr, 'b')
    if pos != 1 || found != false {
        t.Fatalf("Binary Search failed for Rune Arr\nPosition:\n Ex: %d Got: %d\nFound:\n Ex: %v Got: %v", 1, pos, false, found)
    }
}

type CustomComparableType struct {
    key string
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
            key: "a",
            value: 1,
        },
        CustomComparableType{
            key: "b",
            value: 2,
        },
        CustomComparableType{
            key: "c",
            value: 5,
        },
        CustomComparableType{
            key: "d",
            value: 11,
        },
    }
    pos, found := BinarySearchCustom(customArr, CustomComparableType{key: "b", value: 2})
    if pos != 1 || found != true {
        t.Fatalf("Binary Search failed for CustomComparableType Arr\nPosition:\n Ex: %d Got: %d\nFound:\n Ex: %v Got: %v", 1, pos, true, found)
    }
}

func TestCustomComparableArrBinarySearchNotFound(t *testing.T) {
    var customArr SortableCustomType = SortableCustomType{
        CustomComparableType{
            key: "a",
            value: 1,
        },
        CustomComparableType{
            key: "b",
            value: 2,
        },
        CustomComparableType{
            key: "c",
            value: 5,
        },
        CustomComparableType{
            key: "d",
            value: 11,
        },
    }
    pos, found := BinarySearchCustom(customArr, CustomComparableType{key: "b", value: 3})
    if pos != 2 || found != false {
        t.Fatalf("Binary Search failed for CustomComparableType Arr\nPosition:\n Ex: %d Got: %d\nFound:\n Ex: %v Got: %v", 2, pos, false, found)
    }
}

