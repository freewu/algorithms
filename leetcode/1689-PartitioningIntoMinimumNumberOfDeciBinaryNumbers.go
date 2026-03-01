package main

// 1689. Partitioning Into Minimum Number Of Deci-Binary Numbers
// A decimal number is called deci-binary if each of its digits is either 0 or 1 without any leading zeros.
// For example, 101 and 1100 are deci-binary, while 112 and 3001 are not.

// Given a string n that represents a positive decimal integer, 
// return the minimum number of positive deci-binary numbers needed so that they sum up to n.

// Example 1:
// Input: n = "32"
// Output: 3
// Explanation: 10 + 11 + 11 = 32

// Example 2:
// Input: n = "82734"
// Output: 8

// Example 3:
// Input: n = "27346209830709182346"
// Output: 9

// Constraints:
//     1 <= n.length <= 10^5
//     n consists of only digits.
//     n does not contain any leading zeros and represents a positive integer.

import "fmt"
import "slices"

func minPartitions(n string) int {
    res := byte(0)
    for i := 0; i < len(n); i++ {
        if n[i] == '9' {
            return 9
        }
        if n[i] > res {
            res = n[i]
        }
    }
    return int(res - '0')
}

func minPartitions1(n string) int {
    // 贪心,脑筋急转弯
    // 拆分最大的数字即可,其余位可以顺带着0或者1
    return int(slices.Max([]byte(n)) - '0')
}

func main() {
    // Example 1:
    // Input: n = "32"
    // Output: 3
    // Explanation: 10 + 11 + 11 = 32
    fmt.Println(minPartitions("32")) // 3
    // Example 2:
    // Input: n = "82734"
    // Output: 8
    fmt.Println(minPartitions("82734")) // 8
    // Example 3:
    // Input: n = "27346209830709182346"
    // Output: 9
    fmt.Println(minPartitions("82734")) // 8

    fmt.Println(minPartitions("1")) // 1
    fmt.Println(minPartitions("1024")) // 4
    fmt.Println(minPartitions("10000")) // 1
    fmt.Println(minPartitions("99999")) // 9
    fmt.Println(minPartitions("1000000000")) // 1
    fmt.Println(minPartitions("1000000007")) // 7

    fmt.Println(minPartitions1("32")) // 3
    fmt.Println(minPartitions1("82734")) // 8
    fmt.Println(minPartitions1("82734")) // 8
    fmt.Println(minPartitions1("1")) // 1
    fmt.Println(minPartitions1("1024")) // 4
    fmt.Println(minPartitions1("10000")) // 1
    fmt.Println(minPartitions1("99999")) // 9
    fmt.Println(minPartitions1("1000000000")) // 1
    fmt.Println(minPartitions1("1000000007")) // 7
}