package main

// 2638. Count the Number of K-Free Subsets
// You are given an integer array nums, which contains distinct elements and an integer k.

// A subset is called a k-Free subset if it contains no two elements with an absolute difference equal to k. 
// Notice that the empty set is a k-Free subset.

// Return the number of k-Free subsets of nums.

// A subset of an array is a selection of elements (possibly none) of the array.

// Example 1:
// Input: nums = [5,4,6], k = 1
// Output: 5
// Explanation: There are 5 valid subsets: {}, {5}, {4}, {6} and {4, 6}.

// Example 2:
// Input: nums = [2,3,5,8], k = 5
// Output: 12
// Explanation: There are 12 valid subsets: {}, {2}, {3}, {5}, {8}, {2, 3}, {2, 3, 5}, {2, 5}, {2, 5, 8}, {2, 8}, {3, 5} and {5, 8}.

// Example 3:
// Input: nums = [10,5,9,11], k = 20
// Output: 16
// Explanation: All subsets are valid. Since the total count of subsets is 24 = 16, so the answer is 16. 

// Constraints:
//     1 <= nums.length <= 50
//     1 <= nums[i] <= 1000
//     1 <= k <= 1000

import "fmt"
import "sort"

func countTheNumOfKFreeSubsets(nums []int, k int) int64 {
    sort.Ints(nums)
    res, mp := 1, make(map[int][]int)
    for _, v := range nums {
        mp[v % k] = append(mp[v % k], v)
    }
    for _, arr := range mp {
        n := len(arr)
        dp := make([]int, n + 1)
        dp[0], dp[1] = 1, 2
        for i := 2; i <= n; i++ {
            if arr[i - 1] - arr[i - 2] == k {
                dp[i] = dp[i-1] + dp[i-2]
            } else {
                dp[i] = dp[i-1] * 2
            }
        }
        res *= dp[n]
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [5,4,6], k = 1
    // Output: 5
    // Explanation: There are 5 valid subsets: {}, {5}, {4}, {6} and {4, 6}.
    fmt.Println(countTheNumOfKFreeSubsets([]int{5,4,6}, 1)) // 5
    // Example 2:
    // Input: nums = [2,3,5,8], k = 5
    // Output: 12
    // Explanation: There are 12 valid subsets: {}, {2}, {3}, {5}, {8}, {2, 3}, {2, 3, 5}, {2, 5}, {2, 5, 8}, {2, 8}, {3, 5} and {5, 8}.
    fmt.Println(countTheNumOfKFreeSubsets([]int{2,3,5,8}, 5)) // 12
    // Example 3:
    // Input: nums = [10,5,9,11], k = 20
    // Output: 16
    // Explanation: All subsets are valid. Since the total count of subsets is 24 = 16, so the answer is 16. 
    fmt.Println(countTheNumOfKFreeSubsets([]int{10,5,9,11}, 20)) // 16
}