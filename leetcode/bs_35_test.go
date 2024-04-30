package leetcode

/*
tags: Leetcode, Binary Search, Easy
name: 35 - Search Insert Position
*/

import (
	"testing"

	"github.com/cool-pants/go-dsa/dsa"
	"github.com/stretchr/testify/assert"
)

func searchInsert(nums []int, target int) int {
    pos, _ := dsa.BinarySearch(nums, target)
    return pos
}

func TestSearchInsert(t *testing.T) {
    var (
        nums []int
        target int
        ans int
    )
    nums, target = []int{1,3,5,6}, 5
    ans = searchInsert(nums, target)
    assert.Equal(t, 2, ans)

    nums, target = []int{1,3,5,6}, 2
    ans = searchInsert(nums, target)
    assert.Equal(t, 1, ans)

    nums, target = []int{1,3,5,6}, 7
    ans = searchInsert(nums, target)
    assert.Equal(t, 4, ans)
}
