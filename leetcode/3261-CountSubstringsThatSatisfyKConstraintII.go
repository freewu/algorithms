package main

// 3261. Count Substrings That Satisfy K-Constraint II
// You are given a binary string s and an integer k.

// You are also given a 2D integer array queries, where queries[i] = [li, ri].

// A binary string satisfies the k-constraint if either of the following conditions holds:
//     The number of 0's in the string is at most k.
//     The number of 1's in the string is at most k.

// Return an integer array answer, where answer[i] is the number of substrings of s[li..ri] that satisfy the k-constraint.

// Example 1:
// Input: s = "0001111", k = 2, queries = [[0,6]]
// Output: [26]
// Explanation:
// For the query [0, 6], all substrings of s[0..6] = "0001111" satisfy the k-constraint except for the substrings s[0..5] = "000111" and s[0..6] = "0001111".

// Example 2:
// Input: s = "010101", k = 1, queries = [[0,5],[1,4],[2,3]]
// Output: [15,9,3]
// Explanation:
// The substrings of s with a length greater than 3 do not satisfy the k-constraint.

// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is either '0' or '1'.
//     1 <= k <= s.length
//     1 <= queries.length <= 10^5
//     queries[i] == [li, ri]
//     0 <= li <= ri < s.length
//     All queries are distinct.

import "fmt"
import "sort"

// // 超出时间限制 620 / 627
// func countKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
//     n, zeros , ones, i , j := len(s), 0, 0, 0, 0
//     lastValid := make([]int, n) // lastValid[i] = min valid index for which substring ending at i satisfy the constraints 
//     for j < n {
//         if s[j] == '1' {
//             ones++
//         } else {
//             zeros++
//         }
//         for zeros > k && ones > k {
//             if s[i] == '1' {
//                 ones--
//             } else {
//                 zeros--
//             }
//             i++
//         }
//         lastValid[j] = i
//         j++ 
//     }
//     res, index := make([]int64, len(queries)), 0
//     for _, q := range queries {
//         curr := 0
//         for j := q[0]; j <= q[1]; j++ {
//             if lastValid[j] <= q[0] { // if newRange start lies in between index j (i.e., substring ending at j) and its min valid index then all the subarrays between will satisfy the constraints 
//                 curr += (j - q[0] + 1)
//             } else { // if newRange start is prior to index j's min vaild index then we know we can only stastify constranits till min valid index for index j
//                 curr += j - lastValid[j] + 1
//             }
//         }
//         res[index] = int64(curr)
//         index++
//     }
//     return res
// }

func countKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
    n, zeros , ones, i , j := len(s), 0, 0, 0, 0
    lastValid := make([]int, n) // lastValid[i] = min valid index for which substring ending at i satisfy the constraints 
    for j < n {
        if s[j] == '1' {
            ones++
        } else {
            zeros++
        }
        for zeros > k && ones > k {
            if s[i] == '1' {
                ones--
            } else {
                zeros--
            }
            i++
        }
        lastValid[j] = i
        j++ 
    }
    prefixSum := make([]int, n)
    prefixSum[0] = 1
    for i := 1; i < n; i++ {
        prefixSum[i] += prefixSum[i - 1] + (i - lastValid[i] + 1)
    }
    getMaxValid := func(lastValid []int,  low,  high, target int) int {
        for low <= high {
            mid := (low + high) / 2
            if  lastValid[mid] <= target {
                low = mid + 1
            } else {
                high = mid - 1
            }
        }
        return low - 1
    }
    res, index := make([]int64, len(queries)), 0
    for _, q := range queries {
        valid := getMaxValid(lastValid, q[0], q[1], q[0])
        res[index] = int64(((valid - q[0] + 1) * (valid - q[0] + 2) / 2) + prefixSum[q[1]] - prefixSum[valid])
        index++
    }
    return res
}

func countKConstraintSubstrings1(s string, k int, queries [][]int) []int64 {
    n, l := len(s), 0
    left, sum, count := make([]int, n), make([]int, n+1), [2]int{}
    for i, v := range s {
        count[v & 1]++
        for count[0] > k && count[1] > k {
            count[s[l] & 1]--
            l++
        }
        left[i] = l
        sum[i+1] = sum[i] + i - l + 1 // 计算 i-left[i]+1 的前缀和
    }
    res := make([]int64, len(queries))
    for i, q := range queries {
        l, r := q[0], q[1]
        j := l + sort.SearchInts(left[l:r+1], l)
        res[i] = int64(sum[r+1] - sum[j] + (j - l + 1) * (j - l) / 2)
    }
    return res
}

func countKConstraintSubstrings2(s string, k int, queries [][]int) []int64 {
    n, l := len(s), 0
    right, sum, count := make([]int, n), make([]int, n + 1), [2]int{}
    for i, v := range s {
        count[v & 1]++
        for count[0] > k && count[1] > k {
            count[s[l] & 1]--
            right[l] = i
            l++
        }
        sum[i+1] = sum[i] + i - l + 1
    }
    for ; l < n; l++ { // 剩余没填的 right[l] 均为 n
        right[l] = n
    }
    res := make([]int64, len(queries))
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i, q := range queries {
        l, r := q[0], q[1]
        j := min(right[l], r + 1)
        res[i] = int64(sum[r+1] - sum[j] + (j - l + 1) * (j - l) / 2)
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "0001111", k = 2, queries = [[0,6]]
    // Output: [26]
    // Explanation:
    // For the query [0, 6], all substrings of s[0..6] = "0001111" satisfy the k-constraint except for the substrings s[0..5] = "000111" and s[0..6] = "0001111".
    fmt.Println(countKConstraintSubstrings("0001111", 2, [][]int{{0,6}})) // [26]
    // Example 2:
    // Input: s = "010101", k = 1, queries = [[0,5],[1,4],[2,3]]
    // Output: [15,9,3]
    // Explanation:
    // The substrings of s with a length greater than 3 do not satisfy the k-constraint.
    fmt.Println(countKConstraintSubstrings("010101", 1, [][]int{{0,5},{1,4},{2,3}})) // [15,9,3]

    fmt.Println(countKConstraintSubstrings1("0001111", 2, [][]int{{0,6}})) // [26]
    fmt.Println(countKConstraintSubstrings1("010101", 1, [][]int{{0,5},{1,4},{2,3}})) // [15,9,3]

    fmt.Println(countKConstraintSubstrings2("0001111", 2, [][]int{{0,6}})) // [26]
    fmt.Println(countKConstraintSubstrings2("010101", 1, [][]int{{0,5},{1,4},{2,3}})) // [15,9,3]
}