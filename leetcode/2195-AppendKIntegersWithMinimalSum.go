package main

// 2195. Append K Integers With Minimal Sum
// You are given an integer array nums and an integer k. 
// Append k unique positive integers that do not appear in nums to nums such that the resulting total sum is minimum.

// Return the sum of the k integers appended to nums.

// Example 1:
// Input: nums = [1,4,25,10,25], k = 2
// Output: 5
// Explanation: The two unique positive integers that do not appear in nums which we append are 2 and 3.
// The resulting sum of nums is 1 + 4 + 25 + 10 + 25 + 2 + 3 = 70, which is the minimum.
// The sum of the two integers appended is 2 + 3 = 5, so we return 5.

// Example 2:
// Input: nums = [5,6], k = 6
// Output: 25
// Explanation: The six unique positive integers that do not appear in nums which we append are 1, 2, 3, 4, 7, and 8.
// The resulting sum of nums is 5 + 6 + 1 + 2 + 3 + 4 + 7 + 8 = 36, which is the minimum. 
// The sum of the six integers appended is 1 + 2 + 3 + 4 + 7 + 8 = 25, so we return 25.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     1 <= k <= 10^8

import "fmt"
import "sort"

func minimalKSum(nums []int, k int) int64 {
    sort.Ints(nums)
    res, last := (1 + k) * k / 2, 0
    for _, v := range nums {
        if v <= k {
            if last != v {
                k++
                res += (k - v)
            }
            last = v
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,4,25,10,25], k = 2
    // Output: 5
    // Explanation: The two unique positive integers that do not appear in nums which we append are 2 and 3.
    // The resulting sum of nums is 1 + 4 + 25 + 10 + 25 + 2 + 3 = 70, which is the minimum.
    // The sum of the two integers appended is 2 + 3 = 5, so we return 5.
    fmt.Println(minimalKSum([]int{1,4,25,10,25}, 2)) // 5
    // Example 2:
    // Input: nums = [5,6], k = 6
    // Output: 25
    // Explanation: The six unique positive integers that do not appear in nums which we append are 1, 2, 3, 4, 7, and 8.
    // The resulting sum of nums is 5 + 6 + 1 + 2 + 3 + 4 + 7 + 8 = 36, which is the minimum. 
    // The sum of the six integers appended is 1 + 2 + 3 + 4 + 7 + 8 = 25, so we return 25.
    fmt.Println(minimalKSum([]int{5, 6}, 6)) // 25
}