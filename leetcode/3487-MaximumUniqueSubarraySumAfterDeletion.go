package main

// 3487. Maximum Unique Subarray Sum After Deletion
// You are given an integer array nums.

// You are allowed to delete any number of elements from nums without making it empty. 
// After performing the deletions, select a subarray of nums such that:
//     1. All elements in the subarray are unique.
//     2. The sum of the elements in the subarray is maximized.

// Return the maximum sum of such a subarray.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [1,2,3,4,5]
// Output: 15
// Explanation:
// Select the entire array without deleting any element to obtain the maximum sum.

// Example 2:
// Input: nums = [1,1,0,1,1]
// Output: 1
// Explanation:
// Delete the element nums[0] == 1, nums[1] == 1, nums[2] == 0, and nums[3] == 1. Select the entire array [1] to obtain the maximum sum.

// Example 3:
// Input: nums = [1,2,-1,-2,1,0,-1]
// Output: 3
// Explanation:
// Delete the elements nums[2] == -1 and nums[3] == -2, and select the subarray [2, 1] from [1, 2, 1, 0, -1] to obtain the maximum sum.

// Constraints:
//     1 <= nums.length <= 100
//     -100 <= nums[i] <= 100

import "fmt"

func maxSum(nums []int) int {
    set := make(map[int]bool)
    res, mx, found := 0, -1 << 31, false
    for _, v := range nums {
        if !set[v] {
            if v >= 0 {
                set[v] = true
                res += v
                found = true
            } else {
                mx = max(mx, v)
            }
        }
    }
    if !found {
        return mx
    }
    return res
}

func maxSum1(nums []int) int {
    sum, mx := 0, -10000
    set := make(map[int]bool)
    for _, v := range nums {
        set[v] = true
        if v > mx {
            mx = v
        }
    }
    if mx <= 0 { return mx }
    for k := range set {
        if k > 0 {
            sum += k
        }
    }
    return sum
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5]
    // Output: 15
    // Explanation:
    // Select the entire array without deleting any element to obtain the maximum sum.
    fmt.Println(maxSum([]int{1,2,3,4,5})) // 15
    // Example 2:
    // Input: nums = [1,1,0,1,1]
    // Output: 1
    // Explanation:
    // Delete the element nums[0] == 1, nums[1] == 1, nums[2] == 0, and nums[3] == 1. Select the entire array [1] to obtain the maximum sum.
    fmt.Println(maxSum([]int{1,1,0,1,1})) // 1
    // Example 3:
    // Input: nums = [1,2,-1,-2,1,0,-1]
    // Output: 3
    // Explanation:
    // Delete the elements nums[2] == -1 and nums[3] == -2, and select the subarray [2, 1] from [1, 2, 1, 0, -1] to obtain the maximum sum.
    fmt.Println(maxSum([]int{1,2,-1,-2,1,0,-1})) // 3

    fmt.Println(maxSum([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxSum([]int{9,8,7,6,5,4,3,2,1})) // 45

    fmt.Println(maxSum1([]int{1,2,3,4,5})) // 15
    fmt.Println(maxSum1([]int{1,1,0,1,1})) // 1
    fmt.Println(maxSum1([]int{1,2,-1,-2,1,0,-1})) // 3
    fmt.Println(maxSum1([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxSum1([]int{9,8,7,6,5,4,3,2,1})) // 45
}