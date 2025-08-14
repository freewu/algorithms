package main

// 2264. Largest 3-Same-Digit Number in String
// You are given a string num representing a large integer. 
// An integer is good if it meets the following conditions:
//     It is a substring of num with length 3.
//     It consists of only one unique digit.

// Return the maximum good integer as a string or an empty string "" if no such integer exists.

// Note:
//     A substring is a contiguous sequence of characters within a string.
//     There may be leading zeroes in num or a good integer.

// Example 1:
// Input: num = "6777133339"
// Output: "777"
// Explanation: There are two distinct good integers: "777" and "333".
// "777" is the largest, so we return "777".

// Example 2:
// Input: num = "2300019"
// Output: "000"
// Explanation: "000" is the only good integer.

// Example 3:
// Input: num = "42352338"
// Output: ""
// Explanation: No substring of length 3 consists of only one unique digit. Therefore, there are no good integers.

// Constraints:
//     3 <= num.length <= 1000
//     num only consists of digits.

import "fmt"

func largestGoodInteger(num string) string {
    res := byte('0'-1)
    for i := 1; i < len(num) - 1; i++ {
        if num[i] == num[i + 1] && num[i] == num[i-1] && num[i] > res {
            res = num[i]
        }
    }
    if res == byte('0' - 1) {
        return ""
    }
    return string([]byte{ res, res, res })
}

func largestGoodInteger1(num string) string {
    res := ""
    max := func (x, y string) string { if x > y { return x; }; return y; }
    for i := 2; i < len(num); i++ {
        if num[i] == num[i-1] && num[i] == num[i-2] {
            res = max(res, num[i-2:i+1])
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: num = "6777133339"
    // Output: "777"
    // Explanation: There are two distinct good integers: "777" and "333".
    // "777" is the largest, so we return "777".
    fmt.Println(largestGoodInteger("6777133339")) // "777"
    // Example 2:
    // Input: num = "2300019"
    // Output: "000"
    // Explanation: "000" is the only good integer.
    fmt.Println(largestGoodInteger("2300019")) // 000
    // Example 3:
    // Input: num = "42352338"
    // Output: ""
    // Explanation: No substring of length 3 consists of only one unique digit. Therefore, there are no good integers.
    fmt.Println(largestGoodInteger("42352338")) // ""

    fmt.Println(largestGoodInteger("123456789")) // ""
    fmt.Println(largestGoodInteger("987654321")) // ""

    fmt.Println(largestGoodInteger1("6777133339")) // "777"
    fmt.Println(largestGoodInteger1("2300019")) // 000
    fmt.Println(largestGoodInteger1("42352338")) // ""
    fmt.Println(largestGoodInteger1("123456789")) // ""
    fmt.Println(largestGoodInteger1("987654321")) // ""
}