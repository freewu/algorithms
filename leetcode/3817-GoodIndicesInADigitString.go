package main

// 3817. Good Indices in a Digit String
// You are given a string s consisting of digits.

// An index i is called good if there exists a substring of s that ends at index i and is equal to the decimal representation of i.

// Return an integer array of all good indices in increasing order.

// Example 1:
// Input: s = "0234567890112"
// Output: [0,11,12]
// Explanation:​​​​​​​
// At index 0, the decimal representation of the index is "0". The substring s[0] is "0", which matches, so index 0 is good.
// At index 11, the decimal representation is "11". The substring s[10..11] is "11", which matches, so index 11 is good.
// At index 12, the decimal representation is "12". The substring s[11..12] is "12", which matches, so index 12 is good.
// No other index has a substring ending at it that equals its decimal representation. Therefore, the answer is [0, 11, 12].

// Example 2:
// Input: s = "01234"
// Output: [0,1,2,3,4]
// Explanation:
// For every index i from 0 to 4, the decimal representation of i is a single digit, and the substring s[i] matches that digit.
// Therefore, a valid substring ending at each index exists, making all indices good.

// Example 3:
// Input: s = "12345"
// Output: []
// Explanation:
// No index has a substring ending at it that matches its decimal representation.
// Therefore, there are no good indices and the result is an empty array.

// Constraints:
//     1 <= s.length <= 10^5
//     s only consists of digits from '0' to '9'.

import "fmt"
import "strconv"

func goodIndices(s string) []int {
    res, n := []int{}, len(s)
    for i := 0; i < n; i++ {
        // 将索引 i 转为字符串，获取其长度（即需要匹配的子串长度）
        target := strconv.Itoa(i)
        m := len(target)
        // 计算子串的起始位置：i - m + 1
        start := i - m + 1
        // 起始位置不能小于0，否则无法截取有效子串
        if start < 0 {
            continue
        }
        // 截取以i结尾、长度为 m 的子串
        subStr := s[start : i + 1] // Go切片是左闭右开，所以结束位置是i+1
        // 比较子串是否等于目标字符串
        if subStr == target {
            res = append(res, i)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "0234567890112"
    // Output: [0,11,12]
    // Explanation:​​​​​​​
    // At index 0, the decimal representation of the index is "0". The substring s[0] is "0", which matches, so index 0 is good.
    // At index 11, the decimal representation is "11". The substring s[10..11] is "11", which matches, so index 11 is good.
    // At index 12, the decimal representation is "12". The substring s[11..12] is "12", which matches, so index 12 is good.
    // No other index has a substring ending at it that equals its decimal representation. Therefore, the answer is [0, 11, 12].
    fmt.Println(goodIndices("0234567890112")) // [0,11,12]
    // Example 2:
    // Input: s = "01234"
    // Output: [0,1,2,3,4]
    // Explanation:
    // For every index i from 0 to 4, the decimal representation of i is a single digit, and the substring s[i] matches that digit.
    // Therefore, a valid substring ending at each index exists, making all indices good.
    fmt.Println(goodIndices("01234")) // [0,1,2,3,4]
    // Example 3:
    // Input: s = "12345"
    // Output: []
    // Explanation:
    // No index has a substring ending at it that matches its decimal representation.
    // Therefore, there are no good indices and the result is an empty array.
    fmt.Println(goodIndices("12345")) // []

    fmt.Println(goodIndices("0123456789")) // [0 1 2 3 4 5 6 7 8 9]
    fmt.Println(goodIndices("9876543210")) // []
    fmt.Println(goodIndices("123456789")) // []
    fmt.Println(goodIndices("987654321")) // []
}