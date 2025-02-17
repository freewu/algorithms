package main

// 3020. Find the Maximum Number of Elements in Subset
// You are given an array of positive integers nums.

// You need to select a subset of nums which satisfies the following condition:
//     1. You can place the selected elements in a 0-indexed array such that it follows the pattern: 
//        [x, x2, x4, ..., xk/2, xk, xk/2, ..., x4, x2, x] (Note that k can be be any non-negative power of 2). 
//        For example, [2, 4, 16, 4, 2] and [3, 9, 3] follow the pattern while [2, 4, 8, 4, 2] does not.

// Return the maximum number of elements in a subset that satisfies these conditions.

// Example 1:
// Input: nums = [5,4,1,2,2]
// Output: 3
// Explanation: We can select the subset {4,2,2}, which can be placed in the array as [2,4,2] which follows the pattern and 22 == 4. Hence the answer is 3.

// Example 2:
// Input: nums = [1,3,2,4]
// Output: 1
// Explanation: We can select the subset {1}, which can be placed in the array as [1] which follows the pattern. Hence the answer is 1. Note that we could have also selected the subsets {2}, {3}, or {4}, there may be multiple subsets which provide the same answer. 

// Constraints:
//     2 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"
import "sort"

func maximumLength(nums []int) int {
    res := 1
    freq, dp := make(map[int]int), make(map[int]int)
    for _, v := range nums { 
        freq[v]++ 
    }
    if freq[1] % 2 == 0 { // let's handle all ones
        res = freq[1] - 1
    } else {
        res = freq[1]
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    sort.Ints(nums)
    for i := len(nums) - 1; i >= 0; i-- {
        if freq[nums[i]] >= 1 {
            if dp[nums[i]*nums[i]] > 0 && freq[nums[i]] >= 2 {
                dp[nums[i]] += (dp[nums[i] * nums[i]] + 2)
            } else {
                dp[nums[i]]++
            }
            freq[nums[i]] = 0
        }
        res = max(res, dp[nums[i]])
    }
    return res
}

func maximumLength1(nums []int) int {
    count := make(map[int]int)
    for _, v := range nums{
        count[v]++
    }
    res := count[1] - 1 | 1
    delete(count, 1)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for v := range count {
        val := 0
        for ; count[v] >= 2; v *= v {
            val += 2
        }
        if count[v] == 1 { // 正好
            val++
        } else {  // 根号x +多了
            val--
        }
        res = max(res, val) // 保证 res 是奇数
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [5,4,1,2,2]
    // Output: 3
    // Explanation: We can select the subset {4,2,2}, which can be placed in the array as [2,4,2] which follows the pattern and 22 == 4. Hence the answer is 3.
    fmt.Println(maximumLength([]int{5,4,1,2,2})) // 3
    // Example 2:
    // Input: nums = [1,3,2,4]
    // Output: 1
    // Explanation: We can select the subset {1}, which can be placed in the array as [1] which follows the pattern. Hence the answer is 1. Note that we could have also selected the subsets {2}, {3}, or {4}, there may be multiple subsets which provide the same answer. 
    fmt.Println(maximumLength([]int{1,3,2,4})) // 1

    fmt.Println(maximumLength([]int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(maximumLength([]int{9,8,7,6,5,4,3,2,1})) // 1

    fmt.Println(maximumLength1([]int{5,4,1,2,2})) // 3
    fmt.Println(maximumLength1([]int{1,3,2,4})) // 1
    fmt.Println(maximumLength1([]int{1,2,3,4,5,6,7,8,9})) // 1
    fmt.Println(maximumLength1([]int{9,8,7,6,5,4,3,2,1})) // 1
}