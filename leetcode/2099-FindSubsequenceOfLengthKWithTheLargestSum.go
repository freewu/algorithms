package main

// 2099. Find Subsequence of Length K With the Largest Sum
// You are given an integer array nums and an integer k. 
// You want to find a subsequence of nums of length k that has the largest sum.

// Return any such subsequence as an integer array of length k.

// A subsequence is an array that can be derived from another array by deleting some 
// or no elements without changing the order of the remaining elements.

// Example 1:
// Input: nums = [2,1,3,3], k = 2
// Output: [3,3]
// Explanation:
// The subsequence has the largest sum of 3 + 3 = 6.

// Example 2:
// Input: nums = [-1,-2,3,4], k = 3
// Output: [-1,3,4]
// Explanation: 
// The subsequence has the largest sum of -1 + 3 + 4 = 6.

// Example 3:
// Input: nums = [3,4,3,3], k = 2
// Output: [3,4]
// Explanation:
// The subsequence has the largest sum of 3 + 4 = 7. 
// Another possible subsequence is [4, 3].

// Constraints:
//     1 <= nums.length <= 1000
//     -10^5 <= nums[i] <= 10^5
//     1 <= k <= nums.length

import "fmt"
import "sort"

func maxSubsequence(nums []int, k int) []int {
    minIndex := func (nums ...int) int { // 找到最小的 index
        mn, index := nums[0], 0
        for i, v := range nums {
            if v < mn {
                mn, index = v, i
            }
        }
        return index
    }
    for k < len(nums) {
        index := minIndex(nums...)
        // remove minIndex num 
        nums = append(nums[:index], nums[index + 1:]...)
    }
    return nums
}

func maxSubsequence1(nums []int, k int) []int {
    if k >= len(nums) { return nums }
    type Pair struct { val, index int }
    pairs := make([]Pair, len(nums))
    for i, v := range nums {
        pairs[i] = Pair{ val: v, index: i }
    }
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i].val > pairs[j].val || (pairs[i].val == pairs[j].val && pairs[i].index < pairs[j].index)
    })
    selected := pairs[:k] // pick k items
    sort.Slice(selected, func(i, j int) bool {
        return selected[i].index < selected[j].index
    })
    res := make([]int, k)
    for i := range selected {
        res[i] = selected[i].val
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,1,3,3], k = 2
    // Output: [3,3]
    // Explanation:
    // The subsequence has the largest sum of 3 + 3 = 6.
    fmt.Println(maxSubsequence([]int{2,1,3,3}, 2)) // [3,3]
    // Example 2:
    // Input: nums = [-1,-2,3,4], k = 3
    // Output: [-1,3,4]
    // Explanation: 
    // The subsequence has the largest sum of -1 + 3 + 4 = 6.
    fmt.Println(maxSubsequence([]int{-1,-2,3,4}, 3)) // [-1,3,4]
    // Example 3:
    // Input: nums = [3,4,3,3], k = 2
    // Output: [3,4]
    // Explanation:
    // The subsequence has the largest sum of 3 + 4 = 7. 
    // Another possible subsequence is [4, 3].
    fmt.Println(maxSubsequence([]int{3,4,3,3}, 2)) // [3,4]

    fmt.Println(maxSubsequence([]int{1,2,3,4,5,6,7,8,9}, 2)) // [8 9]
    fmt.Println(maxSubsequence([]int{9,8,7,6,5,4,3,2,1}, 2)) // [9 8]

    fmt.Println(maxSubsequence1([]int{2,1,3,3}, 2)) // [3,3]
    fmt.Println(maxSubsequence1([]int{-1,-2,3,4}, 3)) // [-1,3,4]
    fmt.Println(maxSubsequence1([]int{3,4,3,3}, 2)) // [3,4]
    fmt.Println(maxSubsequence1([]int{1,2,3,4,5,6,7,8,9}, 2)) // [8 9]
    fmt.Println(maxSubsequence1([]int{9,8,7,6,5,4,3,2,1}, 2)) // [9 8]
}