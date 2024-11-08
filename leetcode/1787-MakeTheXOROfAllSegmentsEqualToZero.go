package main

// 1787. Make the XOR of All Segments Equal to Zero
// You are given an array nums​​​ and an integer k​​​​​. 
// The XOR of a segment [left, right] where left <= right is the XOR of all the elements with indices between left and right, 
// inclusive: nums[left] XOR nums[left+1] XOR ... XOR nums[right].

// Return the minimum number of elements to change in the array 
// such that the XOR of all segments of size k​​​​​​ is equal to zero.

// Example 1:
// Input: nums = [1,2,0,3,0], k = 1
// Output: 3
// Explanation: Modify the array from [1,2,0,3,0] to from [0,0,0,0,0].

// Example 2:
// Input: nums = [3,4,5,2,1,7,3,4,7], k = 3
// Output: 3
// Explanation: Modify the array from [3,4,5,2,1,7,3,4,7] to [3,4,7,3,4,7,3,4,7].

// Example 3:
// Input: nums = [1,2,4,1,2,5,1,2,6], k = 3
// Output: 3
// Explanation: Modify the array from [1,2,4,1,2,5,1,2,6] to [1,2,3,1,2,3,1,2,3].

// Constraints:
//     1 <= k <= nums.length <= 2000
//     ​​​​​​0 <= nums[i] < 2^10

import "fmt"

func minChanges(nums []int, k int) int {
    n, inf := len(nums), 1 << 31
    freq, count := make([]int, k), make([][1024]int, k)
    for i := 0; i < n; i++ {
        freq[i % k]++
        count[i % k][nums[i]]++
    }
    dp := make([]int, 1024)
    for i := range dp { // fill inf
        dp[i] = inf
    }
    dp[0] = 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < k; i++ {
        ndp, mn := make([]int, 1024), inf
        for j := 0; j < 1024; j++ {
            mn = min(mn, dp[j])
        }
        v := mn + freq[i]
        for j := range ndp {
            ndp[j] = v
        }
        for j := 0; j < 1024; j++ {
            if count[i][j] > 0 {
                for x := 0; x < 1024; x++ {
                    ndp[x ^ j] = min(ndp[x ^ j], dp[x] + freq[i] - count[i][j])
                }
            }
        }
        dp = ndp
    }
    return dp[0]
}

func main() {
    // Example 1:
    // Input: nums = [1,2,0,3,0], k = 1
    // Output: 3
    // Explanation: Modify the array from [1,2,0,3,0] to from [0,0,0,0,0].
    fmt.Println(minChanges([]int{1,2,0,3,0}, 1)) // 3
    // Example 2:
    // Input: nums = [3,4,5,2,1,7,3,4,7], k = 3
    // Output: 3
    // Explanation: Modify the array from [3,4,5,2,1,7,3,4,7] to [3,4,7,3,4,7,3,4,7].
    fmt.Println(minChanges([]int{3,4,5,2,1,7,3,4,7}, 3)) // 3
    // Example 3:
    // Input: nums = [1,2,4,1,2,5,1,2,6], k = 3
    // Output: 3
    // Explanation: Modify the array from [1,2,4,1,2,5,1,2,6] to [1,2,3,1,2,3,1,2,3].
    fmt.Println(minChanges([]int{1,2,4,1,2,5,1,2,6}, 3)) // 3
}