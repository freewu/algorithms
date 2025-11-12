package main

// 3228. Maximum Number of Operations to Move Ones to the End
// You are given a binary string s.

// You can perform the following operation on the string any number of times:
//     1. Choose any index i from the string where i + 1 < s.length such that s[i] == '1' and s[i + 1] == '0'.
//     2. Move the character s[i] to the right until it reaches the end of the string or another '1'. 
//        For example, for s = "010010", if we choose i = 1, the resulting string will be s = "000110".

// Return the maximum number of operations that you can perform.

// Example 1:
// Input: s = "1001101"
// Output: 4
// Explanation:
// We can perform the following operations:
// Choose index i = 0. The resulting string is s = "0011101".
// Choose index i = 4. The resulting string is s = "0011011".
// Choose index i = 3. The resulting string is s = "0010111".
// Choose index i = 2. The resulting string is s = "0001111".

// Example 2:
// Input: s = "00111"
// Output: 0

// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is either '0' or '1'.

import "fmt"

func maxOperations(s string) int {
    res, cur, n, start, end := 0, 1, len(s), -1, -1
    for i := n - 1; i >= 0; i-- {
        if s[i] == '0' {
            start = i
            break
        }
    }
    for i := start; i >= 0; i-- {
        if s[i] == '0' {
            continue
        } else {
            for end = i; end >= 0; end-- {
                if s[end] == '0' {
                    break
                }
            }
            res += cur * (i - end)
            cur++
            i = end + 1
        }
    }
    return res
}

func maxOperations1(s string) int {
    res, count1 := 0, 0
    for i, v := range s {
        if v == '1' {
            count1++
        } else if i > 0 && s[i - 1] == '1'{
            res += count1
        }
    }
    return res
}

func maxOperations2(s string) int {
    res, carried, last, n := 0, 0, 0, len(s)
    for i := 0; i < n; {
        if s[i] == '1' {
            res += carried
            for i < n && s[i] == '1' {
                last = i
                carried++
                i++
            }
            continue
        }
        i++
    }
    if last < n - 1 {
        res += (carried)
    }
    return res
}

func maxOperations3(s string) int {
    res, one, i := 0, 0, 0
    for i < len(s) {
        if s[i] == '0' {
            for i + 1 < len(s) && s[i + 1] == '0' {
                i++
            }
            res += one
        } else {
            one++
        }
        i++
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "1001101"
    // Output: 4
    // Explanation:
    // We can perform the following operations:
    // Choose index i = 0. The resulting string is s = "0011101".
    // Choose index i = 4. The resulting string is s = "0011011".
    // Choose index i = 3. The resulting string is s = "0010111".
    // Choose index i = 2. The resulting string is s = "0001111".
    fmt.Println(maxOperations("1001101")) // 4
    // Example 2:
    // Input: s = "00111"
    // Output: 0
    fmt.Println(maxOperations("00111")) // 0

    fmt.Println(maxOperations("00000000")) // 0
    fmt.Println(maxOperations("11111111")) // 0
    fmt.Println(maxOperations("10101010")) // 10
    fmt.Println(maxOperations("01010101")) // 6

    fmt.Println(maxOperations1("1001101")) // 4
    fmt.Println(maxOperations1("00111")) // 0
    fmt.Println(maxOperations1("00000000")) // 0
    fmt.Println(maxOperations1("11111111")) // 0
    fmt.Println(maxOperations1("10101010")) // 10
    fmt.Println(maxOperations1("01010101")) // 6

    fmt.Println(maxOperations2("1001101")) // 4
    fmt.Println(maxOperations2("00111")) // 0
    fmt.Println(maxOperations2("00000000")) // 0
    fmt.Println(maxOperations2("11111111")) // 0
    fmt.Println(maxOperations2("10101010")) // 10
    fmt.Println(maxOperations2("01010101")) // 6

    fmt.Println(maxOperations3("1001101")) // 4
    fmt.Println(maxOperations3("00111")) // 0
    fmt.Println(maxOperations3("00000000")) // 0
    fmt.Println(maxOperations3("11111111")) // 0
    fmt.Println(maxOperations3("10101010")) // 10
    fmt.Println(maxOperations3("01010101")) // 6
}