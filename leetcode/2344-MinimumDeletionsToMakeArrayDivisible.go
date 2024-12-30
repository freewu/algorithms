package main

// 2344. Minimum Deletions to Make Array Divisible
// You are given two positive integer arrays nums and numsDivide. 
// You can delete any number of elements from nums.

// Return the minimum number of deletions such that the smallest element in nums divides all the elements of numsDivide. 
// If this is not possible, return -1.

// Note that an integer x divides y if y % x == 0.

// Example 1:
// Input: nums = [2,3,2,4,3], numsDivide = [9,6,9,3,15]
// Output: 2
// Explanation: 
// The smallest element in [2,3,2,4,3] is 2, which does not divide all the elements of numsDivide.
// We use 2 deletions to delete the elements in nums that are equal to 2 which makes nums = [3,4,3].
// The smallest element in [3,4,3] is 3, which divides all the elements of numsDivide.
// It can be shown that 2 is the minimum number of deletions needed.

// Example 2:
// Input: nums = [4,3,6], numsDivide = [8,2,6,10]
// Output: -1
// Explanation: 
// We want the smallest element in nums to divide all the elements of numsDivide.
// There is no way to delete elements from nums to allow this.

// Constraints:
//     1 <= nums.length, numsDivide.length <= 10^5
//     1 <= nums[i], numsDivide[i] <= 10^9

import "fmt"
import "sort"

func minOperations(nums []int, numsDivide []int) int {
    divide := func(n int) bool {
        for _, v := range numsDivide {
            if v % n != 0 { return false }
        }
        return true
    }
    sort.Ints(nums)
    prev := -1
    for i := 0; i < len(nums); i++ {
        if prev == nums[i] { continue }
        prev = nums[i]
        if divide(nums[i]) { return i }
    }
    return -1
}

func minOperations1(nums []int, numsDivide []int) int {
    d := numsDivide[0]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for _, b := range numsDivide {
        d, b = max(d, b), min(d, b)
        d = gcd(d, b)
    }
    res, mn := 0, d + 1
    for _, v := range nums {
        if d % v == 0 {
            mn = min(mn, v)
        }
    }
    if mn == d + 1 {  return -1 }
    for _, v := range nums {
        if v < mn {
            res++ 
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,3,2,4,3], numsDivide = [9,6,9,3,15]
    // Output: 2
    // Explanation: 
    // The smallest element in [2,3,2,4,3] is 2, which does not divide all the elements of numsDivide.
    // We use 2 deletions to delete the elements in nums that are equal to 2 which makes nums = [3,4,3].
    // The smallest element in [3,4,3] is 3, which divides all the elements of numsDivide.
    // It can be shown that 2 is the minimum number of deletions needed.
    fmt.Println(minOperations([]int{2,3,2,4,3}, []int{9,6,9,3,15})) // 2
    // Example 2:
    // Input: nums = [4,3,6], numsDivide = [8,2,6,10]
    // Output: -1
    // Explanation: 
    // We want the smallest element in nums to divide all the elements of numsDivide.
    // There is no way to delete elements from nums to allow this.
    fmt.Println(minOperations([]int{4,3,6}, []int{8,2,6,10})) // -1

    fmt.Println(minOperations1([]int{2,3,2,4,3}, []int{9,6,9,3,15})) // 2
    fmt.Println(minOperations1([]int{4,3,6}, []int{8,2,6,10})) // -1
}