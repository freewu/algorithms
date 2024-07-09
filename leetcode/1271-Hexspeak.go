package main

// 1271. Hexspeak
// A decimal number can be converted to its Hexspeak representation by first converting it to an uppercase hexadecimal string, 
// then replacing all occurrences of the digit '0' with the letter 'O', and the digit '1' with the letter 'I'. 
// Such a representation is valid if and only if it consists only of the letters in the set {'A', 'B', 'C', 'D', 'E', 'F', 'I', 'O'}.

// Given a string num representing a decimal integer n, 
// return the Hexspeak representation of n if it is valid, otherwise return "ERROR".

// Example 1:

// Input: num = "257"
// Output: "IOI"
// Explanation: 257 is 101 in hexadecimal.

// Example 2:
// Input: num = "3"
// Output: "ERROR"

// Constraints:
//     1 <= num.length <= 12
//     num does not contain leading zeros.
//     num represents an integer in the range [1, 10^12].

import "fmt"
import "strconv"

func toHexspeak(num string) string {
    i, _ := strconv.Atoi(num)
    hex := []byte(fmt.Sprintf("%x", i)) // 转成16进制数组
    maps := map[byte]byte{ '0':'O', '1':'I', 'a':'A', 'b':'B', 'c':'C', 'd':'D', 'e':'E', 'f':'F', }
    res := ""
    for _, v := range hex {
        if _, ok := maps[v]; !ok { // 出现了 2 - 9 的数字
            return  "ERROR"
        }
        res += string(maps[v])
    }
    return res
}

func main() {
    // Input: num = "257"
    // Output: "IOI"
    // Explanation: 257 is 101 in hexadecimal.
    fmt.Println(toHexspeak("257")) // "IOI"
    // Example 2:
    // Input: num = "3"
    // Output: "ERROR"
    fmt.Println(toHexspeak("3")) // "ERROR"

    fmt.Println(toHexspeak("11")) // "B"
}