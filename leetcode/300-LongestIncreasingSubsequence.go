package main

// 300. Longest Increasing Subsequence
// Given an integer array nums, return the length of the longest strictly increasing subsequence.

// Example 1:
// Input: nums = [10,9,2,5,3,7,101,18]
// Output: 4
// Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

// Example 2:
// Input: nums = [0,1,0,3,2,3]
// Output: 4

// Example 3:
// Input: nums = [7,7,7,7,7,7,7]
// Output: 1

// Constraints:
//     1 <= nums.length <= 2500
//     -10^4 <= nums[i] <= 10^4

// Follow up: Can you come up with an algorithm that runs in O(n log(n)) time complexity?
// 给定一个无序的整数数组，找到其中最长上升子序列的长度

import "fmt"
import "sort"

// O(n^2) DP
func lengthOfLIS(nums []int) int {
    dp, res := make([]int, len(nums)+1), 0
    dp[0] = 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= len(nums); i++ {
        for j := 1; j < i; j++ {
            if nums[j-1] < nums[i-1] {
                dp[i] = max(dp[i], dp[j])
            }
        }
        dp[i] = dp[i] + 1
        res = max(res, dp[i])
    }
    return res
}

//  O(nlogn) DP
func lengthOfLIS1(nums []int) int {
    dp := []int{}
    for _, num := range nums {
        i := sort.SearchInts(dp, num)
        // fmt.Printf("i = %v %v\n", i, num)
        if i == len(dp) {
            //fmt.Printf("append before %v %v\n", dp, num)
            dp = append(dp, num)
            //fmt.Printf("append after %v %v\n", dp, num)
        } else {
            dp[i] = num
        }
    }
    return len(dp)
}

// 二分 + 贪心
func lengthOfLIS2(nums []int) int { 
    dp := make([]int, 0)
    twoDivide := func (nums []int, target int) int {
        l := 0
        r := len(nums) - 1
        for l <= r {
            mid := l + (r-l)/2
            if nums[mid] >= target {
                r = mid - 1
            } else {
                l = mid + 1
            }
        }
        return l
    }
    for _, v := range nums {
        l := twoDivide(dp, v)
        if l == len(dp) { // 当前元素 v 是最大的元素
            dp = append(dp, v)
        } else { // 当前元素 v 可以替换原有的元素
            dp[l] = v
        }
    }
    return len(dp)
}

func main() {
    // Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
    fmt.Printf("lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}) = %v\n", lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4 [2,3,7,101]
    fmt.Printf("lengthOfLIS([]int{0,1,0,3,2,3}) = %v\n", lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))                          // 4 [0,0,2,3]
    fmt.Printf("lengthOfLIS([]int{7,7,7,7,7,7,7}) = %v\n", lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))                     // 1

    fmt.Printf("lengthOfLIS1([]int{10, 9, 2, 5, 3, 7, 101, 18}) = %v\n", lengthOfLIS1([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4 [2,3,7,101]
    fmt.Printf("lengthOfLIS1([]int{0,1,0,3,2,3}) = %v\n", lengthOfLIS1([]int{0, 1, 0, 3, 2, 3}))                          // 4 [0,0,2,3]
    fmt.Printf("lengthOfLIS1([]int{7,7,7,7,7,7,7}) = %v\n", lengthOfLIS1([]int{7, 7, 7, 7, 7, 7, 7}))                    // 1
    
    fmt.Printf("lengthOfLIS2([]int{10, 9, 2, 5, 3, 7, 101, 18}) = %v\n", lengthOfLIS2([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4 [2,3,7,101]
    fmt.Printf("lengthOfLIS2([]int{0,1,0,3,2,3}) = %v\n", lengthOfLIS2([]int{0, 1, 0, 3, 2, 3}))                          // 4 [0,0,2,3]
    fmt.Printf("lengthOfLIS2([]int{7,7,7,7,7,7,7}) = %v\n", lengthOfLIS2([]int{7, 7, 7, 7, 7, 7, 7}))                    // 1
}
