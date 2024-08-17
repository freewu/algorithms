package main

// 2055. Plates Between Candles
// There is a long table with a line of plates and candles arranged on top of it. 
// You are given a 0-indexed string s consisting of characters '*' and '|' only, where a '*' represents a plate and a '|' represents a candle.

// You are also given a 0-indexed 2D integer array queries where queries[i] = [lefti, righti] denotes the substring s[lefti...righti] (inclusive). 
// For each query, you need to find the number of plates between candles that are in the substring. 
// A plate is considered between candles if there is at least one candle to its left and at least one candle to its right in the substring.
//     For example, s = "||**||**|*", and a query [3, 8] denotes the substring "*||**|". 
//     The number of plates between candles in this substring is 2, as each of the two plates has at least one candle in the substring to its left and right.

// Return an integer array answer where answer[i] is the answer to the ith query.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/10/04/ex-1.png" />
// Input: s = "**|**|***|", queries = [[2,5],[5,9]]
// Output: [2,3]
// Explanation:
// - queries[0] has two plates between candles.
// - queries[1] has three plates between candles.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/10/04/ex-2.png" />
// Input: s = "***|**|*****|**||**|*", queries = [[1,17],[4,5],[14,17],[5,11],[15,16]]
// Output: [9,0,0,0,0]
// Explanation:
// - queries[0] has nine plates between candles.
// - The other queries have zero plates between candles.

// Constraints:
//     3 <= s.length <= 10^5
//     s consists of '*' and '|' characters.
//     1 <= queries.length <= 10^5
//     queries[i].length == 2
//     0 <= lefti <= righti < s.length

import "fmt"

func platesBetweenCandles(s string, queries [][]int) []int {
    prefixPlates := make([]int, len(s)+1)
    prefixPlates[0] = 0
    candles := make([]int, 0, len(s))
    for i := 1; i < len(prefixPlates); i++ {
        prefixPlates[i] = prefixPlates[i-1]
        if s[i-1] == '*' {
            prefixPlates[i]++
        } else {
            candles = append(candles, i-1)
        }
    }
    lessOrEqual := func(candles []int, idx int) int {
        if len(candles) == 0 { return -1 }
        l, r := 0, len(candles)-1
        for l+1 < r {
            mid := l + (r-l)/2
            if candles[mid] < idx {
                l = mid
            } else {
                r = mid
            }
        }
        if candles[r] <= idx { return candles[r] }
        if candles[l] <= idx { return candles[l] }
        return -1
    }
    upperOrEqual := func(candles []int, idx int) int {
        if len(candles) == 0 { return -1 }
        l, r := 0, len(candles)-1
        for l+1 < r {
            mid := l + (r-l)/2
            if candles[mid] <= idx {
                l = mid
            } else {
                r = mid
            }
        }
        if candles[l] >= idx { return candles[l] }
        if candles[r] >= idx { return candles[r] }
        return -1
    }
    res := make([]int, 0, len(queries))
    for _, query := range queries {
        left, right := upperOrEqual(candles, query[0]), lessOrEqual(candles, query[1])
        if right == -1 || left == -1 || right <= left {
            res = append(res, 0)
        } else {
            res = append(res, prefixPlates[right] - prefixPlates[left])
        }
    }
    return res
}

func platesBetweenCandles1(s string, queries [][]int) []int {
    res := make([]int,len(queries))
    sum, left, right := make([]int,len(s)+1), make([]int,len(s)), make([]int,len(s))
    p :=-1
    for i, c := range s {
        sum[i+1] = sum[i]
        if c == '|' {
            p = i
        } else {
            sum[i+1]++
        }
        left[i]=p
    }
    for i, p :=len(s)-1, len(s); i >= 0; i-- {
        if s[i]=='|' {
            p = i
        }
        right[i]=p
    }
    for i, q := range queries {
        l, r := right[q[0]], left[q[1]]
        if l < r {
            res[i] = sum[r] - sum[l]
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/10/04/ex-1.png" />
    // Input: s = "**|**|***|", queries = [[2,5],[5,9]]
    // Output: [2,3]
    // Explanation:
    // - queries[0] has two plates between candles.
    // - queries[1] has three plates between candles.
    fmt.Println(platesBetweenCandles("**|**|***|", [][]int{{2,5},{5,9}})) // [2,3]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/10/04/ex-2.png" />
    // Input: s = "***|**|*****|**||**|*", queries = [[1,17],[4,5],[14,17],[5,11],[15,16]]
    // Output: [9,0,0,0,0]
    // Explanation:
    // - queries[0] has nine plates between candles.
    // - The other queries have zero plates between candles.
    fmt.Println(platesBetweenCandles("***|**|*****|**||**|*", [][]int{{1,17},{4,5},{14,17},{5,11},{15,16}})) // [9,0,0,0,0]
    
    fmt.Println(platesBetweenCandles1("**|**|***|", [][]int{{2,5},{5,9}})) // [2,3]
    fmt.Println(platesBetweenCandles1("***|**|*****|**||**|*", [][]int{{1,17},{4,5},{14,17},{5,11},{15,16}})) // [9,0,0,0,0]
}