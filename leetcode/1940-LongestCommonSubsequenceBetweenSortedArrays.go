package main

// 1940. Longest Common Subsequence Between Sorted Arrays
// Given an array of integer arrays arrays where each arrays[i] is sorted in strictly increasing order, 
// return an integer array representing the longest common subsequence among all the arrays.

// A subsequence is a sequence that can be derived from another sequence 
// by deleting some elements (possibly none) without changing the order of the remaining elements.

// Example 1:
// Input: arrays = [[1,3,4],
//                  [1,4,7,9]]
// Output: [1,4]
// Explanation: The longest common subsequence in the two arrays is [1,4].

// Example 2:
// Input: arrays = [[2,3,6,8],
//                  [1,2,3,5,6,7,10],
//                  [2,3,4,6,9]]
// Output: [2,3,6]
// Explanation: The longest common subsequence in all three arrays is [2,3,6].

// Example 3:
// Input: arrays = [[1,2,3,4,5],
//                  [6,7,8]]
// Output: []
// Explanation: There is no common subsequence between the two arrays.

// Constraints:
//     2 <= arrays.length <= 100
//     1 <= arrays[i].length <= 100
//     1 <= arrays[i][j] <= 100
//     arrays[i] is sorted in strictly increasing order.

import "fmt"
import "sort"

func longestCommonSubsequence(arrays [][]int) []int {
    res, mp := []int{}, make(map[int]bool)
    for _, v := range arrays[0] {
        mp[v] = true
    }
    for _, rows := range arrays[1:] {
        mp1 := make(map[int]bool)
        for _, v := range rows {
            if mp[v] {
                mp1[v] = true
            }
        }
        if len(mp1) == 0 { return res }
        mp = mp1
    }
    for i := range mp {
        res = append(res, i)
    }
    sort.Ints(res)
    return res
}

func longestCommonSubsequence1(arrays [][]int) []int { // 【 计数解法: 】
    n := len(arrays) 
    res, count := []int{}, make([]int, 101) 
    for _, rows := range arrays {
        for _, v := range rows {
                count[v] += 1 
        }
    }
    for i := 1; i <= 100; i++ {
        if count[i] == n {
            res = append(res, i) 
        }
    }
    return res 
}

func main() {
    // Example 1:
    // Input: arrays = [[1,3,4],
    //                  [1,4,7,9]]
    // Output: [1,4]
    // Explanation: The longest common subsequence in the two arrays is [1,4].
    arr1 := [][]int{
        {1,3,4},
        {1,4,7,9},
    }
    fmt.Println(longestCommonSubsequence(arr1)) // [1,4]
    // Example 2:
    // Input: arrays = [[2,3,6,8],
    //                  [1,2,3,5,6,7,10],
    //                  [2,3,4,6,9]]
    // Output: [2,3,6]
    // Explanation: The longest common subsequence in all three arrays is [2,3,6].
    arr2 := [][]int{
        {2,3,6,8},
        {1,2,3,5,6,7,10},
        {2,3,4,6,9},
    }
    fmt.Println(longestCommonSubsequence(arr2)) // [2,3,6]
    // Example 3:
    // Input: arrays = [[1,2,3,4,5],
    //                  [6,7,8]]
    // Output: []
    // Explanation: There is no common subsequence between the two arrays.
    arr3 := [][]int{
        {1,2,3,4,5},
        {6,7,8},
    }
    fmt.Println(longestCommonSubsequence(arr3)) // [2 3 6]

    arr4 := [][]int{
        {2,3,6,8},
        {1,2,3,5,6,7,10},
        {2,3,4,6,9},
    }
    fmt.Println(longestCommonSubsequence(arr4)) // []


    fmt.Println(longestCommonSubsequence1(arr1)) // [1,4]
    fmt.Println(longestCommonSubsequence1(arr2)) // [2,3,6]
    fmt.Println(longestCommonSubsequence1(arr3)) // [2 3 6]
    fmt.Println(longestCommonSubsequence1(arr4)) // []
}