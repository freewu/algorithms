package main

// 1027. Longest Arithmetic Subsequence
// Given an array nums of integers, return the length of the longest arithmetic subsequence in nums.
// Note that:
//     A subsequence is an array that can be derived from another array by deleting some or no elements without changing the order of the remaining elements.
//     A sequence seq is arithmetic if seq[i + 1] - seq[i] are all the same value (for 0 <= i < seq.length - 1).
    
// Example 1:
// Input: nums = [3,6,9,12]
// Output: 4
// Explanation:  The whole array is an arithmetic sequence with steps of length = 3.

// Example 2:
// Input: nums = [9,4,7,2,10]
// Output: 3
// Explanation:  The longest arithmetic subsequence is [4,7,10].

// Example 3:
// Input: nums = [20,1,15,3,10,5,8]
// Output: 4
// Explanation:  The longest arithmetic subsequence is [20,15,10,5].
 
// Constraints:
//     2 <= nums.length <= 1000
//     0 <= nums[i] <= 500

import "fmt"

func longestArithSeqLength(nums []int) int {
    mx,res := 0, 1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // 找出最大值
    for _, v := range nums {
        mx = max(mx, v)
    }
    for k := -mx; k <= mx; k++ { // check every possible diff 
        m := make([]int, 1501) // 0 <= nums[i] <= 500
        m[nums[0] + 500] = 1
        for i := 1 ; i < len(nums); i++ { // LIS for fixed diff value
            x := nums[i] - k 
            m[nums[i] + 500] = max(m[nums[i] + 500], m[x + 500] + 1)
            res = max(res, m[nums[i]+500])
        } 
    }
    return res
}

// dp
func longestArithSeqLength1(nums []int) int {
    res, n := 0, len(nums)
    dp := make([][1001]int, n)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        for j := i- 1; j >= 0; j-- {
            gap := nums[j] - nums[i] + 500 // 计算出间隔
            if dp[i][gap] == 0 {
                dp[i][gap] = dp[j][gap] + 1
                res = max(res, dp[i][gap])
            }
        }
    }
    return res + 1
}

func main() {
    // Explanation:  The whole array is an arithmetic sequence with steps of length = 3.
    fmt.Println(longestArithSeqLength([]int{3,6,9,12})) // 4
    // Explanation:  The longest arithmetic subsequence is [4,7,10].
    fmt.Println(longestArithSeqLength([]int{9,4,7,2,10})) // 3
    // Explanation:  The longest arithmetic subsequence is [20,15,10,5].
    fmt.Println(longestArithSeqLength([]int{20,1,15,3,10,5,8})) // 4

    fmt.Println(longestArithSeqLength1([]int{3,6,9,12})) // 4
    fmt.Println(longestArithSeqLength1([]int{9,4,7,2,10})) // 3
    fmt.Println(longestArithSeqLength1([]int{20,1,15,3,10,5,8})) // 4
}