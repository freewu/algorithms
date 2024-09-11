package main

// 3167. Better Compression of String
// You are given a string compressed representing a compressed version of a string. 
// The format is a character followed by its frequency. 
// For example, "a3b1a1c2" is a compressed version of the string "aaabacc".

// We seek a better compression with the following conditions:
//     1. Each character should appear only once in the compressed version.
//     2. The characters should be in alphabetical order.

// Return the better compression of compressed.

// Note: In the better version of compression, the order of letters may change, which is acceptable.

// Example 1:
// Input: compressed = "a3c9b2c1"
// Output: "a3b2c10"
// Explanation:
// Characters "a" and "b" appear only once in the input, but "c" appears twice, once with a size of 9 and once with a size of 1.
// Hence, in the resulting string, it should have a size of 10.

// Example 2:
// Input: compressed = "c2b3a1"
// Output: "a1b3c2"

// Example 3:
// Input: compressed = "a2b4c1"
// Output: "a2b4c1"

// Constraints:
//     1 <= compressed.length <= 6 * 10^4
//     compressed consists only of lowercase English letters and digits.
//     compressed is a valid compression, i.e., each character is followed by its frequency.
//     Frequencies are in the range [1, 10^4] and have no leading zeroes.

import "fmt"
import "strconv"

func betterCompression(compressed string) string {
    mp, k, num, n := [26]int{}, byte('a'), 0, len(compressed)
    isLetter := func(c byte) bool { return c >= 'a' && c <= 'z' }
    isNumber := func(c byte) bool { return c >= '0' || c <= '9' }
    for i := range compressed {
        c := compressed[i]
        //fmt.Printf("%c,%v ,%v\n", c, isLetter(c), isNumber(c))
        if isLetter(c) { // 字符
            k = c
        } else if isNumber(c) { // 数字
            num = num * 10 + int(c - '0')
            if i == n - 1 || isLetter(compressed[i + 1]) {
                mp[int( k - 'a')] += num
                num = 0
            }
        }
    }
    res := ""
    for k, v := range mp {
        if v > 0 {
            res += string(k + 'a') + strconv.Itoa(v)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: compressed = "a3c9b2c1"
    // Output: "a3b2c10"
    // Explanation:
    // Characters "a" and "b" appear only once in the input, but "c" appears twice, once with a size of 9 and once with a size of 1.
    // Hence, in the resulting string, it should have a size of 10.
    fmt.Println(betterCompression("a3c9b2c1")) // "a3b2c10"
    // Example 2:
    // Input: compressed = "c2b3a1"
    // Output: "a1b3c2"
    fmt.Println(betterCompression("c2b3a1")) // "a1b3c2"
    // Example 3:
    // Input: compressed = "a2b4c1"
    // Output: "a2b4c1"
    fmt.Println(betterCompression("a2b4c1")) // "a2b4c1"
}