package main

// 3006. Find Beautiful Indices in the Given Array I
// You are given a 0-indexed string s, a string a, a string b, and an integer k.

// An index i is beautiful if:
//     0 <= i <= s.length - a.length
//     s[i..(i + a.length - 1)] == a
//     There exists an index j such that:
//         0 <= j <= s.length - b.length
//         s[j..(j + b.length - 1)] == b
//         |j - i| <= k

// Return the array that contains beautiful indices in sorted order from smallest to largest.

// Example 1:
// Input: s = "isawsquirrelnearmysquirrelhouseohmy", a = "my", b = "squirrel", k = 15
// Output: [16,33]
// Explanation: There are 2 beautiful indices: [16,33].
// - The index 16 is beautiful as s[16..17] == "my" and there exists an index 4 with s[4..11] == "squirrel" and |16 - 4| <= 15.
// - The index 33 is beautiful as s[33..34] == "my" and there exists an index 18 with s[18..25] == "squirrel" and |33 - 18| <= 15.
// Thus we return [16,33] as the result.

// Example 2:
// Input: s = "abcd", a = "a", b = "a", k = 4
// Output: [0]
// Explanation: There is 1 beautiful index: [0].
// - The index 0 is beautiful as s[0..0] == "a" and there exists an index 0 with s[0..0] == "a" and |0 - 0| <= 4.
// Thus we return [0] as the result.

// Constraints:
//     1 <= k <= s.length <= 10^5
//     1 <= a.length, b.length <= 10
//     s, a, and b contain only lowercase English letters.

import "fmt"

func beautifulIndices(s string, a string, b string, k int) []int {
    res := []int{}
    getPatternMatchingIndex := func(s string, a string) []int {
        t := a + "@" + s
        lps := []int{0}
        for i := 1; i < len(t); i++ { 
            index := lps[i - 1]
            for index > 0 && t[index] != t[i] { 
                index = lps[index - 1] 
            }
            if t[index] == t[i] {
                lps = append(lps, index + 1)
            } else {
                lps = append(lps, 0)
            }
        }
        res := []int{}
        for i := 0; i < len(lps); i++ {
            if lps[i] == len(a) {
                res = append(res, i - 2 * len(a))
            }
        }
        return res
    }
    v1, v2 := getPatternMatchingIndex(s, a), getPatternMatchingIndex(s, b)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i, j := 0, 0; i < len(v1); i++ {
        for j < len(v2) && v1[i] > v2[j] && abs(v1[i] - v2[j]) > k {
            j++
        }
        if j < len(v2) && abs(v1[i] - v2[j]) <= k {
            res = append(res, v1[i])
        }
    }
    return res
}

func beautifulIndices1(s string, a string, b string, k int) []int {
    kmp := func(text, pattern string) []int { // 获取pattern在text中的所有出现位置
        n := len(pattern)
        res, next := []int{}, make([]int, n)
        for i, j := 1, 0; i < n; i++ {
            for j > 0 && pattern[j] != pattern[i] {
                j = next[j - 1]
            }
            if pattern[j] == pattern[i] {
                j++
                next[i] = j
            }
        }
        count := 0
        for i, v := range text {
            for count > 0 && pattern[count] != byte(v) {
                count = next[count-1]
            }
            if pattern[count] == byte(v) {
                count++
                if count == n {
                    res = append(res, i - n + 1) // 往回跳 n - 1 格子
                    count = next[count-1] // 后退一格
                }
            }
        }
        return res
    }
    posA, posB := kmp(s, a), kmp(s, b)
    res, j, m := []int{}, 0, len(posB)
    for _, p := range posA {
        for j < m && posB[j] < p - k { // 两个数组都为递增数组,所以只需寻找 >=leftBorder的第一个数,看是否满足 <=rightBorder即可
            j++
        }
        if j < m && posB[j] <= p + k {
            res = append(res, p)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "isawsquirrelnearmysquirrelhouseohmy", a = "my", b = "squirrel", k = 15
    // Output: [16,33]
    // Explanation: There are 2 beautiful indices: [16,33].
    // - The index 16 is beautiful as s[16..17] == "my" and there exists an index 4 with s[4..11] == "squirrel" and |16 - 4| <= 15.
    // - The index 33 is beautiful as s[33..34] == "my" and there exists an index 18 with s[18..25] == "squirrel" and |33 - 18| <= 15.
    // Thus we return [16,33] as the result.
    fmt.Println(beautifulIndices("isawsquirrelnearmysquirrelhouseohmy", "my" ,"squirrel", 15)) // [16,33]
    // Example 2:
    // Input: s = "abcd", a = "a", b = "a", k = 4
    // Output: [0]
    // Explanation: There is 1 beautiful index: [0].
    // - The index 0 is beautiful as s[0..0] == "a" and there exists an index 0 with s[0..0] == "a" and |0 - 0| <= 4.
    // Thus we return [0] as the result.
    fmt.Println(beautifulIndices("abcd", "a" ,"a", 4)) // [0]

    fmt.Println(beautifulIndices1("isawsquirrelnearmysquirrelhouseohmy", "my" ,"squirrel", 15)) // [16,33]
    fmt.Println(beautifulIndices1("abcd", "a" ,"a", 4)) // [0]
}