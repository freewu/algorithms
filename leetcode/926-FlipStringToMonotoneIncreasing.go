package main

// 926. Flip String to Monotone Increasing
// A binary string is monotone increasing if it consists of some number of 0's (possibly none), followed by some number of 1's (also possibly none).
// You are given a binary string s. You can flip s[i] changing it from 0 to 1 or from 1 to 0.
// Return the minimum number of flips to make s monotone increasing.

// Example 1:
// Input: s = "00110"
// Output: 1
// Explanation: We flip the last digit to get 00111.

// Example 2:
// Input: s = "010110"
// Output: 2
// Explanation: We flip to get 011111, or alternatively 000111.

// Example 3:
// Input: s = "00011000"
// Output: 2
// Explanation: We flip to get 00000000.

// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is either '0' or '1'.

import "fmt"

func minFlipsMonoIncr(s string) int {
    n, i, count0, count1, res := len(s), 0, 0, 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i < n {
        if s[i] == '1' {
            count1++
        } else {
            count0++
        }
        i++
        if count1 < count0 {
            res += count1
            count1, count0 = 0, 0
        }
    }
    return res + min(count0, count1)
}

func minFlipsMonoIncr1(s string) int {
    cnt1, cnt2 := 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, c := range s {
        t1, t2 := 0, 0
        if c == '0' {
            t1 = cnt1
            t2 = min(cnt1, cnt2) + 1
        } else {
            t1 = cnt1 + 1
            t2 = min(cnt1, cnt2)
        }
        cnt1, cnt2 = t1, t2
    }
    return min(cnt1, cnt2)
}

func main() {
    // Example 1:
    // Input: s = "00110"
    // Output: 1
    // Explanation: We flip the last digit to get 00111.
    fmt.Println(minFlipsMonoIncr("00110")) // 1
    // Example 2:
    // Input: s = "010110"
    // Output: 2
    // Explanation: We flip to get 011111, or alternatively 000111.
    fmt.Println(minFlipsMonoIncr("010110")) // 2
    // Example 3:
    // Input: s = "00011000"
    // Output: 2
    // Explanation: We flip to get 00000000.
    fmt.Println(minFlipsMonoIncr("00011000")) // 2

    fmt.Println(minFlipsMonoIncr1("00110")) // 1
    fmt.Println(minFlipsMonoIncr1("010110")) // 2
    fmt.Println(minFlipsMonoIncr1("00011000")) // 2
}