package main

// 3176. Find the Maximum Length of a Good Subsequence I
// You are given an integer array nums and a non-negative integer k. 
// A sequence of integers seq is called good if there are at most k indices i in the range [0, seq.length - 2] such that seq[i] != seq[i + 1].

// Return the maximum possible length of a good subsequence of nums.

// Example 1:
// Input: nums = [1,2,1,1,3], k = 2
// Output: 4
// Explanation:
// The maximum length subsequence is [1,2,1,1,3].

// Example 2:
// Input: nums = [1,2,3,4,5,1], k = 0
// Output: 2
// Explanation:
// The maximum length subsequence is [1,2,3,4,5,1].

// Constraints:
//     1 <= nums.length <= 500
//     1 <= nums[i] <= 10^9
//     0 <= k <= min(nums.length, 25)

import "fmt"

func maximumLength(nums []int, k int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, n, mn := 1, len(nums), min(k + 1, 51)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, mn)
        for j := range dp[i] {
            dp[i][j] = 1 // fill 1
        }
    }
    for i := 1; i < n; i++ {
        for j := 0; j < min(k, 50) + 1; j++ {
            for k := 0; k < i; k++ {
                if nums[k] == nums[i] {
                    dp[i][j] = max(dp[i][j], dp[k][j] + 1)
                } else if j > 0 {
                    dp[i][j] = max(dp[i][j], dp[k][j - 1] + 1)
                }
            }
            res = max(res, dp[i][j])
        }
    }
    return res
}

// class Solution:
//     def maximumLength(self, nums: List[int], k: int) -> int:
//         n = len(nums)
//         dp = [[1] * min(k + 1, 51) for _ in range(n)]

//         max_length = 1
        
//         for i in range(1, n):
//             for j in range(min(k, 50) + 1):
//                 for m in range(i):
//                     if nums[m] == nums[i]:
//                         dp[i][j] = max(dp[i][j], dp[m][j] + 1)
//                     elif j > 0:
//                         dp[i][j] = max(dp[i][j], dp[m][j - 1] + 1)
                
//                 max_length = max(max_length, dp[i][j])
        
//         return max_length


func maximumLength1(nums []int, k int) int {
    fs, records := map[int][]int{}, make([]struct{ mx, mx2, num int }, k+1)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, x := range nums {
        if fs[x] == nil {
            fs[x] = make([]int, k+1)
        }
        f := fs[x]
        for j := k; j >= 0; j-- {
            f[j]++
            if j > 0 {
                p := records[j-1]
                m := p.mx
                if x == p.num {
                    m = p.mx2
                }
                f[j] = max(f[j], m+1)
            }
            // records[j] 维护 fs[.][j] 的 mx,mx2,num
            v := f[j]
            p := &records[j]
            if v > p.mx {
                if x != p.num {
                    p.num = x
                    p.mx2 = p.mx
                }
                p.mx = v
            } else if x != p.num && v > p.mx2 {
                p.mx2 = v
            }
        }
    }
    return records[k].mx
}

func main() {
    // Example 1:
    // Input: nums = [1,2,1,1,3], k = 2
    // Output: 4
    // Explanation:
    // The maximum length subsequence is [1,2,1,1,3].
    fmt.Println(maximumLength([]int{1,2,1,1,3}, 2)) // 4
    // Input: nums = [1,2,3,4,5,1], k = 0
    // Output: 2
    // Explanation:
    // The maximum length subsequence is [1,2,3,4,5,1].
    fmt.Println(maximumLength([]int{1,2,3,4,5,1}, 0)) // 2

    fmt.Println(maximumLength1([]int{1,2,1,1,3}, 2)) // 4
    fmt.Println(maximumLength1([]int{1,2,3,4,5,1}, 0)) // 2
}