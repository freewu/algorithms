package main

// 674. Longest Continuous Increasing Subsequence
// Given an unsorted array of integers nums, return the length of the longest continuous increasing subsequence (i.e. subarray). 
// The subsequence must be strictly increasing.

// A continuous increasing subsequence is defined by two indices l and r (l < r) such 
// that it is [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] and for each l <= i < r, nums[i] < nums[i + 1].

// Example 1:
// Input: nums = [1,3,5,4,7]
// Output: 3
// Explanation: The longest continuous increasing subsequence is [1,3,5] with length 3.
// Even though [1,3,5,7] is an increasing subsequence, it is not continuous as elements 5 and 7 are separated by element 4.

// Example 2:
// Input: nums = [2,2,2,2,2]
// Output: 1
// Explanation: The longest continuous increasing subsequence is [2] with length 1. Note that it must be strictly
// increasing.
 
// Constraints:
//     1 <= nums.length <= 10^4
//     -10^9 <= nums[i] <= 10^9

import "fmt"

func findLengthOfLCIS(nums []int) int {
    if len(nums) == 0 { 
        return 0 
    }
    res, tmp := 1, 1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i - 1] {
            tmp++
        } else {
            res = max(res, tmp)
            tmp = 1
        }
    }
    return max(res, tmp)
}

func findLengthOfLCIS1(nums []int) int {
    res, count := 1, 1
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i - 1] {
            count++
            continue
        }
        if count > res {
            res = count
        }
        count = 1
    }
    if count > res {
        return count
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,3,5,4,7]
    // Output: 3
    // Explanation: The longest continuous increasing subsequence is [1,3,5] with length 3.
    // Even though [1,3,5,7] is an increasing subsequence, it is not continuous as elements 5 and 7 are separated by element 4.
    fmt.Println(findLengthOfLCIS([]int{1,3,5,4,7})) // 3
    // Example 2:
    // Input: nums = [2,2,2,2,2]
    // Output: 1
    // Explanation: The longest continuous increasing subsequence is [2] with length 1. Note that it must be strictly
    // increasing.
    fmt.Println(findLengthOfLCIS([]int{2,2,2,2,2})) // 1

    fmt.Println(findLengthOfLCIS1([]int{1,3,5,4,7})) // 3
    fmt.Println(findLengthOfLCIS1([]int{2,2,2,2,2})) // 1
}