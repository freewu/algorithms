package main

// 3501. Maximize Active Section with Trade II
// You are given a binary string s of length n, where:
//     '1' represents an active section.
//     '0' represents an inactive section.

// You can perform at most one trade to maximize the number of active sections in s. 
// In a trade, you:
//     1. Convert a contiguous block of '1's that is surrounded by '0's to all '0's.
//     2. Afterward, convert a contiguous block of '0's that is surrounded by '1's to all '1's.

// Additionally, you are given a 2D array queries, where queries[i] = [li, ri] represents a substring s[li...ri].

// For each query, determine the maximum possible number of active sections in s after making the optimal trade on the substring s[li...ri].

// Return an array answer, where answer[i] is the result for queries[i].

// A substring is a contiguous non-empty sequence of characters within a string.

// Note
//     For each query, treat s[li...ri] as if it is augmented with a '1' at both ends, forming t = '1' + s[li...ri] + '1'. The augmented '1's do not contribute to the final count.
//     The queries are independent of each other.

// Example 1:
// Input: s = "01", queries = [[0,1]]
// Output: [1]
// Explanation:
// Because there is no block of '1's surrounded by '0's, no valid trade is possible. The maximum number of active sections is 1.

// Example 2:
// Input: s = "0100", queries = [[0,3],[0,2],[1,3],[2,3]]
// Output: [4,3,1,1]
// Explanation:
// Query [0, 3] → Substring "0100" → Augmented to "101001"
// Choose "0100", convert "0100" → "0000" → "1111".
// The final string without augmentation is "1111". The maximum number of active sections is 4.
// Query [0, 2] → Substring "010" → Augmented to "10101"
// Choose "010", convert "010" → "000" → "111".
// The final string without augmentation is "1110". The maximum number of active sections is 3.
// Query [1, 3] → Substring "100" → Augmented to "11001"
// Because there is no block of '1's surrounded by '0's, no valid trade is possible. The maximum number of active sections is 1.
// Query [2, 3] → Substring "00" → Augmented to "1001"
// Because there is no block of '1's surrounded by '0's, no valid trade is possible. The maximum number of active sections is 1.

// Example 3:
// Input: s = "1000100", queries = [[1,5],[0,6],[0,4]]
// Output: [6,7,2]
// Explanation:
// Query [1, 5] → Substring "00010" → Augmented to "1000101"
// Choose "00010", convert "00010" → "00000" → "11111".
// The final string without augmentation is "1111110". The maximum number of active sections is 6.
// Query [0, 6] → Substring "1000100" → Augmented to "110001001"
// Choose "000100", convert "000100" → "000000" → "111111".
// The final string without augmentation is "1111111". The maximum number of active sections is 7.
// Query [0, 4] → Substring "10001" → Augmented to "1100011"
// Because there is no block of '1's surrounded by '0's, no valid trade is possible. The maximum number of active sections is 2.

// Example 4:
// Input: s = "01010", queries = [[0,3],[1,4],[1,3]]
// Output: [4,4,2]
// Explanation:
// Query [0, 3] → Substring "0101" → Augmented to "101011"
// Choose "010", convert "010" → "000" → "111".
// The final string without augmentation is "11110". The maximum number of active sections is 4.
// Query [1, 4] → Substring "1010" → Augmented to "110101"
// Choose "010", convert "010" → "000" → "111".
// The final string without augmentation is "01111". The maximum number of active sections is 4.
// Query [1, 3] → Substring "101" → Augmented to "11011"
// Because there is no block of '1's surrounded by '0's, no valid trade is possible. The maximum number of active sections is 2.

// Constraints:
//     1 <= n == s.length <= 10^5
//     1 <= queries.length <= 10^5
//     s[i] is either '0' or '1'.
//     queries[i] = [li, ri]
//     0 <= li <= ri < n

import "fmt"
import "sort"
import "math/bits"

