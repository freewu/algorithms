package main

// 3375. Minimum Operations to Make Array Values Equal to K
// You are given an integer array nums and an integer k.

// An integer h is called valid if all values in the array that are strictly greater than h are identical.

// For example, if nums = [10, 8, 10, 8], a valid integer is h = 9 because all nums[i] > 9 are equal to 10, but 5 is not a valid integer.

// You are allowed to perform the following operation on nums:
//     1. Select an integer h that is valid for the current values in nums.
//     2. For each index i where nums[i] > h, set nums[i] to h.

// Return the minimum number of operations required to make every element in nums equal to k. 
// If it is impossible to make all elements equal to k, return -1.

// Example 1:
// Input: nums = [5,2,5,4,5], k = 2
// Output: 2
// Explanation:
// The operations can be performed in order using valid integers 4 and then 2.

// Example 2:
// Input: nums = [2,1,2], k = 2
// Output: -1
// Explanation:
// It is impossible to make all the values equal to 2.

// Example 3:
// Input: nums = [9,7,5,3], k = 1
// Output: 4
// Explanation:
// The operations can be performed using valid integers in the order 7, 5, 3, and 1.

// Constraints:
//     1 <= nums.length <= 100 
//     1 <= nums[i] <= 100
//     1 <= k <= 100

import "fmt"
import "sort"
import "slices"

func minOperations(nums []int, k int) int {
    sort.Ints(nums)
    if nums[0] < k { return -1 }
    res := 0
    for i := len(nums) - 1; i >= 1; i-- {
        if nums[i] != nums[i-1] {
           res++
        }
    }
    if nums[0] > k {
        res++
    }
    return res
}

func minOperations1(nums []int, k int) int {
    mn := slices.Min(nums)
    if k > mn { return -1 }
    set := make(map[int]struct{})
    for _, v := range nums {
        set[v] = struct{}{}
    }
    if k == mn {
        return len(set) - 1
    }
    return len(set)
}

func main() {
    // Example 1:
    // Input: nums = [5,2,5,4,5], k = 2
    // Output: 2
    // Explanation:
    // The operations can be performed in order using valid integers 4 and then 2.
    fmt.Println(minOperations([]int{5,2,5,4,5}, 2)) // 2
    // Example 2:
    // Input: nums = [2,1,2], k = 2
    // Output: -1
    // Explanation:
    // It is impossible to make all the values equal to 2.
    fmt.Println(minOperations([]int{2,1,2}, 2)) // -1
    // Example 3:
    // Input: nums = [9,7,5,3], k = 1
    // Output: 4
    // Explanation:
    // The operations can be performed using valid integers in the order 7, 5, 3, and 1.
    fmt.Println(minOperations([]int{9,7,5,3}, 1)) // 4

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9}, 1)) // 8
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1}, 1)) // 8

    fmt.Println(minOperations1([]int{5,2,5,4,5}, 2)) // 2
    fmt.Println(minOperations1([]int{2,1,2}, 2)) // -1
    fmt.Println(minOperations1([]int{9,7,5,3}, 1)) // 4
    fmt.Println(minOperations1([]int{1,2,3,4,5,6,7,8,9}, 1)) // 8
    fmt.Println(minOperations1([]int{9,8,7,6,5,4,3,2,1}, 1)) // 8
}