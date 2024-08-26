package main

// 1256. Encode Number
// Given a non-negative integer num, Return its encoding string.

// The encoding is done by converting the integer to a string using a secret function 
// that you should deduce from the following table:
// <img src="https://assets.leetcode.com/uploads/2019/06/21/encode_number.png" / >
// 0 ""
// 1 "0"
// 2 "1"
// 3 "00"
// 4 "01"
// 5 "10"
// 6 "11"
// 7 "000"

// Example 1:
// Input: num = 23
// Output: "1000"

// Example 2:
// Input: num = 107
// Output: "101100"

// Constraints:
//     0 <= num <= 10^9

import "fmt"

func encode(num int) string {
    return fmt.Sprintf("%b", num+1)[1:]
}

func encode1(num int) string {
    helper := func (x int) string {
        if x == 0 {
            return "0"
        }
        res := []byte{}
        for x > 0 {
            m := x % 2
            x /= 2
            res = append([]byte{byte( m + '0' )}, res...)
        }
        return string(res)
    }
    
    res := helper(num + 1)
    return res[1:]
}

func main() {
    // Example 1:
    // Input: num = 23
    // Output: "1000"
    fmt.Println(encode(23)) // 1000
    // Example 2:
    // Input: num = 107
    // Output: "101100"
    fmt.Println(encode(107)) // 101100

    fmt.Println(encode1(23)) // 1000
    fmt.Println(encode1(107)) // 101100
}