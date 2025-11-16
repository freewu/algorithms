package main

// 1513. Number of Substrings With Only 1s
// Given a binary string s, return the number of substrings with all characters 1's. 
// Since the answer may be too large, return it modulo 10^9 + 7.

// Example 1:
// Input: s = "0110111"
// Output: 9
// Explanation: There are 9 substring in total with only 1's characters.
// "1" -> 5 times.
// "11" -> 3 times.
// "111" -> 1 time.

// Example 2:
// Input: s = "101"
// Output: 2
// Explanation: Substring "1" is shown 2 times in s.

// Example 3:
// Input: s = "111111"
// Output: 21
// Explanation: Each substring contains only 1's characters.

// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is either '0' or '1'.

import "fmt"

func numSub(s string) int {
    res, n, mod := 0, 0, 1_000_000_007
    for _, v := range s {
        if v == '1' {
            n += 1
        } else {
            n = 0
        }
        res = (res + n) % mod
    }
    return res
}

func numSub1(s string) int {
    res, count, n := int64(0), 0, len(s)
    prefixSum := make([]int64,n + 1)
    for i := 1 ; i < n + 1; i++ {
        prefixSum[i] = prefixSum[i-1] + int64(i)
    }
    for _, v := range s {
        if v == '1' { // 1
            count += 1
        } else { // 0
            res = res + prefixSum[count]
            count = 0
        }
    }
    if count != 0 {
        res = res + prefixSum[count]
    }
    return int(res % 1_000_000_007)
}

func main() {
    // Example 1:
    // Input: s = "0110111"
    // Output: 9
    // Explanation: There are 9 substring in total with only 1's characters.
    // "1" -> 5 times.
    // "11" -> 3 times.
    // "111" -> 1 time.
    fmt.Println(numSub("0110111")) // 9
    // Example 2:
    // Input: s = "101"
    // Output: 2
    // Explanation: Substring "1" is shown 2 times in s.
    fmt.Println(numSub("101")) // 2
    // Example 3:
    // Input: s = "111111"
    // Output: 21
    // Explanation: Each substring contains only 1's characters.
    fmt.Println(numSub("111111")) // 21

    fmt.Println(numSub("1111111111")) // 55
    fmt.Println(numSub("0000000000")) // 0
    fmt.Println(numSub("1111100000")) // 15
    fmt.Println(numSub("0000011111")) // 15
    fmt.Println(numSub("0101010101")) // 5
    fmt.Println(numSub("1010101010")) // 5

    fmt.Println(numSub1("0110111")) // 9
    fmt.Println(numSub1("101")) // 2
    fmt.Println(numSub1("111111")) // 21
    fmt.Println(numSub1("1111111111")) // 55
    fmt.Println(numSub1("0000000000")) // 0
    fmt.Println(numSub1("1111100000")) // 15
    fmt.Println(numSub1("0000011111")) // 15
    fmt.Println(numSub1("0101010101")) // 5
    fmt.Println(numSub1("1010101010")) // 5
}