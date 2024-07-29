package main

// LCR 163. 找到第 k 位数字
// 某班级学号记录系统发生错乱，原整数学号序列 [0,1,2,3,4,...] 分隔符丢失后变为 01234... 的字符序列。
// 请实现一个函数返回该字符序列中的第 k 位数字。

// 示例 1：
// 输入：k = 5
// 输出：5

// 示例 2：
// 输入：k = 12
// 输出：1
// 解释：第 12 位数字在序列 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... 里是 1 ，它是 11 的一部分。

// 提示：
//     0 <= k < 2^31

import "fmt"
import "math"

// bits = 1 的时候有 1,2,3,4,5,6,7,8,9 这 9 个数; 9 = math.Pow10(bits - 1) * bits
// bits = 2 的时候有 10-99 这 90 个数; 90 = math.Pow10(bits - 1) * bits
// n 不断减去 bits 从 1 开始的数字总数,求出 n 所在的数字是几位数即 bits
// 计算 n 所在的数字 num，等于初始值加上 (n - 1) / bits
// 计算 n 所在这个数字的第几位 digitIdx 等于 (n - 1) % bits
// 计算出 digitIdx 位的数字
// 以11 为例:
//     11 - 9 = 2
//     (2 - 1) / 2 = 0
//     (2 - 1) % 2 = 1
// 也就是说第 11 位数字是位数是 2 的第一个数字的第二位，即是 0
func findKthNumber(n int) int {
    if n <= 9 {
        return n
    }
    bits := 1
    for n > 9 * int(math.Pow10(bits-1)) * bits {
        n -= 9 * int(math.Pow10(bits-1)) * bits
        bits++
    }
    idx, start := n - 1, int(math.Pow10(bits - 1))
    num, digitIdx := start + idx / bits, idx % bits
    return num / int(math.Pow10(bits - digitIdx - 1)) % 10
}

func main() {
    // Example 1:
    // Input: n = 3
    // Output: 3
    fmt.Println(findKthNumber(3)) // 3
    // Example 2:
    // Input: n = 11
    // Output: 0
    // Explanation: The 11th digit of the sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... is a 0, which is part of the number 10.
    fmt.Println(findKthNumber(11)) // 0
}