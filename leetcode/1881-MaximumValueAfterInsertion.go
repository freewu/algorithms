package main

// 1881. Maximum Value after Insertion
// You are given a very large integer n, represented as a string,​​​​​​ and an integer digit x. 
// The digits in n and the digit x are in the inclusive range [1, 9], and n may represent a negative number.

// You want to maximize n's numerical value by inserting x anywhere in the decimal representation of n​​​​​​. 
// You cannot insert x to the left of the negative sign.
//     1. For example, if n = 73 and x = 6, it would be best to insert it between 7 and 3, making n = 763.
//     2. If n = -55 and x = 2, it would be best to insert it before the first 5, making n = -255.

// Return a string representing the maximum value of n​​​​​​ after the insertion.

// Example 1:
// Input: n = "99", x = 9
// Output: "999"
// Explanation: The result is the same regardless of where you insert 9.

// Example 2:
// Input: n = "-13", x = 2
// Output: "-123"
// Explanation: You can make n one of {-213, -123, -132}, and the largest of those three is -123.

// Constraints:
//     1 <= n.length <= 10^5
//     1 <= x <= 9
//     The digits in n​​​ are in the range [1, 9].
//     n is a valid representation of an integer.
//     In the case of a negative n,​​​​​​ it will begin with '-'.

import "fmt"

func maxValue(n string, x int) string {
    i, xb := 0, '0' + byte(x)
    for i < len(n) && (n[0] == '-' || xb <= n[i]) && (n[0] != '-' || xb >= n[i]) || (i == 0 && n[i] == '-') {
        i++
    }
    return n[:i] + string(xb) + n[i:]
}

func maxValue1(n string, x int) string {
    i, neg, xb := 0, false, '0' + byte(x)
    if n[0] == '-' { // 判断是否是负数
        neg = true
        i = 1
    }
    for i < len(n) {
        if (!neg && xb > n[i]) || (neg && xb < n[i]) { break }
        i++
    }
    return n[:i] + string(xb) + n[i:]
}

func main() {
    // Example 1:
    // Input: n = "99", x = 9
    // Output: "999"
    // Explanation: The result is the same regardless of where you insert 9.
    fmt.Println(maxValue("99", 9)) // "999"
    // Example 2:
    // Input: n = "-13", x = 2
    // Output: "-123"
    // Explanation: You can make n one of {-213, -123, -132}, and the largest of those three is -123.
    fmt.Println(maxValue("-13", 2)) // "-123"

    fmt.Println(maxValue1("99", 9)) // "999"
    fmt.Println(maxValue1("-13", 2)) // "-123"
}