package leetcode

import (
    "testing"
    "github.com/stretchr/testify/assert"
)


func mySqrt(x int) int {
    if x == 1 {
        return x
    }
    var start, end int = 0, x/2
    var mid, ans int

    for start <= end {
        mid = start + (end - start)/2
        if mid * mid == x {
            return mid
        }
        if mid*mid < x {
            ans = max(ans, mid)
            start = mid + 1
        } else {
            end = mid - 1
        }
    }
    return ans
}

func TestSqrtFunc(t *testing.T) {
    assert.Equal(t, 2, mySqrt(4))
    assert.Equal(t, 2, mySqrt(8))
    assert.Equal(t, 3, mySqrt(9))
    assert.Equal(t, 1, mySqrt(1))
    assert.Equal(t, 0, mySqrt(0))
    assert.Equal(t, 3, mySqrt(11))
}
