package main

// 3917. Count Indices With Opposite Parity
// You are given an integer array nums of length n.

// The score of an index i is defined as the number of indices j such that:
//     1. i < j < n, and
//     2. nums[i] and nums[j] have different parity (one is even and the other is odd).

// Return an integer array answer of length n, where answer[i] is the score of index i.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: [2,1,1,0]
// Explanation:
// nums[0] = 1, which is odd. Thus, the indices j = 1 and j = 3 satisfy the conditions, so the score of index 0 is 2.
// nums[1] = 2, which is even. Thus, the index j = 2 satisfies the conditions, so the score of index 1 is 1.
// nums[2] = 3, which is odd. Thus, the index j = 3 satisfies the conditions, so the score of index 2 is 1.
// nums[3] = 4, which is even. Thus, no index satisfies the conditions, so the score of index 3 is 0.
// Thus, the answer = [2, 1, 1, 0].

// Example 2:
// Input: nums = [1]
// Output: [0]
// Explanation:
// There is only one element in nums. Thus, the score of index 0 is 0.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100

import "fmt"

func countOppositeParity(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    for i := 0; i<n; i++ {
        for j := i+1; j<n; j++ {
            if nums[i]%2 != nums[j] % 2 {
                res[i] += 1
            }
        }
    }
    return res
}

func countOppositeParity1(nums []int) []int {
    n, evens, odds := len(nums), 0, 0
    res := make([]int, n)
    for i := n - 1; i >= 0; i-- {
        if nums[i] % 2 == 0 {
            res[i] = odds
            evens++
        } else {
            res[i] = evens
            odds++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4]
    // Output: [2,1,1,0]
    // Explanation:
    // nums[0] = 1, which is odd. Thus, the indices j = 1 and j = 3 satisfy the conditions, so the score of index 0 is 2.
    // nums[1] = 2, which is even. Thus, the index j = 2 satisfies the conditions, so the score of index 1 is 1.
    // nums[2] = 3, which is odd. Thus, the index j = 3 satisfies the conditions, so the score of index 2 is 1.
    // nums[3] = 4, which is even. Thus, no index satisfies the conditions, so the score of index 3 is 0.
    // Thus, the answer = [2, 1, 1, 0].
    fmt.Println(countOppositeParity([]int{1,2,3,4})) // [2,1,1,0]
    // Example 2:
    // Input: nums = [1]
    // Output: [0]
    // Explanation:
    // There is only one element in nums. Thus, the score of index 0 is 0.
    fmt.Println(countOppositeParity([]int{1})) // [0]

    fmt.Println(countOppositeParity([]int{1,2,3,4,5,6,7,8,9})) // [4 4 3 3 2 2 1 1 0]
    fmt.Println(countOppositeParity([]int{9,8,7,6,5,4,3,2,1})) // [4 4 3 3 2 2 1 1 0]

    fmt.Println(countOppositeParity1([]int{1,2,3,4})) // [2,1,1,0]
    fmt.Println(countOppositeParity1([]int{1})) // [0]
    fmt.Println(countOppositeParity1([]int{1,2,3,4,5,6,7,8,9})) // [4 4 3 3 2 2 1 1 0]
    fmt.Println(countOppositeParity1([]int{9,8,7,6,5,4,3,2,1})) // [4 4 3 3 2 2 1 1 0]
}