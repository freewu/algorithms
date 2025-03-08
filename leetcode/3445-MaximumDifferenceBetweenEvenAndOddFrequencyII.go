package main

// 3445. Maximum Difference Between Even and Odd Frequency II
// You are given a string s and an integer k. 
// Your task is to find the maximum difference between the frequency of two characters, freq[a] - freq[b], in a substring subs of s, such that:
//     1. subs has a size of at least k.
//     2. Character a has an odd frequency in subs.
//     3. Character b has an even frequency in subs.

// Return the maximum difference.

// Note that subs can contain more than 2 distinct characters.

// Example 1:
// Input: s = "12233", k = 4
// Output: -1
// Explanation:
// For the substring "12233", the frequency of '1' is 1 and the frequency of '3' is 2. The difference is 1 - 2 = -1.

// Example 2:
// Input: s = "1122211", k = 3
// Output: 1
// Explanation:
// For the substring "11222", the frequency of '2' is 3 and the frequency of '1' is 2. The difference is 3 - 2 = 1.

// Example 3:
// Input: s = "110", k = 3
// Output: -1

// Constraints:
//     3 <= s.length <= 3 * 10^4
//     s consists only of digits '0' to '4'.
//     The input is generated that at least one substring has a character with an even frequency and a character with an odd frequency.
//     1 <= k <= s.length

import "fmt"

func maxDifference(s string, k int) int {
    res, inf := -1 << 31, 1 << 31
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for x := 0; x < 5; x++ {
        for y := 0; y < 5; y++ {
            if y == x { continue }
            cur, pre, mn, left := [5]int{}, [5]int{}, [2][2]int{{ inf, inf }, { inf, inf }}, 0
            for i, b := range s {
                cur[b - '0']++
                r := i + 1
                for r - left >= k && cur[x] > pre[x] && cur[y] > pre[y] {
                    mn[pre[x] & 1][pre[y] & 1] = min(mn[pre[x] & 1][pre[y] & 1], pre[x] - pre[y])
                    pre[s[left] - '0']++
                    left++
                }
                res = max(res, cur[x] - cur[y] - mn[cur[x] & 1 ^ 1][cur[y] & 1])
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "12233", k = 4
    // Output: -1
    // Explanation:
    // For the substring "12233", the frequency of '1' is 1 and the frequency of '3' is 2. The difference is 1 - 2 = -1.
    fmt.Println(maxDifference("12233", 4)) // -1
    // Example 2:
    // Input: s = "1122211", k = 3
    // Output: 1
    // Explanation:
    // For the substring "11222", the frequency of '2' is 3 and the frequency of '1' is 2. The difference is 3 - 2 = 1.
    fmt.Println(maxDifference("1122211", 3)) // 1
    // Example 3:
    // Input: s = "110", k = 3
    // Output: -1
    fmt.Println(maxDifference("110", 3)) // -1

    fmt.Println(maxDifference("0000000000", 3)) // -2147483638
    fmt.Println(maxDifference("1111111111", 3)) // -2147483638
    fmt.Println(maxDifference("1010101010", 3)) // 1
    fmt.Println(maxDifference("0101010101", 3)) // 1
    fmt.Println(maxDifference("1111100000", 3)) // 3
    fmt.Println(maxDifference("0000011111", 3)) // 3
    fmt.Println(maxDifference("0123401234", 3)) // -1
    fmt.Println(maxDifference("4321043210", 3)) // -1
    fmt.Println(maxDifference("0011223344", 3)) // -1
    fmt.Println(maxDifference("4433221100", 3)) // -1
}