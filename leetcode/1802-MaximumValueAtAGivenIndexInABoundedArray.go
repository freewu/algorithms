package main

// 1802. Maximum Value at a Given Index in a Bounded Array
// You are given three positive integers: n, index, and maxSum. 
// You want to construct an array nums (0-indexed) that satisfies the following conditions:
//     nums.length == n
//     nums[i] is a positive integer where 0 <= i < n.
//     abs(nums[i] - nums[i+1]) <= 1 where 0 <= i < n-1.
//     The sum of all the elements of nums does not exceed maxSum.
//     nums[index] is maximized.

// Return nums[index] of the constructed array.
// Note that abs(x) equals x if x >= 0, and -x otherwise.

// Example 1:
// Input: n = 4, index = 2,  maxSum = 6
// Output: 2
// Explanation: nums = [1,2,2,1] is one array that satisfies all the conditions.
// There are no arrays that satisfy all the conditions and have nums[2] == 3, so 2 is the maximum nums[2].

// Example 2:
// Input: n = 6, index = 1,  maxSum = 10
// Output: 3

// Constraints:
//     1 <= n <= maxSum <= 10^9
//     0 <= index < n

import "fmt"

func maxValue(n int, index int, maxSum int) int {
    if n == maxSum {
        return 1
    }
    res := (maxSum + (index*(index+1)+(n-1-index)*(n-index))/2) / n
    if res > index && res > n-1 - index { // maxSum is large enough
        return res
    }
    total, left, right := n, 0, 0
    res = 1
    for total < maxSum {
        res++
        if left < index {
            left++
        }
        if right < n - 1 - index {
            right++
        }
        total += left + right + 1
    }
    return res
}

func maxValue1(n, index, maxSum int) int {
    left, right := 1, maxSum
    cal := func(max, length int) int {
        if length == 0 {
            return 0
        }
        if length <= max {
            return (2 * max + 1 - length) * length / 2
        }
        return (max + 1)* max / 2 + length - max
    }
    for left < right {
        mid := left + (right-left) >> 1
        if mid + cal(mid, index) + cal(mid, n-index-1) < maxSum {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}

func main() {
    // Input: n = 4, index = 2,  maxSum = 6
// Output: 2
    // Explanation: nums = [1,2,2,1] is one array that satisfies all the conditions.
    // There are no arrays that satisfy all the conditions and have nums[2] == 3, so 2 is the maximum nums[2].
    fmt.Println(maxValue(4, 2, 6)) // 2
    fmt.Println(maxValue(6, 1, 10)) // 3

    fmt.Println(maxValue1(4, 2, 6)) // 2
    fmt.Println(maxValue1(6, 1, 10)) // 3
}