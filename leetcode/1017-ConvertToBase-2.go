package main

// 1017. Convert to Base -2
// Given an integer n, return a binary string representing its representation in base -2.
// Note that the returned string should not have leading zeros unless the string is "0".

// Example 1:
// Input: n = 2
// Output: "110"
// Explantion: (-2)^2 + (-2)^1 = 2

// Example 2:
// Input: n = 3
// Output: "111"
// Explantion: (-2)^2 + (-2)^1 + (-2)^0 = 3

// Example 3:
// Input: n = 4
// Output: "100"
// Explantion: (-2)^2 = 4

// Constraints:
//     0 <= n <= 10^9

import "fmt"
import "strconv"
import "strings"

func baseNeg2(n int) string {
    if n == 0 {
        return "0"
    }
    binary, i, dozen := []rune{}, 0, 0
    for n > 0 {
        res := (n + dozen) % 2
        if dozen == 1 && n % 2 == 0 {
            dozen = 0
        }
        if dozen == 0 && i % 2 == 1 && res == 1 {
            dozen = 1
        }
        n /= 2
        binary = append([]rune{rune(res + '0')}, binary...)
        i++
    }
    if dozen == 1 {
        binary = append([]rune{rune(dozen + '0')}, binary...)
        if i % 2 == 1 {
            binary = append([]rune{rune(dozen + '0')}, binary...)
        }
    }
    return string(binary)
}

func baseNeg21(n int) string {
    var sb strings.Builder
    if n == 0{
        return "0"
    }
    for n != 0{
        rem := n&1 
        sb.WriteString(strconv.Itoa(rem))
        n = -(n >> 1)
    }
    reverseString := func (s []byte) string{
        n := len(s)
        for i := 0;i< n/2; i++{
            temp := s[n-i-1]
            s[n-i-1] = s[i]
            s[i] = temp
        }
        return string(s)
    }
    return reverseString([]byte(sb.String()))
}

func main() {
    fmt.Println(baseNeg2(0)) // 0
    fmt.Println(baseNeg2(1)) // 1
    // Example 1:
    // Input: n = 2
    // Output: "110"
    // Explantion: (-2)^2 + (-2)^1 = 2
    fmt.Println(baseNeg2(2)) // 110
    // Example 2:
    // Input: n = 3
    // Output: "111"
    // Explantion: (-2)^2 + (-2)^1 + (-2)^0 = 3
    fmt.Println(baseNeg2(3)) // 111
    // Example 3:
    // Input: n = 4
    // Output: "100"
    // Explantion: (-2)^2 = 4
    fmt.Println(baseNeg2(4)) // 100

    fmt.Println(baseNeg21(0)) // 0
    fmt.Println(baseNeg21(1)) // 1
    fmt.Println(baseNeg21(2)) // 110
    fmt.Println(baseNeg21(3)) // 111
    fmt.Println(baseNeg21(4)) // 100
}