package main

// 1573. Number of Ways to Split a String
// Given a binary string s, you can split s into 3 non-empty strings s1, s2, and s3 where s1 + s2 + s3 = s.

// Return the number of ways s can be split such that the number of ones is the same in s1, s2, and s3. 
// Since the answer may be too large, return it modulo 10^9 + 7.

// Example 1:
// Input: s = "10101"
// Output: 4
// Explanation: There are four ways to split s in 3 parts where each part contain the same number of letters '1'.
// "1|010|1"
// "1|01|01"
// "10|10|1"
// "10|1|01"

// Example 2:
// Input: s = "1001"
// Output: 0

// Example 3:
// Input: s = "0000"
// Output: 3
// Explanation: There are three ways to split s in 3 parts.
// "0|0|00"
// "0|00|0"
// "00|0|0"

// Constraints:
//     3 <= s.length <= 10^5
//     s[i] is either '0' or '1'.

import "fmt"

func numWays(s string) int {
    numOnes, mod := 0, 1_000_000_007
    for i := range s {
        if s[i] == '1' { // 统计 1 出现的次数
            numOnes++
        }
    }
    if numOnes % 3 != 0 { return 0 } // 不能被 3 整除
    calc := func(n int) int { return (n-1) * (n-2) / 2; }
    if numOnes == 0 {
        return calc(len(s)) % mod
    }
    numOnes /= 3
    i := 0
    advanceToSplitAndCountZeros := func() int {
        for j := numOnes; j > 0; i++ { // advance until we have enough "1" for this string
            if s[i] == '1' {
                j--
            }
        }
        // number of ways to split is based on the number of zeros
        // between previous string's last "1" and the next string's first "1".
        res := 0
        for ; s[i] == '0'; i++ {
            res++
        }
        return res
    }
    firstSplit, secondSplit := advanceToSplitAndCountZeros() + 1, advanceToSplitAndCountZeros() + 1
    return (firstSplit * secondSplit) % mod
}

func main() {
    // Example 1:
    // Input: s = "10101"
    // Output: 4
    // Explanation: There are four ways to split s in 3 parts where each part contain the same number of letters '1'.
    // "1|010|1"
    // "1|01|01"
    // "10|10|1"
    // "10|1|01"
    fmt.Println(numWays("10101")) // 4
    // Example 2:
    // Input: s = "1001"
    // Output: 0
    fmt.Println(numWays("1001")) // 0
    // Example 3:
    // Input: s = "0000"
    // Output: 3
    // Explanation: There are three ways to split s in 3 parts.
    // "0|0|00"
    // "0|00|0"
    // "00|0|0"
    fmt.Println(numWays("0000")) // 3
}