func maxActiveSectionsAfterTrade(s string, queries [][]int) []int {
    n := len(s)
    base := 0
    for i := 0; i < n; i++ {
        if s[i] == '1' {
            base++
        }
    }
    // 提取连续段
    var vec [][2]int
    j := 0
    for i := 0; i < n; i++ {
        if i == n-1 || s[i] != s[i+1] {
            vec = append(vec, [2]int{j, i - j + 1})
            j = i + 1
        }
    }
    m := len(vec)
    // RMQ 预处理
    const MAXP = 20
    INF := 1 << 30
    rmq := make([][MAXP]int, m)
    // for i := range rmq {
    //     rmq[i] = make([]int, MAXP)
    // }
    for i := 0; i < m; i++ {
        if s[vec[i][0]] == '1' {
            rmq[i][0] = -INF
        } else {
            if i+2 < m {
                rmq[i][0] = vec[i][1] + vec[i+2][1]
            } else {
                rmq[i][0] = -INF
            }
        }
    }
    for p := 1; (1 << p) <= m; p++ {
        for i := 0; i+(1<<(p-1)) < m; i++ {
            rmq[i][p] = max(rmq[i][p-1], rmq[i+(1<<(p-1))][p-1])
        }
    }
    // 查询函数
    query := func(l, r int) int {
        if l > r {
            return -INF
        }
        k := bits.Len(uint(r-l+1)) - 1
        return max(rmq[l][k], rmq[r-(1<<k)+1][k])
    }
    res := make([]int, len(queries))
    for i, qry := range queries {
        l, r := qry[0], qry[1]
        // 找到被切断的段
        bl := sort.Search(len(vec), func(i int) bool { return vec[i][0] > l }) - 1
        br := sort.Search(len(vec), func(i int) bool { return vec[i][0] > r }) - 1
        if br-bl+1 <= 2 {
            res[i] = base
            continue
        }
        // 计算被切断后的长度
        getNum := func(idx int) int {
            if idx == bl {
                return vec[bl][1] - (l - vec[bl][0])
            }
            if idx == br {
                return r - vec[br][0] + 1
            }
            return vec[idx][1]
        }
        calc := func(idx int) int {
            if s[vec[idx][0]] == '1' {
                return -INF
            }
            return getNum(idx) + getNum(idx+2)
        }
        best := max(query(bl+1, br-3), 0)
        best = max(best, calc(bl))
        best = max(best, calc(br-2))
        res[i] = base + best
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "01", queries = [[0,1]]
    // Output: [1]
    // Explanation:
    // Because there is no block of '1's surrounded by '0's, no valid trade is possible. The maximum number of active sections is 1.
    fmt.Println(maxActiveSectionsAfterTrade("01", [][]int{{0,1}})) // [1]
    // Example 2:
    // Input: s = "0100", queries = [[0,3],[0,2],[1,3],[2,3]]
    // Output: [4,3,1,1]
    // Explanation:
    // Query [0, 3] → Substring "0100" → Augmented to "101001"
    // Choose "0100", convert "0100" → "0000" → "1111".
    // The final string without augmentation is "1111". The maximum number of active sections is 4.
    // Query [0, 2] → Substring "010" → Augmented to "10101"
    // Choose "010", convert "010" → "000" → "111".
    // The final string without augmentation is "1110". The maximum number of active sections is 3.
    // Query [1, 3] → Substring "100" → Augmented to "11001"
    // Because there is no block of '1's surrounded by '0's, no valid trade is possible. The maximum number of active sections is 1.
    // Query [2, 3] → Substring "00" → Augmented to "1001"
    // Because there is no block of '1's surrounded by '0's, no valid trade is possible. The maximum number of active sections is 1.
    fmt.Println(maxActiveSectionsAfterTrade("0100", [][]int{{0,3},{0,2},{1,3},{2,3}})) // [4,3,1,1]
    // Example 3:
    // Input: s = "1000100", queries = [[1,5],[0,6],[0,4]]
    // Output: [6,7,2]
    // Explanation:
    // Query [1, 5] → Substring "00010" → Augmented to "1000101"
    // Choose "00010", convert "00010" → "00000" → "11111".
    // The final string without augmentation is "1111110". The maximum number of active sections is 6.
    // Query [0, 6] → Substring "1000100" → Augmented to "110001001"
    // Choose "000100", convert "000100" → "000000" → "111111".
    // The final string without augmentation is "1111111". The maximum number of active sections is 7.
    // Query [0, 4] → Substring "10001" → Augmented to "1100011"
    // Because there is no block of '1's surrounded by '0's, no valid trade is possible. The maximum number of active sections is 2.
    fmt.Println(maxActiveSectionsAfterTrade("1000100", [][]int{{1,5},{0,6},{0,4}})) // [6,7,2]
    // Example 4:
    // Input: s = "01010", queries = [[0,3],[1,4],[1,3]]
    // Output: [4,4,2]
    // Explanation:
    // Query [0, 3] → Substring "0101" → Augmented to "101011"
    // Choose "010", convert "010" → "000" → "111".
    // The final string without augmentation is "11110". The maximum number of active sections is 4.
    // Query [1, 4] → Substring "1010" → Augmented to "110101"
    // Choose "010", convert "010" → "000" → "111".
    // The final string without augmentation is "01111". The maximum number of active sections is 4.
    // Query [1, 3] → Substring "101" → Augmented to "11011"
    // Because there is no block of '1's surrounded by '0's, no valid trade is possible. The maximum number of active sections is 2.
    fmt.Println(maxActiveSectionsAfterTrade("01010", [][]int{{0,3},{1,4},{1,3}})) // [4,4,2]

    fmt.Println(maxActiveSectionsAfterTrade("0000000000", [][]int{{0,3},{1,4},{1,3}})) // [0 0 0]
    fmt.Println(maxActiveSectionsAfterTrade("1111111111", [][]int{{0,3},{1,4},{1,3}})) // [10 10 10]
    fmt.Println(maxActiveSectionsAfterTrade("0000011111", [][]int{{0,3},{1,4},{1,3}})) // [5 5 5]
    fmt.Println(maxActiveSectionsAfterTrade("1111100000", [][]int{{0,3},{1,4},{1,3}})) // [5 5 5]
    fmt.Println(maxActiveSectionsAfterTrade("0101010101", [][]int{{0,3},{1,4},{1,3}})) // [7 7 5]
    fmt.Println(maxActiveSectionsAfterTrade("1010101010", [][]int{{0,3},{1,4},{1,3}})) // [7 7 7]
}