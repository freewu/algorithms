package main

// 3550. Smallest Index With Digit Sum Equal to Index
// You are given an integer array nums.

// Return the smallest index i such that the sum of the digits of nums[i] is equal to i.

// If no such index exists, return -1.

// Example 1:
// Input: nums = [1,3,2]
// Output: 2
// Explanation:
// For nums[2] = 2, the sum of digits is 2, which is equal to index i = 2. Thus, the output is 2.

// Example 2:
// Input: nums = [1,10,11]
// Output: 1
// Explanation:
// For nums[1] = 10, the sum of digits is 1 + 0 = 1, which is equal to index i = 1.
// For nums[2] = 11, the sum of digits is 1 + 1 = 2, which is equal to index i = 2.
// Since index 1 is the smallest, the output is 1.

// Example 3:
// Input: nums = [1,2,3]
// Output: -1
// Explanatio
// Since no index satisfies the condition, the output is -1.

// Constraints:
//     1 <= nums.length <= 100
//     0 <= nums[i] <= 1000

import "fmt"

func smallestIndex(nums []int) int {
    calc := func(num int) int {
        sum := 0
        for num > 0 {
            sum += num % 10
            num /= 10
        }
        return sum
    }
    for i, v := range nums {
        if calc(v) == i {
            return i
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: nums = [1,3,2]
    // Output: 2
    // Explanation:
    // For nums[2] = 2, the sum of digits is 2, which is equal to index i = 2. Thus, the output is 2.
    fmt.Println(smallestIndex([]int{1,3,2})) // 2
    // Example 2:
    // Input: nums = [1,10,11]
    // Output: 1
    // Explanation:
    // For nums[1] = 10, the sum of digits is 1 + 0 = 1, which is equal to index i = 1.
    // For nums[2] = 11, the sum of digits is 1 + 1 = 2, which is equal to index i = 2.
    // Since index 1 is the smallest, the output is 1.
    fmt.Println(smallestIndex([]int{1,10,11})) // 1
    // Example 3:
    // Input: nums = [1,2,3]
    // Output: -1
    // Explanatio
    // Since no index satisfies the condition, the output is -1.
    fmt.Println(smallestIndex([]int{1,2,3})) // -1

    fmt.Println(smallestIndex([]int{1,2,3,4,5,6,7,8,9})) // -1
    fmt.Println(smallestIndex([]int{9,8,7,6,5,4,3,2,1})) // -1
}