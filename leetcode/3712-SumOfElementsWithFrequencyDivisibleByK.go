package main

// 3712. Sum of Elements With Frequency Divisible by K
// You are given an integer array nums and an integer k.

// Return an integer denoting the sum of all elements in nums whose frequency is divisible by k, or 0 if there are no such elements.

// Note: An element is included in the sum exactly as many times as it appears in the array if its total frequency is divisible by k.

// Example 1:
// Input: nums = [1,2,2,3,3,3,3,4], k = 2
// Output: 16
// Explanation:
// The number 1 appears once (odd frequency).
// The number 2 appears twice (even frequency).
// The number 3 appears four times (even frequency).
// The number 4 appears once (odd frequency).
// So, the total sum is 2 + 2 + 3 + 3 + 3 + 3 = 16.

// Example 2:
// Input: nums = [1,2,3,4,5], k = 2
// Output: 0
// Explanation:
// There are no elements that appear an even number of times, so the total sum is 0.

// Example 3:
// Input: nums = [4,4,4,1,2,3], k = 3
// Output: 12
// Explanation:
// The number 1 appears once.
// The number 2 appears once.
// The number 3 appears once.
// The number 4 appears three times.
// So, the total sum is 4 + 4 + 4 = 12.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100
//     1 <= k <= 100

import "fmt"

func sumDivisibleByK(nums []int, k int) int {
    res, count := 0, make([]int,101)
    for _, v := range nums {
        count[v]++
    }    
    for i, c := range count {   
        if c % k == 0 {
            res += i * c
        }
    } 
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,2,3,3,3,3,4], k = 2
    // Output: 16
    // Explanation:
    // The number 1 appears once (odd frequency).
    // The number 2 appears twice (even frequency).
    // The number 3 appears four times (even frequency).
    // The number 4 appears once (odd frequency).
    // So, the total sum is 2 + 2 + 3 + 3 + 3 + 3 = 16.
    fmt.Println(sumDivisibleByK([]int{1,2,2,3,3,3,3,4}, 2)) // 16
    // Example 2:
    // Input: nums = [1,2,3,4,5], k = 2
    // Output: 0
    // Explanation:
    // There are no elements that appear an even number of times, so the total sum is 0.
    fmt.Println(sumDivisibleByK([]int{1,2,3,4,5}, 2)) // 0
    // Example 3:
    // Input: nums = [4,4,4,1,2,3], k = 3
    // Output: 12
    // Explanation:
    // The number 1 appears once.
    // The number 2 appears once.
    // The number 3 appears once.
    // The number 4 appears three times.
    // So, the total sum is 4 + 4 + 4 = 12.
    fmt.Println(sumDivisibleByK([]int{4,4,4,1,2,3}, 3)) // 12   

    fmt.Println(sumDivisibleByK([]int{1,2,3,4,5,6,7,8,9}, 2)) // 0
    fmt.Println(sumDivisibleByK([]int{9,8,7,6,5,4,3,2,1}, 2)) // 0
}