package main

// 3399. Smallest Substring With Identical Characters II
// You are given a binary string s of length n and an integer numOps.

// You are allowed to perform the following operation on s at most numOps times:
//     1. Select any index i (where 0 <= i < n) and flip s[i]. 
//        If s[i] == '1', change s[i] to '0' and vice versa.

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
//     1 <= n == s.length <= 10^5
//     s consists only of '0' and '1'.
//     0 <= numOps <= n

import "fmt"

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
    count, arr := 0, []int{}
    for i, j := 0, 0; i < len(s); i++ {
        if s[i] != s[j] {
            arr = append(arr, count)
            count, j = 0, i
        }
        count++
    }
    arr = append(arr, count)
    check0 := func(st int) bool {
        count := numOps
        for _, v := range s {
            if byte(v) != ('0') + byte(st) {
                count--
            }
            st = st ^ 1
        }
        return count >= 0
    }
    check := func(dist int) bool {
        if dist == 1 { // 1001 无解 10001 ok 100001
            return check0(0) || check0(1)
        }
        count := numOps
        for _, v := range arr {
            if v <= dist { // 不需要操作
                continue
            }
            for ; v > dist; v -= (dist + 1) {
                count--
            }
            if count < 0 {
                return false
            }
        }
        return count >= 0
    }
    l, r := 1, len(s)
    for l < r {
        mid := (l + r) / 2
        if check(mid) {
            r = mid
        } else {
            l = mid + 1
        }
    }
    return r
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