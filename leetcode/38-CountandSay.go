package main

// 38. Count and Say
// The count-and-say sequence is a sequence of digit strings defined by the recursive formula:
//     1. countAndSay(1) = "1"
//     2. countAndSay(n) is the run-length encoding of countAndSay(n - 1).

// Run-length encoding (RLE) is a string compression method that works by replacing consecutive identical characters (repeated 2 or more times) with the concatenation of the character and the number marking the count of the characters (length of the run). For example, to compress the string "3322251" we replace "33" with "23", replace "222" with "32", replace "5" with "15" and replace "1" with "11". Thus the compressed string becomes "23321511".

// Given a positive integer n, return the nth element of the count-and-say sequence.

// Example 1:
// Input: n = 4
// Output: "1211"
// Explanation:
// countAndSay(1) = "1"
// countAndSay(2) = RLE of "1" = "11"
// countAndSay(3) = RLE of "11" = "21"
// countAndSay(4) = RLE of "21" = "1211"

// Example 2:
// Input: n = 1
// Output: "1"
// Explanation:
// This is the base case.

// Constraints:
//     1 <= n <= 30

// Follow up: Could you solve it iteratively?

import "fmt"
import "strconv"

func countAndSay(n int) string {
    if n <= 0 { return "" }
    res := "1"
    for { // 循环n次
        if (n - 1) == 0 { break }
        s, r, c, l := "", 0, res[0], len(res)
        // 统计字符串
        for i := 0; i < l; i++ {
            if c == res[i] {
                r++
                continue
            }
            s += strconv.Itoa(r) + string(c)
            c, r = res[i], 1
        }
        s += strconv.Itoa(r) + string(c)
        res = s
        n--
    }
    return res
}

func countAndSay1(n int) string {
    res := "1"
    for i := 0; i < n - 1; i++ {
        j, tmp := 0,""
        for k, v := range res {
            if v != rune(res[j]) {
                tmp += strconv.Itoa(k - j) + string(res[j])
                j = k
            }
        }
        res = tmp + strconv.Itoa(len(res) - j) + string(res[j])
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 4
    // Output: "1211"
    // Explanation:
    // countAndSay(1) = "1"
    // countAndSay(2) = RLE of "1" = "11"
    // countAndSay(3) = RLE of "11" = "21"
    // countAndSay(4) = RLE of "21" = "1211"
    fmt.Println(countAndSay(4)) // 1211
    // Example 2:
    // Input: n = 1
    // Output: "1"
    // Explanation:
    // This is the base case.
    fmt.Println(countAndSay(1)) // 1

    fmt.Println(countAndSay(2)) // 21
    fmt.Println(countAndSay(3)) // 21
    fmt.Println(countAndSay(5)) // 111221
    fmt.Println(countAndSay(6)) // 1312211
    fmt.Println(countAndSay(7)) // 13112221
    fmt.Println(countAndSay(8)) // 1113213211
    //fmt.Println(countAndSay(30)) // 

    fmt.Println(countAndSay1(4)) // 1211
    fmt.Println(countAndSay1(1)) // 1
    fmt.Println(countAndSay1(2)) // 21
    fmt.Println(countAndSay1(3)) // 21
    fmt.Println(countAndSay1(5)) // 111221
    fmt.Println(countAndSay1(6)) // 1312211
    fmt.Println(countAndSay1(7)) // 13112221
    fmt.Println(countAndSay1(8)) // 1113213211
    //fmt.Println(countAndSay(30)) // 1
}
