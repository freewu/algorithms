package main

// 3602. Hexadecimal and Hexatrigesimal Conversion
// You are given an integer n.

// Return the concatenation of the hexadecimal representation of n2 and the hexatrigesimal representation of n3.

// A hexadecimal number is defined as a base-16 numeral system that uses the digits 0 – 9 and the uppercase letters A - F to represent values from 0 to 15.

// A hexatrigesimal number is defined as a base-36 numeral system that uses the digits 0 – 9 and the uppercase letters A - Z to represent values from 0 to 35.

// Example 1:
// Input: n = 13
// Output: "A91P1"
// Explanation:
// n2 = 13 * 13 = 169. In hexadecimal, it converts to (10 * 16) + 9 = 169, which corresponds to "A9".
// n3 = 13 * 13 * 13 = 2197. In hexatrigesimal, it converts to (1 * 362) + (25 * 36) + 1 = 2197, which corresponds to "1P1".
// Concatenating both results gives "A9" + "1P1" = "A91P1".

// Example 2:
// Input: n = 36
// Output: "5101000"
// Explanation:
// n2 = 36 * 36 = 1296. In hexadecimal, it converts to (5 * 162) + (1 * 16) + 0 = 1296, which corresponds to "510".
// n3 = 36 * 36 * 36 = 46656. In hexatrigesimal, it converts to (1 * 363) + (0 * 362) + (0 * 36) + 0 = 46656, which corresponds to "1000".
// Concatenating both results gives "510" + "1000" = "5101000".

// Constraints:
//     1 <= n <= 1000

import "fmt"
import "strings"

func concatHex36(n int) string {
    n2 := n * n
    n3 := n * n * n
    Base36 := func(n int) string {
        if n == 0 { return "0" }
        digits := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        var sb strings.Builder
        for n > 0 {
            sb.WriteByte(digits[n%36])
            n /= 36
        }
        res := []rune(sb.String())
        for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
            res[i], res[j] = res[j], res[i]
        }
        return string(res)
    }
    hex := strings.ToUpper(fmt.Sprintf("%X", n2))
    base36 := strings.ToUpper(fmt.Sprintf("%s", Base36(n3)))
    return hex + base36
}

func main() {
    // Example 1:
    // Input: n = 13
    // Output: "A91P1"
    // Explanation:
    // n2 = 13 * 13 = 169. In hexadecimal, it converts to (10 * 16) + 9 = 169, which corresponds to "A9".
    // n3 = 13 * 13 * 13 = 2197. In hexatrigesimal, it converts to (1 * 362) + (25 * 36) + 1 = 2197, which corresponds to "1P1".
    // Concatenating both results gives "A9" + "1P1" = "A91P1".
    fmt.Println(concatHex36(13)) // A91P1
    // Example 2:
    // Input: n = 36
    // Output: "5101000"
    // Explanation:
    // n2 = 36 * 36 = 1296. In hexadecimal, it converts to (5 * 162) + (1 * 16) + 0 = 1296, which corresponds to "510".
    // n3 = 36 * 36 * 36 = 46656. In hexatrigesimal, it converts to (1 * 363) + (0 * 362) + (0 * 36) + 0 = 46656, which corresponds to "1000".
    // Concatenating both results gives "510" + "1000" = "5101000".
    fmt.Println(concatHex36(36)) // 5101000

    fmt.Println(concatHex36(1)) // 11
    fmt.Println(concatHex36(999)) // F3A71GHL8FR
    fmt.Println(concatHex36(1000)) // F4240GJDGXS
}