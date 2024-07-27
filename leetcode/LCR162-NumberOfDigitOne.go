package main

// LCR 162. 数字 1 的个数
// 给定一个整数 num，计算所有小于等于 num 的非负整数中数字 1 出现的个数。

// 示例 1：
// 输入：num = 0
// 输出：0

// 示例 2：
// 输入：num = 13
// 输出：6

// 提示：
//     0 <= num < 10^9

import "fmt"

func digitOneInNumber(n int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    count, base := 0, 1
    for base <= n {
        divider := base * 10
        count += (n / divider) * base + min(max(n % divider - base + 1, 0), base)
        base *= 10
    }
    return count
}

func main() {
    // Example 1:
    // Input: n = 13
    // Output: 6
    // 1 10 11 12 13
    fmt.Println(digitOneInNumber(13)) // 6
    // Example 2:
    // Input: n = 0
    // Output: 0
    fmt.Println(digitOneInNumber(0)) // 0

    fmt.Println(digitOneInNumber(1)) // 1
}