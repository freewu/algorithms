package main

// 3852. Smallest Pair With Different Frequencies
// You are given an integer array nums.

// Consider all pairs of distinct values x and y from nums such that:
//     1. x < y
//     2. x and y have different frequencies in nums.

// Among all such pairs:
//     1. Choose the pair with the smallest possible value of x.
//     2. If multiple pairs have the same x, choose the one with the smallest possible value of y.
    
// Return an integer array [x, y]. If no valid pair exists, return [-1, -1].

// The frequency of a value x is the number of times it occurs in the array.

// Example 1:
// Input: nums = [1,1,2,2,3,4]
// Output: [1,3]
// Explanation:
// The smallest value is 1 with a frequency of 2, and the smallest value greater than 1 that has a different frequency from 1 is 3 with a frequency of 1. Thus, the answer is [1, 3].

// Example 2:
// Input: nums = [1,5]
// Output: [-1,-1]
// Explanation:
// Both values have the same frequency, so no valid pair exists. Return [-1, -1].

// Example 3:
// Input: nums = [7]
// Output: [-1,-1]
// Explanation:
// There is only one value in the array, so no valid pair exists. Return [-1, -1].

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100

import "fmt"

func minDistinctFreqPair(nums []int) []int {
    const MX = 100
    count := make([]int, MX + 1)
    for _, v := range nums { // 统计出现频率
        count[v]++
    }
    first := -1
    for i := 0; i <= MX; i++ {
        if count[i] > 0 {
            if first == -1 { // 首次
                first = i
            } else if count[i] != count[first] { 
                return []int{ first, i }
            }
        }
    }
    return []int{-1, -1} // 不存在有效的数对 
}

func main() {
    // Example 1:
    // Input: nums = [1,1,2,2,3,4]
    // Output: [1,3]
    // Explanation:
    // The smallest value is 1 with a frequency of 2, and the smallest value greater than 1 that has a different frequency from 1 is 3 with a frequency of 1. Thus, the answer is [1, 3].
    fmt.Println(minDistinctFreqPair([]int{1,1,2,2,3,4})) // [1,3]
    // Example 2:
    // Input: nums = [1,5]
    // Output: [-1,-1]
    // Explanation:
    // Both values have the same frequency, so no valid pair exists. Return [-1, -1].
    fmt.Println(minDistinctFreqPair([]int{1,5})) // [-1,-1]
    // Example 3:
    // Input: nums = [7]
    // Output: [-1,-1]
    // Explanation:
    // There is only one value in the array, so no valid pair exists. Return [-1, -1].
    fmt.Println(minDistinctFreqPair([]int{7})) // [-1,-1]

    fmt.Println(minDistinctFreqPair([]int{1,2,3,4,5,6,7,8,9})) // [-1,-1]
    fmt.Println(minDistinctFreqPair([]int{9,8,7,6,5,4,3,2,1})) // [-1,-1]
}