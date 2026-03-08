package main

// 3862. Find the Smallest Balanced Index
// You are given an integer array nums.

// An index i is balanced if the sum of elements strictly to the left of i equals the product of elements strictly to the right of i.

// If there are no elements to the left, the sum is considered as 0. 
// Similarly, if there are no elements to the right, the product is considered as 1.

// Return an integer denoting the smallest balanced index. 
// If no balanced index exists, return -1.

// Example 1:
// Input: nums = [2,1,2]
// Output: 1
// Explanation:
// For index i = 1:
// Left sum = nums[0] = 2
// Right product = nums[2] = 2
// Since the left sum equals the right product, index 1 is balanced.
// No smaller index satisfies the condition, so the answer is 1.

// Example 2:
// Input: nums = [2,8,2,2,5]
// Output: 2
// Explanation:
// For index i = 2:
// Left sum = 2 + 8 = 10
// Right product = 2 * 5 = 10
// Since the left sum equals the right product, index 2 is balanced.
// No smaller index satisfies the condition, so the answer is 2.

// Example 3:
// Input: nums = [1]
// Output: -1
// For index i = 0:
// The left side is empty, so the left sum is 0.
// The right side is empty, so the right product is 1.
// Since the left sum does not equal the right product, index 0 is not balanced.
// Therefore, no balanced index exists and the answer is -1.
 
// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"

// time: O(n), space: O(1)
func smallestBalancedIndex(nums []int) int {
    sum := int64(0)
    for _, v := range nums {
        sum += int64(v)
    }
    prod := int64(1)
    for i := len(nums)-1; i >= 0 && prod < sum; i-- {
        sum -= int64(nums[i])
        if sum == prod {
            return i
        }
        prod *= int64(nums[i])
    }
    return -1
}

func main() {
    // Example 1:
    // Input: nums = [2,1,2]
    // Output: 1
    // Explanation:
    // For index i = 1:
    // Left sum = nums[0] = 2
    // Right product = nums[2] = 2
    // Since the left sum equals the right product, index 1 is balanced.
    // No smaller index satisfies the condition, so the answer is 1.
    fmt.Println(smallestBalancedIndex([]int{2,1,2})) // 1
    // Example 2:
    // Input: nums = [2,8,2,2,5]
    // Output: 2
    // Explanation:
    // For index i = 2:
    // Left sum = 2 + 8 = 10
    // Right product = 2 * 5 = 10
    // Since the left sum equals the right product, index 2 is balanced.
    // No smaller index satisfies the condition, so the answer is 2.
    fmt.Println(smallestBalancedIndex([]int{2,8,2,2,5})) // 2
    // Example 3:
    // Input: nums = [1]
    // Output: -1
    // For index i = 0:
    // The left side is empty, so the left sum is 0.
    // The right side is empty, so the right product is 1.
    // Since the left sum does not equal the right product, index 0 is not balanced.
    // Therefore, no balanced index exists and the answer is -1.
    fmt.Println(smallestBalancedIndex([]int{1})) // -1

    fmt.Println(smallestBalancedIndex([]int{999,818,984,995,841,822,984,978,960,997,896,926,759,961,1000,562,1,1,1,87,4,1,40})) // 15
    fmt.Println(smallestBalancedIndex([]int{1,2,3,4,5,6,7,8,9})) // -1
    fmt.Println(smallestBalancedIndex([]int{9,8,7,6,5,4,3,2,1})) // -1
}