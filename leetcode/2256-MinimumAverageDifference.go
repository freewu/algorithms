package main

// 2256. Minimum Average Difference
// You are given a 0-indexed integer array nums of length n.

// The average difference of the index i is the absolute difference 
// between the average of the first i + 1 elements of nums and the average of the last n - i - 1 elements. 
// Both averages should be rounded down to the nearest integer.

// Return the index with the minimum average difference. 
// If there are multiple such indices, return the smallest one.

// Note:
//     The absolute difference of two numbers is the absolute value of their difference.
//     The average of n elements is the sum of the n elements divided (integer division) by n.
//     The average of 0 elements is considered to be 0.

// Example 1:
// Input: nums = [2,5,3,9,5,3]
// Output: 3
// Explanation:
// - The average difference of index 0 is: |2 / 1 - (5 + 3 + 9 + 5 + 3) / 5| = |2 / 1 - 25 / 5| = |2 - 5| = 3.
// - The average difference of index 1 is: |(2 + 5) / 2 - (3 + 9 + 5 + 3) / 4| = |7 / 2 - 20 / 4| = |3 - 5| = 2.
// - The average difference of index 2 is: |(2 + 5 + 3) / 3 - (9 + 5 + 3) / 3| = |10 / 3 - 17 / 3| = |3 - 5| = 2.
// - The average difference of index 3 is: |(2 + 5 + 3 + 9) / 4 - (5 + 3) / 2| = |19 / 4 - 8 / 2| = |4 - 4| = 0.
// - The average difference of index 4 is: |(2 + 5 + 3 + 9 + 5) / 5 - 3 / 1| = |24 / 5 - 3 / 1| = |4 - 3| = 1.
// - The average difference of index 5 is: |(2 + 5 + 3 + 9 + 5 + 3) / 6 - 0| = |27 / 6 - 0| = |4 - 0| = 4.
// The average difference of index 3 is the minimum average difference so return 3.

// Example 2:
// Input: nums = [0]
// Output: 0
// Explanation:
// The only index is 0 so return 0.
// The average difference of index 0 is: |0 / 1 - 0| = |0 - 0| = 0.

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^5

import "fmt"

func minimumAverageDifference(nums []int) int {
    res, n := 0, len(nums)
    left, right := make([]int, n), make([]int, n)
    copy(left, nums)
    copy(right, nums)
    for i := 1; i < n; i++ {
        left[i] += left[i-1]
        right[n-1-i] += right[n-i]
    }
    for i := 1; i < n; i++ {
        left[i] /= i + 1
        right[n-1-i] /= i + 1
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    mn := 1 << 31
    for i := 0; i < n - 1; i++ {
        if diff := abs(left[i] - right[i + 1]); diff < mn {
            mn, res = diff, i
        }
    }
    if diff := abs(left[n - 1]); diff < mn {
        res = n - 1
    }
    return res
}

func minimumAverageDifference1(nums []int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    res, s, mn, n := -1, 0, 1 << 31, len(nums)
    for i, v := range nums {
        s += v
        rem, r, ll, rr := sum - s, n - 1 - i, s / (i + 1), 0
        if r > 0 {
            rr = rem / r
        }
        t := abs(ll - rr)
        if t < mn {
            mn, res = t, i
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,5,3,9,5,3]
    // Output: 3
    // Explanation:
    // - The average difference of index 0 is: |2 / 1 - (5 + 3 + 9 + 5 + 3) / 5| = |2 / 1 - 25 / 5| = |2 - 5| = 3.
    // - The average difference of index 1 is: |(2 + 5) / 2 - (3 + 9 + 5 + 3) / 4| = |7 / 2 - 20 / 4| = |3 - 5| = 2.
    // - The average difference of index 2 is: |(2 + 5 + 3) / 3 - (9 + 5 + 3) / 3| = |10 / 3 - 17 / 3| = |3 - 5| = 2.
    // - The average difference of index 3 is: |(2 + 5 + 3 + 9) / 4 - (5 + 3) / 2| = |19 / 4 - 8 / 2| = |4 - 4| = 0.
    // - The average difference of index 4 is: |(2 + 5 + 3 + 9 + 5) / 5 - 3 / 1| = |24 / 5 - 3 / 1| = |4 - 3| = 1.
    // - The average difference of index 5 is: |(2 + 5 + 3 + 9 + 5 + 3) / 6 - 0| = |27 / 6 - 0| = |4 - 0| = 4.
    // The average difference of index 3 is the minimum average difference so return 3.
    fmt.Println(minimumAverageDifference([]int{2,5,3,9,5,3})) // 3
    // Example 2:
    // Input: nums = [0]
    // Output: 0
    // Explanation:
    // The only index is 0 so return 0.
    // The average difference of index 0 is: |0 / 1 - 0| = |0 - 0| = 0.
    fmt.Println(minimumAverageDifference([]int{0})) // 0

    fmt.Println(minimumAverageDifference1([]int{2,5,3,9,5,3})) // 3
    fmt.Println(minimumAverageDifference1([]int{0})) // 0
}