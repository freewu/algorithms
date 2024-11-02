package main

// 1713. Minimum Operations to Make a Subsequence
// You are given an array target that consists of distinct integers 
// and another integer array arr that can have duplicates.

// In one operation, you can insert any integer at any position in arr. 
// For example, if arr = [1,4,1,2], you can add 3 in the middle and make it [1,4,3,1,2]. 
// Note that you can insert the integer at the very beginning or end of the array.

// Return the minimum number of operations needed to make target a subsequence of arr.

// A subsequence of an array is a new array generated 
// from the original array by deleting some elements (possibly none) without changing the remaining elements' relative order. 
// For example, [2,7,4] is a subsequence of [4,2,3,7,2,1,4] (the underlined elements), while [2,4,2] is not.

// Example 1:
// Input: target = [5,1,3], arr = [9,4,2,3,4]
// Output: 2
// Explanation: You can add 5 and 1 in such a way that makes arr = [5,9,4,1,2,3,4], then target will be a subsequence of arr.

// Example 2:
// Input: target = [6,4,8,1,3,2], arr = [4,7,6,2,3,8,6,1]
// Output: 3

// Constraints:
//     1 <= target.length, arr.length <= 10^5
//     1 <= target[i], arr[i] <= 10^9
//     target contains no duplicates.

import "fmt"
import "sort"

func minOperations(target []int, arr []int) int {
    mp := make(map[int]int)
    for i, v := range target {
        mp[v] = i
    }
    for i, v := range arr {
        if val, ok := mp[v]; ok {
            arr[i] = val
        } else {
            arr[i] = -1
        }
    }
    dp := make([]int,0)
    for _, v := range arr {
        if v == -1 { continue }
        if len(dp) == 0 {
           dp = append(dp,v)
        } else if v > dp[len(dp) - 1] {
            dp = append(dp, v)
        } else {
            j := sort.Search(len(dp), func (i int) bool {
                return dp[i] >= v
            })
            dp[j] = v
        }
    }
    return len(target) - len(dp)
}

func main() {
    // Example 1:
    // Input: target = [5,1,3], arr = [9,4,2,3,4]
    // Output: 2
    // Explanation: You can add 5 and 1 in such a way that makes arr = [5,9,4,1,2,3,4], then target will be a subsequence of arr.
    fmt.Println(minOperations([]int{5,1,3}, []int{9,4,2,3,4})) // 2
    // Example 2:
    // Input: target = [6,4,8,1,3,2], arr = [4,7,6,2,3,8,6,1]
    // Output: 3
    fmt.Println(minOperations([]int{6,4,8,1,3,2}, []int{4,7,6,2,3,8,6,1})) // 3
}