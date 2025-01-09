package main

// 2475. Number of Unequal Triplets in Array
// You are given a 0-indexed array of positive integers nums. 
// Find the number of triplets (i, j, k) that meet the following conditions:
//     1. 0 <= i < j < k < nums.length
//     2. nums[i], nums[j], and nums[k] are pairwise distinct.
//         In other words, nums[i] != nums[j], nums[i] != nums[k], and nums[j] != nums[k].

// Return the number of triplets that meet the conditions.

// Example 1:
// Input: nums = [4,4,2,4,3]
// Output: 3
// Explanation: The following triplets meet the conditions:
// - (0, 2, 4) because 4 != 2 != 3
// - (1, 2, 4) because 4 != 2 != 3
// - (2, 3, 4) because 2 != 4 != 3
// Since there are 3 triplets, we return 3.
// Note that (2, 0, 4) is not a valid triplet because 2 > 0.

// Example 2:
// Input: nums = [1,1,1,1,1]
// Output: 0
// Explanation: No triplets meet the conditions so we return 0.

// Constraints:
//     3 <= nums.length <= 100
//     1 <= nums[i] <= 1000

import "fmt"
import "sort"

func unequalTriplets1(nums []int) int {
    sort.Ints(nums)
    res, start, n := 0, 0, len(nums)
    for i, x := range nums[:n-1] {
        if x != nums[i+1] {
            res += start * (i - start + 1) * (n - 1 - i)
            start = i + 1
        }
    }
    return res
}

func unequalTriplets(nums []int) int {
    res, n := 0, len(nums)
    for i := 0; i < n -2; i++ {
        for j := i + 1; j < n - 1; j++ {
            for k := j + 1; k < n; k++ {
                if nums[i] != nums[j] && nums[i] != nums[k] && nums[j] != nums[k] {
                    res++
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [4,4,2,4,3]
    // Output: 3
    // Explanation: The following triplets meet the conditions:
    // - (0, 2, 4) because 4 != 2 != 3
    // - (1, 2, 4) because 4 != 2 != 3
    // - (2, 3, 4) because 2 != 4 != 3
    // Since there are 3 triplets, we return 3.
    // Note that (2, 0, 4) is not a valid triplet because 2 > 0.
    fmt.Println(unequalTriplets([]int{4,4,2,4,3})) // 3
    // Example 2:
    // Input: nums = [1,1,1,1,1]
    // Output: 0
    // Explanation: No triplets meet the conditions so we return 0.
    fmt.Println(unequalTriplets([]int{1,1,1,1,1})) // 0

    fmt.Println(unequalTriplets([]int{1,2,3,4,5,6,7,8,9})) // 84
    fmt.Println(unequalTriplets([]int{9,8,7,6,5,4,3,2,1})) // 84

    fmt.Println(unequalTriplets1([]int{4,4,2,4,3})) // 3
    fmt.Println(unequalTriplets1([]int{1,1,1,1,1})) // 0
    fmt.Println(unequalTriplets1([]int{1,2,3,4,5,6,7,8,9})) // 84
    fmt.Println(unequalTriplets1([]int{9,8,7,6,5,4,3,2,1})) // 84
}