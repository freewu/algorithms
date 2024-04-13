package main

// 1218. Longest Arithmetic Subsequence of Given Difference
// Given an integer array arr and an integer difference, 
// return the length of the longest subsequence in arr which is an arithmetic sequence such 
// that the difference between adjacent elements in the subsequence equals difference.

// A subsequence is a sequence that can be derived from arr by deleting some or no elements without changing the order of the remaining elements.

// Example 1:
// Input: arr = [1,2,3,4], difference = 1
// Output: 4
// Explanation: The longest arithmetic subsequence is [1,2,3,4].

// Example 2:
// Input: arr = [1,3,5,7], difference = 1
// Output: 1
// Explanation: The longest arithmetic subsequence is any single element.

// Example 3:
// Input: arr = [1,5,7,8,5,3,4,2,1], difference = -2
// Output: 4
// Explanation: The longest arithmetic subsequence is [7,5,3,1].
 
// Constraints:
//     1 <= arr.length <= 10^5
//     -10^4 <= arr[i], difference <= 10^4

import "fmt"

func longestSubsequence(arr []int, difference int) int {
    mp, res := make(map[int]int), 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := range arr {
        if v, ok := mp[arr[i] - difference]; ok {
            res = max(res, v + 1)
            mp[arr[i]] = v + 1
        } else {
            res = max(res, 1)
            mp[arr[i]] = 1
        }
    }
    return res
}

func longestSubsequence1(arr []int, difference int) int {
    res, k, f := 1, 100000, make([]int, 200000)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(arr); i ++ {
        f[arr[i] + k] = max(f[arr[i] + k - difference] + 1, f[arr[i] + k]) 
        res = max(res, f[arr[i] + k])
    }
    return res
}

func main() {
    // Explanation: The longest arithmetic subsequence is [1,2,3,4].
    fmt.Println(longestSubsequence([]int{1,2,3,4},1)) // 4
    // Explanation: The longest arithmetic subsequence is any single element.
    fmt.Println(longestSubsequence([]int{1,3,5,7},1)) // 1
    // Explanation: The longest arithmetic subsequence is [7,5,3,1].
    fmt.Println(longestSubsequence([]int{1,5,7,8,5,3,4,2,1},-2)) // 4

    fmt.Println(longestSubsequence1([]int{1,2,3,4},1)) // 4
    fmt.Println(longestSubsequence1([]int{1,3,5,7},1)) // 1
    fmt.Println(longestSubsequence1([]int{1,5,7,8,5,3,4,2,1},-2)) // 4
}