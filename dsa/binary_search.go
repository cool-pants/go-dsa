package dsa

import "cmp"


type Comparable interface {
    Equals(x Comparable) bool
    Less(x Comparable) bool
}

func BinarySearch[K cmp.Ordered](a []K, x K) (int, bool) {
    start, end := 0, len(a)
    var mid int
    for start <= end {
        mid = start + (end-start)/2
        if a[mid] == x{
            return mid, true
        } else if a[mid] < x {
            start = mid + 1
        } else {
            end = mid - 1
        }
    }
    return start, false
}

func BinarySearchCustom[K Comparable](a []K, x K) (int, bool) {
    start, end := 0, len(a)
    var mid int
    for start <= end {
        mid = start + (end-start)/2
        if a[mid].Equals(x){
            return mid, true
        } else if a[mid].Less(x) {
            start = mid + 1
        } else {
            end = mid - 1
        }
    }
    return start, false
}
