package main

// LCP 17. 速算机器人
// 小扣在秋日市集发现了一款速算机器人。
// 店家对机器人说出两个数字（记作 x 和 y），请小扣说出计算指令：
//     1. "A" 运算：使 x = 2 * x + y；
//     2. "B" 运算：使 y = 2 * y + x。

// 在本次游戏中，店家说出的数字为 x = 1 和 y = 0，小扣说出的计算指令记作仅由大写字母 A、B 组成的字符串 s，
// 字符串中字符的顺序表示计算顺序，请返回最终 x 与 y 的和为多少。

// 示例 1：
// 输入：s = "AB"
// 输出：4
// 解释： 经过一次 A 运算后，x = 2, y = 0。 再经过一次 B 运算，x = 2, y = 2。 最终 x 与 y 之和为 4。

// 提示：
//     0 <= s.length <= 10
//     s 由 'A' 和 'B' 组成

import "fmt"

func calculate(s string) int {
    x, y := 1, 0
    for i := 0; i < len(s); i++ {
        if s[i] == 'A' { // "A" 运算：使 x = 2 * x + y；
            x = 2 * x + y
        }
        if s[i] == 'B' { // "B" 运算：使 y = 2 * y + x。
            y = 2 * y + x
        }
    }
    return x + y
}

func calculate1(s string) int {
    return 1 << len(s)
}

func main() {
    // 示例 1：
    // 输入：s = "AB"
    // 输出：4
    // 解释： 经过一次 A 运算后，x = 2, y = 0。 再经过一次 B 运算，x = 2, y = 2。 最终 x 与 y 之和为 4。
    fmt.Println(calculate("AB")) // 4

    fmt.Println(calculate("A")) // 2
    fmt.Println(calculate("B")) // 2
    fmt.Println(calculate("AA")) // 4
    fmt.Println(calculate("BB")) // 4
    fmt.Println(calculate("AAB")) // 8
    fmt.Println(calculate("BBA")) // 8
    fmt.Println(calculate("AABB")) // 16
    fmt.Println(calculate("BBAA")) // 16
    fmt.Println(calculate("ABAB")) // 16
    fmt.Println(calculate("BABA")) // 16

    fmt.Println(calculate1("AB")) // 4
    fmt.Println(calculate1("A")) // 2
    fmt.Println(calculate1("B")) // 2
    fmt.Println(calculate1("AA")) // 4
    fmt.Println(calculate1("BB")) // 4
    fmt.Println(calculate1("AAB")) // 8
    fmt.Println(calculate1("BBA")) // 8
    fmt.Println(calculate1("AABB")) // 16
    fmt.Println(calculate1("BBAA")) // 16
    fmt.Println(calculate1("ABAB")) // 16
    fmt.Println(calculate1("BABA")) // 16
}