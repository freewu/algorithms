package main

// 3663. Find The Least Frequent Digit
// Given an integer n, find the digit that occurs least frequently in its decimal representation. 
// If multiple digits have the same frequency, choose the smallest digit.

// Return the chosen digit as an integer.

// The frequency of a digit x is the number of times it appears in the decimal representation of n.

// Example 1:
// Input: n = 1553322
// Output: 1
// Explanation:
// The least frequent digit in n is 1, which appears only once. All other digits appear twice.

// Example 2:
// Input: n = 723344511
// Output: 2
// Explanation:
// The least frequent digits in n are 7, 2, and 5; each appears only once.

// Constraints:
//     1 <= n <= 2^31​​​​​​​ - 1

import "fmt"
import "strconv"

func getLeastFrequentDigit(n int) int {
    // 把数字转成字符串
    s := strconv.Itoa(n)
    // 统计每个数字出现的次数
    freq := make(map[int]int)
    for _, v := range s {
        freq[int(v-'0')]++
    }
    // 找到出现次数最少的数字
    mn := 1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, v := range freq {
        mn = min(mn, v)
    }
    // 找到出现次数最少的数字中最小的那个
    res := 10
    for k, v := range freq {
        if v == mn {
            res = min(res, k)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 1553322
    // Output: 1
    // Explanation:
    // The least frequent digit in n is 1, which appears only once. All other digits appear twice.
    fmt.Println(getLeastFrequentDigit(1553322)) // 1
    // Example 2:
    // Input: n = 723344511
    // Output: 2
    // Explanation:
    // The least frequent digits in n are 7, 2, and 5; each appears only once.
    fmt.Println(getLeastFrequentDigit(723344511)) // 2

    fmt.Println(getLeastFrequentDigit(1)) // 1
    fmt.Println(getLeastFrequentDigit(1024)) // 0
    fmt.Println(getLeastFrequentDigit(9999)) // 9
    fmt.Println(getLeastFrequentDigit(10000)) // 1
    fmt.Println(getLeastFrequentDigit(2 << 31)) // 7
}