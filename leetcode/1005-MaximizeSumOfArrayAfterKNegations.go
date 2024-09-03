package main

// 1005. Maximize Sum Of Array After K Negations
// Given an integer array nums and an integer k, modify the array in the following way:
//     choose an index i and replace nums[i] with -nums[i].

// You should apply this process exactly k times. 
// You may choose the same index i multiple times.

// Return the largest possible sum of the array after modifying it in this way.

// Example 1:
// Input: nums = [4,2,3], k = 1
// Output: 5
// Explanation: Choose index 1 and nums becomes [4,-2,3].

// Example 2:
// Input: nums = [3,-1,0,2], k = 3
// Output: 6
// Explanation: Choose indices (1, 2, 2) and nums becomes [3,1,0,2].

// Example 3:
// Input: nums = [2,-3,-1,5,-4], k = 2
// Output: 13
// Explanation: Choose indices (1, 4) and nums becomes [2,3,-1,5,4].

// Constraints:
//     1 <= nums.length <= 10^4
//     -100 <= nums[i] <= 100
//     1 <= k <= 10^4

import "fmt"
import "sort"

func largestSumAfterKNegations(nums []int, k int) int {
    sort.Ints(nums)
    sum, pos := 0, 0
    for k > 0 {
        if nums[pos] == 0 { break }
        if pos < len(nums) - 1 {
            if (nums[pos] < 0 && nums[pos + 1] < 0) || nums[pos]*-1 > nums[pos+1] {
                nums[pos] *= -1
                pos++
            } else {
                nums[pos] *= -1
            }
        } else {
            nums[pos] *= -1
        }
        k--
    }
    for _, v := range nums {
        sum += v
    }
    return sum
}

func main() {
    // Example 1:
    // Input: nums = [4,2,3], k = 1
    // Output: 5
    // Explanation: Choose index 1 and nums becomes [4,-2,3].
    fmt.Println(largestSumAfterKNegations([]int{4,2,3}, 1)) // 5
    // Example 2:
    // Input: nums = [3,-1,0,2], k = 3
    // Output: 6
    // Explanation: Choose indices (1, 2, 2) and nums becomes [3,1,0,2].
    fmt.Println(largestSumAfterKNegations([]int{3,-1,0,2}, 3)) // 6
    // Example 3:
    // Input: nums = [2,-3,-1,5,-4], k = 2
    // Output: 13
    // Explanation: Choose indices (1, 4) and nums becomes [2,3,-1,5,4].
    fmt.Println(largestSumAfterKNegations([]int{2,-3,-1,5,-4}, 2)) // 13

    fmt.Println(largestSumAfterKNegations([]int{4,2,3}, 1)) // 5
}