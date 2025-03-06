package main

// 3398. Smallest Substring With Identical Characters I
// You are given a binary string s of length n and an integer numOps.

// You are allowed to perform the following operation on s at most numOps times:
//     Select any index i (where 0 <= i < n) and flip s[i]. If s[i] == '1', change s[i] to '0' and vice versa.

// You need to minimize the length of the longest substring of s such that all the characters in the substring are identical.

// Return the minimum length after the operations.

// Example 1:
// Input: s = "000001", numOps = 1
// Output: 2
// Explanation: 
// By changing s[2] to '1', s becomes "001001". The longest substrings with identical characters are s[0..1] and s[3..4].

// Example 2:
// Input: s = "0000", numOps = 2
// Output: 1
// Explanation: 
// By changing s[0] and s[2] to '1', s becomes "1010".

// Example 3:
// Input: s = "0101", numOps = 0
// Output: 1

// Constraints:
//     1 <= n == s.length <= 1000
//     s consists only of '0' and '1'.
//     0 <= numOps <= n

import "fmt"
import "sort"

func minLength(s string, numOps int) int {
    n := len(s)
    if n == 1 { return 1 }
    helper := func(s string, mx int) int {
        n := len(s)
        if mx == 1 {
            a, b := 0, 0
            for i := 0; i < n; i++ {
                if s[i] == '0' {
                    a, b = b, a + 1
                } else {
                    a, b = b + 1, a
                }
            }
            if b < a {
                return b
            }
            return a
        }
        p, c, res := s[0], 0, 0
        for i := 0; i < n; i++ {
            if s[i] == p {
                c++
                continue
            }
            if c > mx {
                res += c / (mx + 1)
            }
            p, c = s[i], 1
        }
        if c > mx {
            res += c / (mx + 1)
        }
        return res
    }
    l, r := 0, n
    for r - l > 1 {
        m := (l + r) / 2
        if helper(s, m) > numOps {
            l = m
        } else {
            r = m
        }
    }
    return r
}

func minLength1(s string, numOps int) int {
    n := len(s)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    return 1 + sort.Search(n - 1, func(m int) bool {
        m++
        count := 0
        if m == 1 {
            // 改成 0101...
            for i, b := range s {
                // 如果 s[i] 和 i 的奇偶性不同，count 加一
                count += (int(b) ^ i) & 1
            }
            // n - count 表示改成 1010...
            count = min(count, n - count)
        } else {
            k := 0
            for i := 0; i < n; i++ {
                k++
                // 到达连续相同子串的末尾
                if i == n-1 || s[i] != s[i+1] {
                    count += k / (m + 1)
                    k = 0
                }
            }
        }
        return count <= numOps
    })
}

func main() {
    // Example 1:
    // Input: s = "000001", numOps = 1
    // Output: 2
    // Explanation: 
    // By changing s[2] to '1', s becomes "001001". The longest substrings with identical characters are s[0..1] and s[3..4].
    fmt.Println(minLength("000001", 1)) // 2
    // Example 2:
    // Input: s = "0000", numOps = 2
    // Output: 1
    // Explanation: 
    // By changing s[0] and s[2] to '1', s becomes "1010".
    fmt.Println(minLength("0000", 2)) // 1
    // Example 3:
    // Input: s = "0101", numOps = 0
    // Output: 1
    fmt.Println(minLength("0101", 0)) // 1

    fmt.Println(minLength1("000001", 1)) // 2
    fmt.Println(minLength1("0000", 2)) // 1
    fmt.Println(minLength1("0101", 0)) // 1
}