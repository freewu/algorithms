package main

// 1702. Maximum Binary String After Change
// You are given a binary string binary consisting of only 0's or 1's. 
// You can apply each of the following operations any number of times:
//     Operation 1: If the number contains the substring "00", you can replace it with "10".
//         For example, "00010" -> "10010"
//     Operation 2: If the number contains the substring "10", you can replace it with "01".
//         For example, "00010" -> "00001"

// Return the maximum binary string you can obtain after any number of operations. 
// Binary string x is greater than binary string y if x's decimal representation is greater than y's decimal representation.

// Example 1:
// Input: binary = "000110"
// Output: "111011"
// Explanation: A valid transformation sequence can be:
// "000110" -> "000101" 
// "000101" -> "100101" 
// "100101" -> "110101" 
// "110101" -> "110011" 
// "110011" -> "111011"

// Example 2:
// Input: binary = "01"
// Output: "01"
// Explanation: "01" cannot be transformed any further.

// Constraints:
//     1 <= binary.length <= 10^5
//     binary consist of '0' and '1'.

import "fmt"
import "strings"

func maximumBinaryString(binary string) string {
    res, index := []byte(binary), -1
    for i := 0; i < len(res); i++ {
        if res[i] == '0' { 
            // Example (011110 -> 101111); indexOf [ previousZeroIndex = 0; previousZeroIndex+1 = 1; i = 5]
            // We need to prefer operation 2 before operation 1 as it results in bigger number.
            if index != -1 { // Case of 10 -> 01
                res[index] = '1'
                res[index + 1] = '0'
                res[i] = '1'
                index++
            } else if i+1 < len(res) && res[i + 1] == '0' {  // 00 -> 10
                res[i] = '1'
            } else { // Else we can update the position of the last zero
                index = i
            }
        }
    }
    return string(res)
}

func maximumBinaryString1(binary string) string {
    // 统计 0 的数量
    zeros := 0
    for i := range binary {
        if binary[i] == '0' {
            zeros++
        }
    }
    if zeros <= 1 { // 不超过两个 0 直接返回 
        return binary
    }
    d := 0
    for ; d < len(binary) && binary[d] == '1'; d++ {}
    //fmt.Printf("d = %v\n",d)
    n := len(binary)
    return binary[:d] + strings.Repeat("1", zeros - 1) + "0" + strings.Repeat("1", n - d - zeros)
}

func main() {
    // Explanation: A valid transformation sequence can be:
    // "000110" -> "000101" 
    // "000101" -> "100101" 
    // "100101" -> "110101" 
    // "110101" -> "110011" 
    // "110011" -> "111011"
    fmt.Println(maximumBinaryString("000110")) // 111011
    // Explanation: "01" cannot be transformed any further.
    fmt.Println(maximumBinaryString("01")) // 01
    fmt.Println(maximumBinaryString("0101010101")) // 1111011111

    fmt.Println(maximumBinaryString("1111111100")) // 1111111110

    fmt.Println(maximumBinaryString1("000110")) // 111011
    fmt.Println(maximumBinaryString1("01")) // 01
    fmt.Println(maximumBinaryString1("0101010101")) // 1111011111
    fmt.Println(maximumBinaryString1("1111111100")) // 1111111110
    
}