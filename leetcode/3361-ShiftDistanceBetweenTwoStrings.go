package main

// 3361. Shift Distance Between Two Strings
// You are given two strings s and t of the same length, and two integer arrays nextCost and previousCost.

// In one operation, you can pick any index i of s, and perform either one of the following actions:
//     1. Shift s[i] to the next letter in the alphabet. 
//        If s[i] == 'z', you should replace it with 'a'. 
//        This operation costs nextCost[j] where j is the index of s[i] in the alphabet.
//     2. Shift s[i] to the previous letter in the alphabet. 
//        If s[i] == 'a', you should replace it with 'z'. 
//        This operation costs previousCost[j] where j is the index of s[i] in the alphabet.

// The shift distance is the minimum total cost of operations required to transform s into t.

// Return the shift distance from s to t.

// Example 1:
// Input: s = "abab", t = "baba", nextCost = [100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0], previousCost = [1,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
// Output: 2
// Explanation:
// We choose index i = 0 and shift s[0] 25 times to the previous character for a total cost of 1.
// We choose index i = 1 and shift s[1] 25 times to the next character for a total cost of 0.
// We choose index i = 2 and shift s[2] 25 times to the previous character for a total cost of 1.
// We choose index i = 3 and shift s[3] 25 times to the next character for a total cost of 0.

// Example 2:
// Input: s = "leet", t = "code", nextCost = [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1], previousCost = [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]
// Output: 31
// Explanation:
// We choose index i = 0 and shift s[0] 9 times to the previous character for a total cost of 9.
// We choose index i = 1 and shift s[1] 10 times to the next character for a total cost of 10.
// We choose index i = 2 and shift s[2] 1 time to the previous character for a total cost of 1.
// We choose index i = 3 and shift s[3] 11 times to the next character for a total cost of 11.

// Constraints:
//     1 <= s.length == t.length <= 10^5
//     s and t consist only of lowercase English letters.
//     nextCost.length == previousCost.length == 26
//     0 <= nextCost[i], previousCost[i] <= 10^9

import "fmt"

func shiftDistance(s string, t string, nextCost []int, previousCost []int) int64 {
    res := 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < len(s); i++ {
        if s[i] == t[i] { continue }
        cur, target := int(s[i] - 'a'), int(t[i] - 'a')
        // go forward
        nCost, diff := 0, (target - cur + 26) % 26
        for j := 0; j < diff; j++ {
            nCost += nextCost[(cur + j + 26) % 26]
        }
        // go backward
        pCost, diff := 0, (cur - target + 26) % 26
        for j := 0; j < diff; j++ {
            pCost += previousCost[(cur - j + 26) % 26]
        }
        res += min(nCost, pCost)
    }
    return int64(res)
}

func shiftDistance1(s string, t string, nextCost []int, previousCost []int) int64 {
    next, prev := []int{0}, []int{ previousCost[0] }
    for i := 0; i < len(nextCost); i++{
        next = append(next, next[i] + nextCost[i])
    }
    for i := 1; i < len(previousCost); i++{
        prev = append(prev, prev[i - 1] + previousCost[i])
    }
    min := func (x, y int64) int64 { if x < y { return x; }; return y; }
    res, mx1, mx2 := int64(0), int64(next[26]), int64(prev[25])
    for i := 0; i < len(s); i++{
        sum1, sum2 := int64(next[t[i] - 'a'] - next[s[i] - 'a']), int64(prev[s[i] - 'a'] - prev[t[i] - 'a'])
        if t[i] < s[i] { sum1 = sum1 + mx1 }
        if s[i] < t[i] { sum2 = sum2 + mx2 }
        res += min(sum1, sum2)
    }
    return res
}

func shiftDistance2(s string, t string, nextCost []int, previousCost []int) int64 {
    res, n, m := int64(0), len(s), len(nextCost)
    next, prev := make([]int64, m + 1), make([]int64, m + 1)
    for i := 1; i <= m; i++ {
        next[i], prev[i] = next[i - 1] + int64(nextCost[i - 1]), prev[i - 1] + int64(previousCost[i - 1])
    }
    min := func (x, y int64) int64 { if x < y { return x; }; return y; }
    for i := 0; i < n; i++ {
        if s[i] == t[i] { continue }
        si, ti := int(s[i] - 96), int(t[i] - 96)
        moveNext, movePrev := next[ti - 1] - next[si - 1], prev[si] - prev[ti]
        if si > ti {
            moveNext += next[m]
        } else {
            movePrev += prev[m]
        }
        res += min( moveNext, movePrev )
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abab", t = "baba", nextCost = [100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0], previousCost = [1,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
    // Output: 2
    // Explanation:
    // We choose index i = 0 and shift s[0] 25 times to the previous character for a total cost of 1.
    // We choose index i = 1 and shift s[1] 25 times to the next character for a total cost of 0.
    // We choose index i = 2 and shift s[2] 25 times to the previous character for a total cost of 1.
    // We choose index i = 3 and shift s[3] 25 times to the next character for a total cost of 0.
    fmt.Println(shiftDistance("abab", "baba", []int{100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}, []int{1,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})) // 2 
    // Example 2:
    // Input: s = "leet", t = "code", nextCost = [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1], previousCost = [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]
    // Output: 31
    // Explanation:
    // We choose index i = 0 and shift s[0] 9 times to the previous character for a total cost of 9.
    // We choose index i = 1 and shift s[1] 10 times to the next character for a total cost of 10.
    // We choose index i = 2 and shift s[2] 1 time to the previous character for a total cost of 1.
    // We choose index i = 3 and shift s[3] 11 times to the next character for a total cost of 11.
    fmt.Println(shiftDistance("leet", "code", []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}, []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1})) // 31

    fmt.Println(shiftDistance("blue", "frog", []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}, []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1})) // 18
    fmt.Println(shiftDistance("bluefrog", "leetcode", []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}, []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1})) // 57

    fmt.Println(shiftDistance1("abab", "baba", []int{100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}, []int{1,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})) // 2 
    fmt.Println(shiftDistance1("leet", "code", []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}, []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1})) // 31
    fmt.Println(shiftDistance1("blue", "frog", []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}, []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1})) // 18
    fmt.Println(shiftDistance1("bluefrog", "leetcode", []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}, []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1})) // 57

    fmt.Println(shiftDistance2("abab", "baba", []int{100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}, []int{1,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})) // 2 
    fmt.Println(shiftDistance2("leet", "code", []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}, []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1})) // 31
    fmt.Println(shiftDistance2("blue", "frog", []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}, []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1})) // 18
    fmt.Println(shiftDistance2("bluefrog", "leetcode", []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}, []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1})) // 57
}