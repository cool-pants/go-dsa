package leetcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func missingNumber(nums []int) int {
	sort.Ints(nums)
	start, end, mid := 0, len(nums)-1, 0
	// 0,1,3
	// 0,1,2
	// len = 9
	if nums[end] != end+1 {
		return end + 1
	}
	for start < end {
		mid = start + (end-start)/2
		if nums[mid] == mid {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

func TestMissingNumber(t *testing.T) {
	assert.Equal(t, 2, missingNumber([]int{3, 0, 1}))
	assert.Equal(t, 2, missingNumber([]int{0, 1}))
	assert.Equal(t, 8, missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
