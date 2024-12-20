package main

// 3231. Minimum Number of Increasing Subsequence to Be Removed
// Given an array of integers nums, you are allowed to perform the following operation any number of times:
//     Remove a strictly increasing subsequence from the array.

// Your task is to find the minimum number of operations required to make the array empty.

// Example 1:
// Input: nums = [5,3,1,4,2]
// Output: 3
// Explanation:
// We remove subsequences [1, 2], [3, 4], [5].

// Example 2:
// Input: nums = [1,2,3,4,5]
// Output: 1

// Example 3:
// Input: nums = [5,4,3,2,1]
// Output: 5

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

func minOperations(nums []int) int {
    group := []int{}
    for _, v := range nums {
        l, r := 0, len(group)
        for l < r {
            mid := (l + r) >> 1
            if group[mid] < v {
                r = mid
            } else {
                l = mid + 1
            }
        }
        if l == len(group) {
            group = append(group, v)
        } else {
            group[l] = v
        }
    }
    return len(group)
}

func main() {
    // Example 1:
    // Input: nums = [5,3,1,4,2]
    // Output: 3
    // Explanation:
    // We remove subsequences [1, 2], [3, 4], [5].
    fmt.Println(minOperations([]int{5,3,1,4,2})) // 3
    // Example 2:
    // Input: nums = [1,2,3,4,5]
    // Output: 1
    fmt.Println(minOperations([]int{1,2,3,4,5})) // 1
    // Example 3:
    // Input: nums = [5,4,3,2,1]
    // Output: 5
    fmt.Println(minOperations([]int{5,4,3,2,1})) // 5
}